package cmd

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

func NewCmdInit() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize your database for task management.",
		Long:  `Initialize your database for task management.`,
		RunE: func(cmd *cobra.Command, args []string) error {

			dbPath := GetDbPath()

			DbConnection, _ := sql.Open("sqlite3", dbPath)

			defer DbConnection.Close()

			sqlCmd := `CREATE TABLE IF NOT EXISTS task(
				id INTEGER PRIMARY KEY,
				title TEXT,
				state INTEGER)`
			_, err := DbConnection.Exec(sqlCmd)

			sqlCmd = `CREATE TABLE IF NOT EXISTS log(
				id INTEGER PRIMARY KEY,
				task_id INTEGER,
				start TEXT,
				end TEXT,
				period INTEGER,
				state INTEGER)`
			_, err = DbConnection.Exec(sqlCmd)

			if err != nil {
				cmd.Printf("failed to init dakoku.\n")
				return err
			}

			cmd.Printf("init dakoku.\n")
			return nil
		},
	}
	return cmd
}

func init() {}
