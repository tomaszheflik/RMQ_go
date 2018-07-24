package main

import (
	"fmt"
	"time"
)
func process(rabbitHost string, statusQueseue string, build Build) error {
	build.Status = "Building"
	err := sendToQueue(getBuildJSON(build), rabbitHost, statusQueseue)
	failOnError(err, "Unable to send status")
	if err != nil {
		return err
	}

	time.Sleep(time.Second * build.Payload)
	build.Status = "Done"
	fmt.Printf("Job status: %+v\n", build)
	err = sendToQueue(getBuildJSON(build), rabbitHost, statusQueseue)
	failOnError(err, "Unable to send status")
	if err != nil {
		return err
	}
	return err
}

func main() {
	rabbit := "amqp://guest:guest@192.168.99.100:32100"
	builderQueue := "build"
	statusQueue := "status"
	var build Build
	var toProcess Message

	message := make(chan Message)
	go func(){
			readFromQueue(rabbit, builderQueue, message)
	}()

	for {
		toProcess = <-message
		fmt.Printf("Job for Builder: %s\n", toProcess.BODY)
		build = getBuildStruct(toProcess.BODY)
		err := process(rabbit, statusQueue, build)
		if err != nil {
			return 
		}
		toProcess.ACK <- true
	}
}
