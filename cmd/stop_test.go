package cmd

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestStopSuccess(t *testing.T) {
	cases := []struct {
		command string
		want    string
	}{
		{command: "dakoku stop 1", want: "stop task 1[task_title].\n"},
	}

	dbPath := GetDbPath()
	for _, c := range cases {
		InitDbForTest(new(bytes.Buffer))
		CreateTaskForTest(new(bytes.Buffer))
		StartTaskForTest(new(bytes.Buffer))

		buf := new(bytes.Buffer)
		cmd := NewCmdStop()
		cmd.SetOutput(buf)
		cmdArgs := strings.Split(c.command, " ")
		fmt.Printf("cmdArgs %+v\n", cmdArgs)
		cmd.SetArgs(cmdArgs[2:])
		cmd.Execute()

		get := buf.String()
		if c.want != get {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
		os.Remove(dbPath)
	}
}

func TestStopDbNotInitError(t *testing.T) {
	cases := []struct {
		command string
		want    string
	}{
		{command: "dakoku stop", want: "failed to stop a task. please initialize dakoku.\n"},
	}

	for _, c := range cases {
		buf := new(bytes.Buffer)
		cmd := NewCmdStop()
		cmd.SetOutput(buf)
		cmdArgs := strings.Split(c.command, " ")
		fmt.Printf("cmdArgs %+v\n", cmdArgs)
		cmd.SetArgs(cmdArgs[2:])
		cmd.Execute()

		get := buf.String()
		if !strings.Contains(get, c.want) {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

func TestStopArgsError(t *testing.T) {
	cases := []struct {
		command string
		want    string
	}{
		{command: "dakoku stop", want: "failed to stop a task. task's id is required.\n"},
	}

	dbPath := GetDbPath()
	for _, c := range cases {
		InitDbForTest(new(bytes.Buffer))
		CreateTaskForTest(new(bytes.Buffer))

		buf := new(bytes.Buffer)
		cmd := NewCmdStop()
		cmd.SetOutput(buf)
		cmdArgs := strings.Split(c.command, " ")
		fmt.Printf("cmdArgs %+v\n", cmdArgs)
		cmd.SetArgs(cmdArgs[2:])
		cmd.Execute()

		get := buf.String()
		if !strings.Contains(get, c.want) {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
		os.Remove(dbPath)
	}
}

func TestStopNoTaskError(t *testing.T) {
	cases := []struct {
		command string
		want    string
	}{
		{command: "dakoku stop 1", want: "failed to stop tasks. please check the task exists and is working.\n"},
	}

	dbPath := GetDbPath()
	for _, c := range cases {
		InitDbForTest(new(bytes.Buffer))

		buf := new(bytes.Buffer)
		cmd := NewCmdStop()
		cmd.SetOutput(buf)
		cmdArgs := strings.Split(c.command, " ")
		fmt.Printf("cmdArgs %+v\n", cmdArgs)
		cmd.SetArgs(cmdArgs[2:])
		cmd.Execute()

		get := buf.String()
		if !strings.Contains(get, c.want) {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
		os.Remove(dbPath)
	}
}

func TestStopNotWorkingTaskError(t *testing.T) {
	cases := []struct {
		command string
		want    string
	}{
		{command: "dakoku stop 1", want: "failed to stop tasks. please check the task exists and is working.\n"},
	}

	dbPath := GetDbPath()
	for _, c := range cases {
		InitDbForTest(new(bytes.Buffer))
		CreateTaskForTest(new(bytes.Buffer))

		buf := new(bytes.Buffer)
		cmd := NewCmdStop()
		cmd.SetOutput(buf)
		cmdArgs := strings.Split(c.command, " ")
		fmt.Printf("cmdArgs %+v\n", cmdArgs)
		cmd.SetArgs(cmdArgs[2:])
		cmd.Execute()

		get := buf.String()
		if !strings.Contains(get, c.want) {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
		os.Remove(dbPath)
	}
}
