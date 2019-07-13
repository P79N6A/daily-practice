package utils

import "time"

func GetTime(client chan string) {
	now := time.Now().Format("2006-01-02 15:04:05")
	client <- now
}
