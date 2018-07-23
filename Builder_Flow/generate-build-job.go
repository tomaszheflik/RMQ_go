package main

import (
	"math/rand"
	"time"
	"fmt"
)
 func init() {
	 rand.Seed(time.Now().UnixNano())
 }

func generateID() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	id := make([]rune, 10)
	for i := range id {
		id[i] = letters[rand.Intn(len(letters))]
	}
	return string(id)
}

func generatePayload(limit int) time.Duration {
	return time.Duration(rand.Intn(limit))
}


func main() {
	var build Build
	rabbit := "amqp://guest:guest@192.168.99.100:32100"
	queueName := "init_job"
	
	for {
		build.ID = generateID()
		build.Payload = generatePayload(10)
		build.Status = "Inqueue"
		fmt.Printf("Sending build: %s\n", build)
		err := sendToQueue(getBuildJSON(build), rabbit, queueName) 
		if err != nil {
			break
		}
		time.Sleep(time.Duration(generatePayload(20)) * time.Second)
	}
	fmt.Printf("End of sending.")
}