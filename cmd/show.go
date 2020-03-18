package cmd

import (
	"database/sql"
	"errors"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"log"
	"time"
)

var (
	o = &Options{}
)

func NewCmdShow() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "show your tasks.",
		Long:  `Show your tasks.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			dbPath := GetDbPath()
			if !Exists(dbPath) {
				return errors.New("failed to show tasks. please initialize dakoku.\n")
			}
			DbConnection, _ := sql.Open("sqlite3", dbPath)
			defer DbConnection.Close()

			var (
				sqlCmd string
				rows   *sql.Rows
			)
			if o.isAll {
				sqlCmd = `SELECT task.id, task.title, task.state,
						log.id, log.task_id, log.start, log.end, log.period, sum(log.period), log.state
						FROM task
						LEFT OUTER JOIN log ON log.task_id = task.id
						WHERE task.state <> -1
						GROUP BY task.id`
				rows, _ = DbConnection.Query(sqlCmd)
			} else {
				t := time.Now()
				until := t.Truncate(time.Hour).Add(-time.Duration(t.Hour()) * time.Hour).Add(-time.Duration(time.Hour) * time.Duration(24*o.days)).Format(time.RFC3339)
				sqlCmd = `SELECT task.id, task.title, task.state,
						log.id, log.task_id, log.start, log.end, log.period, sum(log.period), log.state
						FROM task
						LEFT OUTER JOIN log ON log.task_id = task.id
						WHERE task.state <> -1 AND log.end >= ?
						GROUP BY task.id`
				rows, _ = DbConnection.Query(sqlCmd, until)
			}

			defer rows.Close()

			var results []Result
			for rows.Next() {
				var (
					t           Task
					l           Log
					start       string
					end         string
					period      int
					sumOfPeriod int
				)

				err := rows.Scan(&t.id, &t.title, &t.state, &l.id, &l.taskId, &start, &end, &period, &sumOfPeriod, &l.state)

				l.start, _ = time.Parse(time.RFC3339, start)
				l.end, _ = time.Parse(time.RFC3339, end)
				l.period = time.Duration(period)

				if err != nil {
					log.Fatalln(err)
					return errors.New("failed to show tasks.\n")
				}

				var stateView string
				if t.state == 1 {
					stateView = "Now Doing"
				} else {
					stateView = ""
				}

				r := Result{task: &t, log: &l, sumOfPeriod: time.Duration(sumOfPeriod), stateView: stateView}
				results = append(results, r)
			}

			for _, r := range results {
				cmd.Printf("%-v | %-v | %-v | %-v\n", r.task.id, r.task.title, r.sumOfPeriod, aurora.BrightGreen(r.stateView))
			}
			return nil
		},
	}
	cmd.Flags().IntVarP(&o.days, "days", "d", 0, "Show tasks and work times for input days.")
	cmd.Flags().BoolVarP(&o.isAll, "all", "a", false, "Show all tasks and work times.")
	return cmd
}

func init() {}
