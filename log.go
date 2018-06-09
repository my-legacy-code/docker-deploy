package main

import (
	"fmt"
	"time"
)

func log(message string) {
	now := time.Now().Format("2006 Jan 17 _2 15:04:05")
	fmt.Printf("[DOCKER DEPLOY][%s] %s\n", now, message)
}
