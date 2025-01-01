package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

const (
	PRODUCER_HOST     = "http://localhost"
	PRODUCER_PORT     = 9999
	PRODUCER_BASE_URL = "/api/v1/producer/"
	QUEUE_HOST        = "http://localhost"
	QUEUE_PORT        = 10000
	QUEUE_BASE_URL    = "/api/v1/queue/"
)

var (
	ErrOnParse           = errors.New("data could not be parsed")
	ErrOnRequestCreation = errors.New("request could not be created")
	ErrOnRequest         = errors.New("request could not be send")
)

type messageRequest struct {
	Message string `json:"message"`
}

type messageRequestResponse struct {
	Code         int    `json:"code"`
	ErrorMessage string `json:"error_message,omitempty"`
}

type apiHandler struct{}

func newApiHandler() *apiHandler {
	return &apiHandler{}
}

func (ah *apiHandler) messageHandler(c echo.Context) error {
	var requestBody messageRequest

	if err := c.Bind(&requestBody); err != nil {
		response := messageRequestResponse{
			Code:         http.StatusBadRequest,
			ErrorMessage: fmt.Sprintf("[Producer] Request Body could not be parsed! Error: %v\n", err),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	message := requestBody.Message
	fmt.Printf("[Producer] Producer has obtained the following message: %s\n", message)
	err := sendMessageToQueue(message)
	if err != nil {
		response := messageRequestResponse{
			Code:         http.StatusBadRequest,
			ErrorMessage: fmt.Sprintf("[Producer] Could not send message to queue! Error: %v\n", err),
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	fmt.Printf("[Producer] Could successfully send message to queue!\n")

	response := messageRequestResponse{
		Code: http.StatusOK,
	}
	return c.JSON(http.StatusOK, response)
}

func sendMessageToQueue(message string) error {
	url := QUEUE_HOST + ":" + strconv.Itoa(QUEUE_PORT) + QUEUE_BASE_URL + "send"
	data := messageRequest{
		Message: message,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return ErrOnParse
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		return ErrOnRequestCreation
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ErrOnRequest
	}
	defer resp.Body.Close()

	return nil
}

func main() {
	ah := newApiHandler()

	echoServer := echo.New()

	echoServer.POST(PRODUCER_BASE_URL+"message", ah.messageHandler)

	serverAddress := fmt.Sprintf(":%d", PRODUCER_PORT)
	if err := echoServer.Start(serverAddress); err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("[Producer] Could not start server successfully! Error: %s\n", err)
		os.Exit(1)
	}
}
