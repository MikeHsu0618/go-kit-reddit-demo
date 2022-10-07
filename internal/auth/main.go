package main

import (
	"context"
	"fmt"
	"go-kit-reddit-demo/internal/auth/client/http"
)

func main() {
	service, err := http.New("localhost:8081/", nil)
	if err != nil {
		return
	}
	token, err := service.GenerateToken(context.Background(), 123)
	if err != nil {
		return
	}
	fmt.Println(token)
}
