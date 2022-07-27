package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"time"

	v1 "github.kakaoenterprise.in/cloud/msg/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func getInstanceError() *v1.Instance {
	return &v1.Instance{
		Metadata: &v1.ObjectMeta{
			Service:        "MSG",
			Project:         "e30bf7e8fdc24320b0ed1dd184dade7f",
			Name:            "hayz-test-2",
		},
		Spec: &v1.InstanceSpec{
			Compute: &v1.ComputeSpec{
				FlavorId: "090c2665-a882-49ca-963c-b008b7a6075c",
			},
			NetworkPorts: []*v1.NetworkPortSpec{
				{
					Name:      "hayz-port-test-1",
					NetworkId: "8ac1349b-1c98-41b9-a627-25c69f125f23",
					SecurityGroupIds: []string{"7cd61389-0e81-44d7-9392-e483087ceebb"},
				},
			},
			Volumes: []*v1.VolumeSpec{
				{
					Name: "hayz-vol-test-1",
					Size: 100,
					Source: &v1.VolumeSource{
						ImageId: "9e760d22-64c1-45a1-a426-0dc064eb18f6",
					},
				},
			},
			Apps: []*v1.AppSpec{
				{
					Name: "hayz-app-test-pre",
					Run: &v1.Run{
						Command: []string{"wget object.."},
					},
				},
			},
		},
	}
}

func main() {
	dialOptions := make([]grpc.DialOption, 0, 2)
	credential := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
	dialOptions = append(dialOptions, grpc.WithTransportCredentials(credential))
	dialOptions = append(dialOptions, grpc.WithBlock())

	conn, err := grpc.Dial("10.183.48.57:30111", dialOptions...)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewInstanceServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	instance, err := c.Create(ctx, getInstanceError())
	fmt.Println(err)
	fmt.Println(instance)


}
