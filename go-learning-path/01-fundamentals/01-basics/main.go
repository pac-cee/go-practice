// Package main demonstrates basic Go programming concepts
package main

import (
	"fmt"
	"strings"
	"time"
)

// Greeting returns a greeting message based on the time of day
func Greeting(name string) string {
	hour := time.Now().Hour()
	var timeOfDay string

	switch {
	case hour < 12:
		timeOfDay = "morning"
	case hour < 17:
		timeOfDay = "afternoon"
