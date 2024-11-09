package main

import (
	"log"
	"time"
)

type Task struct {
	Name     string         `json:"name"`
	Interval time.Duration  `json:"interval"`
	Timezone *time.Location `json:"timezone"`
	lastRun  time.Time
}

type ScheduleRequest struct {
	Name     string `json:"name"`
	Interval int    `json:"interval"` // in seconds
	Timezone string `json:"timezone"`
}

func NewTask(name string, interval int, tz string) *Task {
	location, err := time.LoadLocation(tz)
	if err != nil {
		log.Printf("Invalid timezone %s, using UTC\n", tz)
		location = time.UTC
	}

	return &Task{
		Name:     name,
		Interval: time.Duration(interval) * time.Second,
		Timezone: location,
	}
}

func (t *Task) ShouldRun(now time.Time) bool {
	if t.lastRun.IsZero() {
		return true
	}
	return now.Sub(t.lastRun) >= t.Interval
}

func (t *Task) Run() {
	t.lastRun = time.Now().In(t.Timezone)
	log.Printf("Running task: %s at %v\n", t.Name, t.lastRun)
	// Here, add your task execution logic, e.g., making an API call or processing a job.
}
