package main

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes"

	micro "github.com/micro/go-micro"
	proto "github.com/wz185/go-example/app/interface/grpc/pause"
)

func main() {
	service := micro.NewService(
		micro.Name("package"),
	)

	service.Init()

	client := proto.NewPauseService("package", service.Client())

	res, err := client.Create(context.TODO(), &proto.PauseRequest{
		StartDate:      ptypes.TimestampNow(),
		EndDate:        ptypes.TimestampNow(),
		Type:           proto.PauseType_STANDARD,
		ValueProductId: "123",
		AccountId:      "123",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.ResposneInfo)
}
