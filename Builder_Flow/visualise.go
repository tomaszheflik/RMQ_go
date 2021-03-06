package main

import (
"fmt"
"time"
)
func update(data []Build, element Build ) []Build {
	for key, value := range data {
		if value.ID == element.ID {
			data[key] = element
			return data
		} 
	}
	return append(data, element)
}

func delete(data []Build, element Build) []Build {
	for key, value := range data {
		if value.ID == element.ID {
			if key == len(data) {
				data = data[:len(data)-1]
			} else {
				return append(data[:key], data[key+1:]... )
			}
		}
	}
	return data
}

func visualise(data []Build){
	fmt.Print("\033[H\033[2J")
	for _, value := range data {
		fmt.Println(value)
	}
}

func main() {
	rabbit := "amqp://guest:guest@192.168.99.100:32100"
	readQueueName := "init_job"
	builderQueue := "build"
	statusQueue := "status"
	var build, statusOfBuild Build
	var toProcess, toStatus string
	var buildData []Build

	message := make(chan string)
	go func(){
			readFromQueue(rabbit, readQueueName, message)
	}()

	statusChan := make(chan string)
	go func(){
			readFromQueue(rabbit, statusQueue, statusChan)
	}()

	go func(){
		for {
			toProcess = <-message
			sendToQueue(toProcess, rabbit, builderQueue)
			build = getBuildStruct(toProcess)
			buildData = update(buildData, build)
		}
	}()

	go func(){
		for {
			toStatus = <-statusChan
			statusOfBuild = getBuildStruct(toStatus)
			if statusOfBuild.Status == "Building" {
				buildData = update(buildData, statusOfBuild)
			}
			if statusOfBuild.Status == "Done" {
				buildData = delete(buildData, statusOfBuild)
			}
		}
	}()
	for {
		visualise(buildData)
		time.Sleep(1 * time.Second)
	}
}
	