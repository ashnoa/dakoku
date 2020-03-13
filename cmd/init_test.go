package cmd

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestInit(t *testing.T) {
	cases := []struct {
		command string
		want    string
	}{
		{command: "dakoku init", want: "init dakoku.\n"},
	}

	dbPath := GetDbPath()

	for _, c := range cases {
		buf := new(bytes.Buffer)
		cmd := NewCmdInit()
		cmd.SetOutput(buf)
		cmdArgs := strings.Split(c.command, " ")
		fmt.Printf("cmdArgs %+v\n", cmdArgs)
		cmd.SetArgs(cmdArgs[1:])
		cmd.Execute()

		get := buf.String()
		if c.want != get {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
		if !Exists(dbPath) {
			t.Errorf("db.sql is nothing. %v", dbPath)
		}
		os.Remove(dbPath)
	}
}
