package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/labstack/echo/v4"
)

const (
	CONSUMER_HOST     = "http://localhost"
	CONSUMER_PORT     = 10001
	CONSUMER_BASE_URL = "/api/v1/consumer/"
	QUEUE_HOST        = "http://localhost"
	QUEUE_PORT        = 10000
	QUEUE_BASE_URL    = "/api/v1/queue/"
)

var messages []string
var mtx sync.RWMutex

type messageRequestResponse struct {
	Code     int      `json:"code"`
	Messages []string `json:"messages"`
}

type messageReceiverRequestBody struct {
	Message string `json:"message"`
}

type messageReceiverRequestResponse struct {
	Code int `json:"code"`
}

type subscribeRequest struct {
	ConsumerAddress string `json:"consumer_address"`
}

type apiHandler struct{}

func newApiHandler() *apiHandler {
	return &apiHandler{}
}

func (ah *apiHandler) messageRequestHandler(c echo.Context) error {
	mtx.RLock()
	response := messageRequestResponse{
		Code:     http.StatusOK,
		Messages: messages,
	}
	mtx.RUnlock()
	return c.JSON(http.StatusOK, response)
}

func (ah *apiHandler) messageReceiverHandler(c echo.Context) error {
	fmt.Printf("[Consumer] The queue tries to push a message!\n")
	var messageRequest messageReceiverRequestBody

	if err := json.NewDecoder(c.Request().Body).Decode(&messageRequest); err != nil {
		response := messageReceiverRequestResponse{
			Code: http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	message := messageRequest.Message
	mtx.Lock()
	messages = append(messages, message)
	mtx.Unlock()
	fmt.Printf("[Consumer] The queue has sent the following message: %s\n", message)

	response := messageReceiverRequestResponse{
		Code: http.StatusOK,
	}
	return c.JSON(http.StatusOK, response)
}

func subscribeToQueue() bool {
	url := QUEUE_HOST + ":" + strconv.Itoa(QUEUE_PORT) + QUEUE_BASE_URL + "subscribe"

	data := subscribeRequest{
		ConsumerAddress: CONSUMER_HOST + ":" + strconv.Itoa(CONSUMER_PORT),
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("[Consumer] Cannot parse data! Error: %v\n", err)
		return false
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("[Consumer] Cannot contact the queue server! Error: %v\n", err)
		return false
	}
	defer resp.Body.Close()
	return true
}

func main() {
	ok := subscribeToQueue()
	if !ok {
		fmt.Print("[Consumer] Could not subscribe to queue server!\n")
		os.Exit(1)
	}

	ah := newApiHandler()

	echoServer := echo.New()

	echoServer.GET(CONSUMER_BASE_URL+"messages", ah.messageRequestHandler)
	echoServer.POST(CONSUMER_BASE_URL+"push", ah.messageReceiverHandler)

	serverAddress := fmt.Sprintf(":%d", CONSUMER_PORT)
	if err := echoServer.Start(serverAddress); err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("[Consumer] Could not start server successfully! Error: %s\n", err)
		os.Exit(1)
	}
}
