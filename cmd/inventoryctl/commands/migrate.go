package commands

import (
	"fmt"
	"time"

	"github.com/open-farms/inventory/internal/settings"
)

type MigrateCmd struct {
	Dry bool `name:"dry" help:"Print out changes before applying"`
}

var defaultDelay = 500 * time.Millisecond

func timetrack(t time.Duration) (*time.Ticker, chan bool) {
	start := time.Now()
	ticker := time.NewTicker(t)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				fmt.Printf("completed in %s\n", time.Since(start).Round(time.Millisecond))
				return
			case <-ticker.C:
				fmt.Printf("time elapsed %s\n", time.Since(start).Round(time.Millisecond))
			}
		}
	}()

	return ticker, done
}

func (m *MigrateCmd) Run(globals *Globals) error {
	if m.Dry {
		fmt.Printf("\n%s\n\n", "Dry run enabled - no changes will be committed.")
	}

	t, done := timetrack(defaultDelay)
	err := settings.Migrate(globals.Config, m.Dry)
	if err != nil {
		return err
	}

	done <- true
	time.Sleep(defaultDelay)
	t.Stop()

	return nil
}
