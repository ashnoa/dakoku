package cmd

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestStartSuccess(t *testing.T) {
	cases := []struct {
		command string
		want    string
	}{
		{command: "dakoku start 1", want: "start task 1[task_title].\n"},
	}

	dbPath := GetDbPath()
	for _, c := range cases {
		InitDbForTest(new(bytes.Buffer))
		CreateTaskForTest(new(bytes.Buffer))

		buf := new(bytes.Buffer)
		cmd := NewCmdStart()
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

func TestStartDbNotInitError(t *testing.T) {
	cases := []struct {
		command string
		want    string
	}{
		{command: "dakoku start", want: "failed to start a task. please initialize dakoku.\n"},
	}

	for _, c := range cases {
		buf := new(bytes.Buffer)
		cmd := NewCmdStart()
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

func TestStartArgsError(t *testing.T) {
	cases := []struct {
		command string
		want    string
	}{
		{command: "dakoku start", want: "failed to start a task. task's id is required.\n"},
	}

	dbPath := GetDbPath()
	for _, c := range cases {
		InitDbForTest(new(bytes.Buffer))
		CreateTaskForTest(new(bytes.Buffer))

		buf := new(bytes.Buffer)
		cmd := NewCmdStart()
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

func TestStartNoTaskError(t *testing.T) {
	cases := []struct {
		command string
		want    string
	}{
		{command: "dakoku start 1", want: "failed to start tasks. please check the task exists and isn't working.\n"},
	}

	dbPath := GetDbPath()
	for _, c := range cases {
		InitDbForTest(new(bytes.Buffer))

		buf := new(bytes.Buffer)
		cmd := NewCmdStart()
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
