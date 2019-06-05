package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/davegarred/wh3/persist"
	"testing"
)

func TestHandleRequest(t *testing.T) {
	handler := &RequestHandler{persist.NewDynamoClient()}
	res,_ := handler.HandleRequest(context.Background(), events.APIGatewayProxyRequest{})
	val,_ := json.Marshal(res)
	fmt.Println(string(val))
}
