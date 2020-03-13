package cmd

import (
	"bytes"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"strings"
)

func InitDbForTest(buf *bytes.Buffer) *bytes.Buffer {
	cmd := NewCmdInit()
	cmd.SetOutput(buf)
	cmdArgs := strings.Split("dakoku init", " ")
	fmt.Printf("cmdArgs %+v\n", cmdArgs)
	cmd.SetArgs(cmdArgs[1:])
	cmd.Execute()
	return buf
}

func CreateTaskForTest(buf *bytes.Buffer) *bytes.Buffer {
	cmd := NewCmdCreate()
	cmd.SetOutput(buf)
	cmdArgs := strings.Split("dakoku create task_title", " ")
	fmt.Printf("cmdArgs %+v\n", cmdArgs)
	cmd.SetArgs(cmdArgs[2:])
	cmd.Execute()
	return buf
}

func StartTaskForTest(buf *bytes.Buffer) *bytes.Buffer {
	cmd := NewCmdStart()
	cmd.SetOutput(buf)
	cmdArgs := strings.Split("dakoku start 1", " ")
	fmt.Printf("cmdArgs %+v\n", cmdArgs)
	cmd.SetArgs(cmdArgs[2:])
	cmd.Execute()
	return buf
}
