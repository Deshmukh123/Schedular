package main

import (
	"log"
	"net/http"
)

func main() {
	scheduler := NewScheduler()
	go scheduler.Start()

	http.HandleFunc("/schedule", scheduler.HandleSchedule)
	http.HandleFunc("/list", scheduler.HandleListTasks)

	log.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
