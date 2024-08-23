package Eureka_config

import (
	"fmt"

	"github.com/ArthurHlt/go-eureka-client/eureka"
)

func EurekaClientConfig() {
	client := eureka.NewClient([]string{"http://eureka-server:8761/eureka"})
	fmt.Println("Printing Client Details..")
	fmt.Println(client)
	instance := eureka.NewInstanceInfo(
		"golang_parser_service",
		"golang_parser_service",
		"127.0.0.1",
		9090,
		30,
		false,
	)

	client.RegisterInstance("golang_parser_service", instance)
	client.SendHeartbeat(instance.App, instance.HostName)
	fmt.Println("Printing Instance Details...")
	fmt.Println(client.GetInstance(instance.App, instance.HostName))
}
