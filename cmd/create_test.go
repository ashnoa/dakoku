package cmd

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestCreateSuccess(t *testing.T) {
	cases := []struct {
		command string
		want    string
	}{
		{command: "dakoku create task_title", want: "create a task [task_title].\n"},
		{command: "dakoku create task_title task_title_another", want: "create a task [task_title].\n"},
	}

	dbPath := GetDbPath()

	for _, c := range cases {
		InitDbForTest(new(bytes.Buffer))

		buf := new(bytes.Buffer)
		cmd := NewCmdCreate()
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

func TestCreateArgsError(t *testing.T) {
	cases := []struct {
		command string
		want    string
	}{
		{command: "dakoku create",
			want: "failed to create a task. task's title is required."},
	}

	dbPath := GetDbPath()

	for _, c := range cases {
		InitDbForTest(new(bytes.Buffer))

		buf := new(bytes.Buffer)
		cmd := NewCmdCreate()
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

func TestCreateDbNotInitError(t *testing.T) {
	cases := []struct {
		command string
		want    string
	}{
		{command: "dakoku create task_title", want: "failed to create a task. please initialize dakoku.\n"},
	}

	for _, c := range cases {
		buf := new(bytes.Buffer)
		cmd := NewCmdCreate()
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
