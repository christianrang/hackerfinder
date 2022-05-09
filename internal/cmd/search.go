package cmd

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/christianrang/find-bad-ip/pkg/vtsdk"
	"github.com/christianrang/find-bad-ip/pkg/vtsdk/ipaddress"
	"github.com/christianrang/find-bad-ip/pkg/vtsdk/ipaddress/output"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var (
	ipFile []string
	ips    []string
	// Used to output data
	csvFilename string
	csvFile     *os.File
	csvWriter   *csv.Writer
	searchCmd   = &cobra.Command{
		Use:   "search [OPTIONS]",
		Short: "searches virustotal",
		Run: func(cmd *cobra.Command, args []string) {
			client := vtsdk.CreateClient(configuration.VTConfig)

			t := initializeTable()

			if csvFilename != "" {
				if _, err := os.Stat(csvFilename); !errors.Is(err, os.ErrNotExist) {
					fmt.Printf("error: file %s already exists\n", csvFilename)
					os.Exit(1)
				}
				csvFile, err := os.Create(csvFilename)
				defer csvFile.Close()
				if err != nil {
					fmt.Printf("error: failed to create csv file %s: %s\n", csvFilename, err)
					os.Exit(1)
				}

				csvWriter = csv.NewWriter(csvFile)
				defer csvWriter.Flush()
				output.WriteRow(csvWriter, output.CreateHeaders())
				csvWriter.Flush()
			}

			handleIp(client, t, csvWriter)
			handleIpFile(client, t, csvWriter)

			t.Render()
		},
	}
)

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringSliceVar(
		&ipFile,
		"ip-file",
		ipFile,
		// TODO: add a useage
		"sets a file of IPs to search. Each IP must be on its own line.",
	)
	searchCmd.Flags().StringSliceVar(
		&ips,
		"ip",
		ips,
		// TODO: add a useage
		"",
	)
	searchCmd.Flags().StringVar(
		&csvFilename,
		"csv",
		"",
		// TODO: add a useage
		"--csv [output filename]",
	)
}

func initializeTable() table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"IP", "Malicious", "Suspicious", "Harmless", "Threat Level"})
	t.AppendSeparator()
	return t
}

func handleIp(client *vtsdk.Client, t table.Writer, csvW *csv.Writer) {
	for _, ip := range ips {
		_, result, err := ipaddress.QueryIp(*client, ip)

		if err != nil {
			log.Fatalln(err)
		}

		if csvW != nil {
			output.WriteRow(csvW, output.CreateRecord(result))
		}

		result.Table(t)
	}
}

func handleIpFile(client *vtsdk.Client, t table.Writer, csvW *csv.Writer) {
	for _, file := range ipFile {
		data, err := os.ReadFile(file)
		if err != nil {
			log.Fatalf("%#v", err)
		}

		contents := strings.Split(string(data), "\n")

		for _, ip := range contents {
			if ip == "" {
				break
			}
			_, result, err := ipaddress.QueryIp(*client, ip)
			if err != nil {
				log.Fatalf("%#v", err)
			}

			if csvW != nil {
				output.WriteRow(csvW, output.CreateRecord(result))
			}

			result.Table(t)
		}
	}
}
