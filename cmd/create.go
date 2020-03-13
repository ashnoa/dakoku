package cmd

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"time"
)

func NewCmdCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "create a task.",
		Long:  `Create a task. You should initialize database by init before execution of this command.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			dbPath := GetDbPath()
			if !Exists(dbPath) {
				return errors.New("failed to create a task. please initialize dakoku.\n")
			}
			if len(args) < 1 {
				return errors.New("failed to create a task. task's title is required.")
			}

			DbConnection, _ := sql.Open("sqlite3", dbPath)
			defer DbConnection.Close()

			sqlCmd := `INSERT INTO task (title, state) VALUES (?, ?)`
			_, err := DbConnection.Exec(sqlCmd, args[0], Off)
			if err != nil {
				return errors.New("failed to create a task.")
			}

			sqlCmd = `SELECT id FROM task WHERE title = ? AND state = 0 ORDER BY id DESC LIMIT 1`
			row := DbConnection.QueryRow(sqlCmd, args[0])
			var id int
			err = row.Scan(&id)
			if err != nil {
				return errors.New("failed to create a task.")
			}

			now := time.Now().Format(time.RFC3339)
			sqlCmd = `INSERT INTO log (task_id, start, end, period, state)
			VALUES (?, ?, ?, 0, 0)`
			_, err = DbConnection.Exec(sqlCmd, id, now, now)
			if err != nil {
				return errors.New("failed to create a task.")
			}

			cmd.Printf("create a task [%v].\n", args[0])
			return nil
		},
	}
	return cmd
}

func init() {}
