package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

type Scheduler struct {
	tasks map[string]*Task
	mu    sync.Mutex
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks: make(map[string]*Task),
	}
}

func (s *Scheduler) Start() {
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		s.runTasks()
	}
}

func (s *Scheduler) runTasks() {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	for _, task := range s.tasks {
		if task.ShouldRun(now) {
			go task.Run()
		}
	}
}

func (s *Scheduler) HandleSchedule(w http.ResponseWriter, r *http.Request) {
	var req ScheduleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	task := NewTask(req.Name, req.Interval, req.Timezone)
	s.mu.Lock()
	s.tasks[task.Name] = task
	s.mu.Unlock()

	log.Printf("Task %s scheduled to run every %v\n", task.Name, task.Interval)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (s *Scheduler) HandleListTasks(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()

	tasks := make([]*Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}
