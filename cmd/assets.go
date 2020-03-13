package cmd

import (
	"os"
	"time"
)

const (
	Off     = 0
	On      = 1
	Deleted = -1
)

type Options struct {
	days  int
	isAll bool
}

type Task struct {
	id    int
	title string
	state int
}

type Log struct {
	id     int
	taskId int
	start  time.Time
	end    time.Time
	period time.Duration
	state  int
}

type Result struct {
	task        *Task
	log         *Log
	sumOfPeriod time.Duration
	stateView   string
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func GetDbPath() string {
	workDir, _ := os.Getwd()
	return workDir + "/db.sql"
}
