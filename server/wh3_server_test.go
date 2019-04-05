package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"testing"
)

func _TestHandleRequest(t *testing.T) {
	res,_ := HandleRequest(context.Background(), events.APIGatewayProxyRequest{})
	val,_ := json.Marshal(res)
	fmt.Println(string(val))
}
