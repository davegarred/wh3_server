package main

import (
	"context"
	"testing"
)

func TestHandleRequest(t *testing.T) {
	_,err := HandleRequest(context.Background(), nil)
	if err != nil {
		panic(err)
	}
}

