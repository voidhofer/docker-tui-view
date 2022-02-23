package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func GetServices() []string {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	services, err := cli.ServiceList(context.Background(), types.ServiceListOptions{})
	if err != nil {
		panic(err)
	}

	var serviceList []string

	for _, service := range services {
		serviceList = append(serviceList, service.ID)
		fmt.Printf("%s %s\n", service.ID)
	}

	return serviceList
}
