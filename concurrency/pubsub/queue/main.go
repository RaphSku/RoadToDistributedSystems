package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	QUEUE_HOST        = "http://localhost"
	QUEUE_PORT        = 10000
	QUEUE_BASE_URL    = "/api/v1/queue/"
	CONSUMER_BASE_URL = "/api/v1/consumer/"
)

var (
	ErrOnParse           = errors.New("parsing failed")
	ErrOnRequestCreation = errors.New("request could not be created")
	ErrOnRequest         = errors.New("request could not be send")
	ErrOnEmptyQueue      = errors.New("queue is empty")
)

type queue struct {
	values []string
}

func newQueue() *queue {
	return &queue{}
}

func (q *queue) enqueue(value string) {
	q.values = append(q.values, value)
}

func (q *queue) dequeue() (string, error) {
	if len(q.values) == 0 {
		return "", ErrOnEmptyQueue
	}

	value := q.values[len(q.values)-1]
	q.values = q.values[:len(q.values)-1]

	return value, nil
}

type subscriberRequest struct {
	ConsumerAddress string `json:"consumer_address"`
}

type subscriberResponse struct {
	Code         int    `json:"code"`
	ErrorMessage string `json:"error_message"`
}

type consumerRequest struct {
	Message string `json:"message"`
}

type producerRequest struct {
	Message string `json:"message"`
}

type producerResponse struct {
	Code         int    `json:"code"`
	ErrorMessage string `json:"error_message"`
}

type apiHandler struct {
	mtx          sync.Mutex
	subscriber   []string
	messageQueue *queue
}

func newApiHandler() *apiHandler {
	queue := newQueue()
	return &apiHandler{
		messageQueue: queue,
	}
}

func (ah *apiHandler) handleSubscription(c echo.Context) error {
	fmt.Printf("[Queue] A consumer wants to subscribe!\n")
	var requestBody subscriberRequest

	if err := json.NewDecoder(c.Request().Body).Decode(&requestBody); err != nil {
		response := subscriberResponse{
			Code:         http.StatusBadRequest,
			ErrorMessage: fmt.Sprintf("[Queue] Handling of subscription went wrong! Error: %v\n", err),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	consumerAddress := requestBody.ConsumerAddress
	ah.mtx.Lock()
	ah.subscriber = append(ah.subscriber, consumerAddress)
	ah.mtx.Unlock()
	fmt.Printf("[Queue] The following consumer has subscribed to this queue: %s\n", consumerAddress)

	response := subscriberResponse{
		Code:         http.StatusOK,
		ErrorMessage: "",
	}
	return c.JSON(http.StatusOK, response)
}

func (ah *apiHandler) handleProducer(c echo.Context) error {
	fmt.Printf("[Queue] A producer is trying to push a message!\n")

	var requestBody producerRequest

	if err := json.NewDecoder(c.Request().Body).Decode(&requestBody); err != nil {
		response := producerResponse{
			Code:         http.StatusBadRequest,
			ErrorMessage: fmt.Sprintf("[Queue] Handling of Producer message went wrong! Error: %v\n", err),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	message := requestBody.Message
	fmt.Printf("[Queue] The following message has been sent: %s\n", message)
	ah.mtx.Lock()
	ah.messageQueue.enqueue(message)
	ah.mtx.Unlock()
	fmt.Printf("[Queue] The following message was added to the queue: %s\n", message)

	response := producerResponse{
		Code:         http.StatusOK,
		ErrorMessage: "",
	}
	return c.JSON(http.StatusOK, response)
}

func deliverMessagesToSubscriber(ctx context.Context, ah *apiHandler) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Gracefully shut down delivery of messages!\n")
			return
		default:
			ah.mtx.Lock()
			message, err := ah.messageQueue.dequeue()
			ah.mtx.Unlock()
			fmt.Printf("[Queue] The following message was taken from queue: %s\n", message)
			if err == ErrOnEmptyQueue {
				time.Sleep(5 * time.Second)
				continue
			}

			for _, consumerAddress := range ah.subscriber {
				consumerUrl := consumerAddress + CONSUMER_BASE_URL + "push"

				data := consumerRequest{
					Message: message,
				}

				jsonData, err := json.Marshal(data)
				if err != nil {
					fmt.Printf("[Queue] Could not parse message during delivery for consumer: %s! Error: %v\n", consumerAddress, ErrOnParse)
					continue
				}

				req, err := http.NewRequest("POST", consumerUrl, bytes.NewBuffer([]byte(jsonData)))
				if err != nil {
					fmt.Printf("[Queue] Could not create new request for consumer: %s! Error: %v\n", consumerAddress, ErrOnRequestCreation)
					continue
				}

				req.Header.Set("Content-Type", "application/json")

				client := &http.Client{}
				resp, err := client.Do(req)
				if err != nil {
					fmt.Printf("[Queue] Could not send the request for consumer: %s! Error: %v\n", consumerAddress, ErrOnRequest)
					continue
				}
				fmt.Printf("[Queue] Message was successfully delivered!\n")
				defer resp.Body.Close()
			}
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ah := newApiHandler()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		deliverMessagesToSubscriber(ctx, ah)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		echoServer := echo.New()

		echoServer.POST(QUEUE_BASE_URL+"subscribe", ah.handleSubscription)
		echoServer.POST(QUEUE_BASE_URL+"send", ah.handleProducer)

		go func() {
			<-ctx.Done()
			if err := echoServer.Shutdown(ctx); err != nil {
				fmt.Printf("[Queue] Error shutting down server: %v\n", err)
			}
		}()

		serverAddress := fmt.Sprintf(":%d", QUEUE_PORT)
		if err := echoServer.Start(serverAddress); err != nil {
			if err == http.ErrServerClosed {
				fmt.Printf("[Queue] Server closes down...\n")
				return
			}
			fmt.Printf("[Queue] Could not start server successfully! Error: %v\n", err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	<-sigChan
	fmt.Printf("[Queue] Termination signal received. Shutting down...\n")

	cancel()

	wg.Wait()
	fmt.Printf("[Queue] Exited gracefully.\n")
}
