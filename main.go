package main

import (
	"log"

	micro "github.com/micro/go-micro"
	applicationPause "github.com/wz185/go-example/app/interface/controllers/pause"
	protoPause "github.com/wz185/go-example/app/interface/grpc/pause"
)

func main() {
	service := micro.NewService(
		micro.Name("package"), // client name needs to match to this service name
	)

	service.Init()

	handler := new(applicationPause.Handler)
	err := protoPause.RegisterPauseHandler(service.Server(), handler)

	if err != nil {
		log.Fatalln(err)
	}

	// Run the server
	if err := service.Run(); err != nil {
		log.Fatalln(err)
	}
}
