package main

import (
	"time"
)

type Build struct {
	ID string
	Payload time.Duration
	Status string
}