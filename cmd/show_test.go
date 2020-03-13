package cmd

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestShowSuccess(t *testing.T) {
	cases := []struct {
		command string
		want    string
	}{
		{command: "dakoku show", want: "1 | task_title | 0s |"},
	}

	dbPath := GetDbPath()
	for _, c := range cases {
		InitDbForTest(new(bytes.Buffer))
		CreateTaskForTest(new(bytes.Buffer))

		buf := new(bytes.Buffer)
		cmd := NewCmdShow()
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

func TestShowDbNotInitError(t *testing.T) {
	cases := []struct {
		command string
		want    string
	}{
		{command: "dakoku show", want: "failed to show tasks. please initialize dakoku.\n"},
	}

	for _, c := range cases {
		buf := new(bytes.Buffer)
		cmd := NewCmdShow()
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
