package cmd

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"time"
)

func NewCmdStop() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stop",
		Short: "stop a task.",
		Long:  "stop a task.",
		RunE: func(cmd *cobra.Command, args []string) error {
			dbPath := GetDbPath()
			if !Exists(dbPath) {
				return errors.New("failed to stop a task. please initialize dakoku.\n")
			}
			if len(args) < 1 {
				return errors.New("failed to stop a task. task's id is required.")
			}

			DbConnection, _ := sql.Open("sqlite3", dbPath)

			defer DbConnection.Close()

			sqlCmd := `SELECT log.id, log.start, task.title
			FROM log
			LEFT OUTER JOIN task ON log.task_id = task.id
			WHERE log.task_id = ? AND log.state = 1
			ORDER BY log.start DESC LIMIT 1`
			row := DbConnection.QueryRow(sqlCmd, args[0])
			var id int
			var start string
			var title string
			err := row.Scan(&id, &start, &title)
			if err != nil {
				if err == sql.ErrNoRows {
					return errors.New("failed to stop tasks. please check the task exists and is working.")
				}
				return errors.New("failed to stop a task.1")
			}

			now := time.Now()
			var startDate time.Time
			startDate, _ = time.Parse(time.RFC3339, start)
			period := now.Sub(startDate).Nanoseconds()

			sqlCmd = `UPDATE log SET end = ?, period = ?, state = ? WHERE id = ?`
			_, err = DbConnection.Exec(sqlCmd, now.Format(time.RFC3339), period, Off, id)
			if err != nil {
				return errors.New("failed to stop a task.2")
			}

			sqlCmd = `UPDATE task SET state = ? WHERE id = ?`
			_, err = DbConnection.Exec(sqlCmd, Off, args[0])
			if err != nil {
				return errors.New("failed to stop a task.3")
			}

			cmd.Printf("stop task %v[%v].\n", args[0], title)
			return nil
		},
	}
	return cmd
}

func init() {}
