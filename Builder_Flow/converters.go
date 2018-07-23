package main

import (
	"encoding/json"
	"fmt"
)

func getBuildJSON(build Build) string {
	jsonBuild, err := json.Marshal(build)
	failOnError(err, "Unable to marshal json")
	return string(jsonBuild)
}

func getBuildStruct(message string) Build {
	var build Build
	err := json.Unmarshal([]byte(message), &build)
	if err != nil {
		fmt.Printf("Unmarshall error: %s\n", err)
	}
	return build
}
