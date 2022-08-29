package logger

import (
	"fmt"
	"time"
)

var currentTime = time.Now().Format("15:04:05")

// Info -> Print an Info line to the console.
func Info(text string) {
	fmt.Printf("[%v] %v[INFO]%v: %v\n", currentTime, "\033[32m", "\033[0m", text)
}

// Debug -> Print a Debug line to the console.
func Debug(text string) {
	fmt.Printf("[%v] %v[DEBUG]%v: %v\n", currentTime, "\033[36m", "\033[0m", text)
}

// Warn -> Print a Warn line to the console.
func Warn(text string) {
	fmt.Printf("[%v] %v[WARN]%v: %v\n", currentTime, "\033[33m", "\033[0m", text)
}

// Error -> Print an Error line to the console.
func Error(text string) {
	fmt.Printf("[%v] %v[ERROR]%v: %v\n", currentTime, "\033[31m", "\033[0m", text)
}
