package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/christianrang/hackerfinder/internal/cmd/helpers/update"
	"github.com/minio/selfupdate"
	"github.com/spf13/cobra"
)

var (
	updateCommand = &cobra.Command{
		Use:   "update",
		Short: "updates hackerfinder",
		Run: func(cmd *cobra.Command, args []string) {
			url, err := update.FindCurrentReleaseUrl()
			if err != nil {
				log.Fatalln("error updating: ", err)
			}
			doUpdate(url)
		},
	}
)

func init() {
	rootCmd.AddCommand(updateCommand)
}

func doUpdate(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = selfupdate.Apply(resp.Body, selfupdate.Options{})
	if err != nil {
		if rollbackErr := selfupdate.RollbackError(err); rollbackErr != nil {
			fmt.Printf("Failed to rollback from bad update: %v\n", rollbackErr)
		}
	}
	return err
}
