package main

import (
	"github.com/sirupsen/logrus"
	"income_flow_rest_service/handler"
)

func main() {
	handlers := handler.NewHandler()

	srv := NewServer()
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		logrus.Error(err)
	}
}
