package main

import (
	"fmt"
	"os"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	eventTime, err := time.Parse(time.RFC3339, "2022-07-20T09:44:21Z")
	eventTime2, err := time.Parse(time.RFC3339, "2022-07-20T09:44:21+09:00")
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(timestamppb.New(eventTime))
	fmt.Println(timestamppb.New(eventTime2))
}