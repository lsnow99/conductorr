package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	urlString, ok := request.QueryStringParameters["url"]
	if !ok {
		return makeErrorResponse(http.StatusBadRequest, fmt.Errorf("path parameter 'url' is missing"))
	}

	u, err := url.Parse(urlString)
	if err != nil {
		return makeErrorResponse(http.StatusBadRequest, fmt.Errorf("error decoding url: %s", err.Error()))
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return makeErrorResponse(http.StatusInternalServerError, fmt.Errorf("error building request: %s", err.Error()))
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return makeErrorResponse(http.StatusBadGateway, fmt.Errorf("error performing proxy request: %s", err.Error()))
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return makeErrorResponse(http.StatusBadGateway, fmt.Errorf("error reading proxy response body: %s", err.Error()))
	}

	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "text/plain"},
		Body:            string(data),
		IsBase64Encoded: false,
	}, nil
}

func makeErrorResponse(statusCode int, err error) (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		StatusCode:      statusCode,
		Headers:         map[string]string{"Content-Type": "text/plain"},
		Body:            err.Error(),
		IsBase64Encoded: false,
	}, err
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
