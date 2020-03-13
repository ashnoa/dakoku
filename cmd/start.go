package cmd

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"time"
)

func NewCmdStart() *cobra.Command {
	// startCmd represents the start command
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start a task.",
		Long:  "start a task.",
		RunE: func(cmd *cobra.Command, args []string) error {
			dbPath := GetDbPath()
			if !Exists(dbPath) {
				return errors.New("failed to start a task. please initialize dakoku.\n")
			}
			if len(args) < 1 {
				return errors.New("failed to start a task. task's id is required.")
			}

			DbConnection, _ := sql.Open("sqlite3", dbPath)

			defer DbConnection.Close()

			sqlCmd := `SELECT id, title FROM task WHERE id = ? AND state = 0`
			row := DbConnection.QueryRow(sqlCmd, args[0])
			var id int
			var title string
			err := row.Scan(&id, &title)
			if err != nil {
				if err == sql.ErrNoRows {
					return errors.New("failed to start tasks. please check the task exists and isn't working.")
				}
				return errors.New("failed to start a task.")
			}

			now := time.Now().Format(time.RFC3339)
			sqlCmd = `INSERT INTO log (task_id, start, end, period, state)
			VALUES (?, ?, ?, 0, ?)`
			_, err = DbConnection.Exec(sqlCmd, id, now, now, On)
			if err != nil {
				return errors.New("failed to start a task.")
			}

			sqlCmd = `UPDATE task SET state = ? WHERE id = ?`
			_, err = DbConnection.Exec(sqlCmd, On, id)
			if err != nil {
				return errors.New("failed to start a task.")
			}

			cmd.Printf("start task %v[%v].\n", id, title)
			return nil
		},
	}
	return cmd
}

func init() {}
