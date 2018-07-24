package main

import (
	"time"
)
// Build - 
type Build struct {
	ID string
	Payload time.Duration
	Status string
}

// Message - 
type Message struct {
	BODY string
	ACK chan bool
}