package main

import (
	"github.com/adjust/queue"
)

func main() {
	connection := queue.OpenConnection("cleaner", "tcp", "localhost:6379", 2)
	queue := connection.OpenQueue("things")
	queue.Purge()
}
