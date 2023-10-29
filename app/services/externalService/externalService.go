package externalService

import (
	"net/http"
	"receipt-processor-module/pkg/models"
	"encoding/json"
	"io"
	"reflect"
	"bytes"
	"fmt"
)


func GetExternalApiResponse(receipt models.Receipt) (interface{}, error )  {
	url := "https://jsonplaceholder.typicode.com/post"
	httpMethod := "POST"
	var buf bytes.Buffer
  json.NewEncoder(&buf).Encode(receipt)

	request, err := http.NewRequest(httpMethod, url, &buf)
	if err != nil {
		return nil, models.CustomError{HttpCode: 500, Message:"Internal server error", DebugMessage:"Could not create Request for external API"}
	}
	response, err := http.DefaultClient.Do(request)
	defer response.Body.Close()
	if err != nil {
		return nil, models.CustomError{HttpCode: 500, Message:"Internal server error", DebugMessage:"Problem fetching external data"}
	}

	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, models.CustomError{HttpCode: 500, Message:"Internal server error", DebugMessage:"Problem fetching external data"}
	}
	
	var parsedBody interface{}
	json.Unmarshal(respBody, &parsedBody)
	fmt.Println( parsedBody, reflect.TypeOf(parsedBody))
	return parsedBody.(map[string]interface{}), nil

}