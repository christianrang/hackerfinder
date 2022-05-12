package cmd

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	outputs "github.com/christianrang/find-bad-ip/internal/outputs/ip"
	"github.com/christianrang/find-bad-ip/pkg/abuseipdbsdk"
	"github.com/christianrang/find-bad-ip/pkg/abuseipdbsdk/check"
	"github.com/christianrang/find-bad-ip/pkg/vtsdk"
	"github.com/christianrang/find-bad-ip/pkg/vtsdk/ipaddress"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var (
	ipOutput outputs.Ip
	ipFile   []string
	ips      []string
	// Used to output data
	csvFilename string
	csvFile     *os.File
	csvWriter   *csv.Writer
	searchCmd   = &cobra.Command{
		Use:   "search [OPTIONS]",
		Short: "searches virustotal and abuseaipdb",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Searching...")
			vtClient := vtsdk.CreateClient(configuration.VTConfig)
			abuseIpClient := abuseipdbsdk.CreateClient(configuration.Abuseipdb)

			t := outputs.InitializeTable()

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
				outputs.WriteRow(csvWriter, outputs.CreateHeaders())
				csvWriter.Flush()
			}

			handleIp(vtClient, abuseIpClient, t, csvWriter)
			handleIpFile(vtClient, abuseIpClient, t, csvWriter)

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

func handleIp(vtClient *vtsdk.Client, abuseIpClient *abuseipdbsdk.Client, t table.Writer, csvW *csv.Writer) {
	for _, ip := range ips {
		_, err := ipaddress.QueryIp(*vtClient, ip, &ipOutput.VtIpAddress)

		_, err = check.QueryCheck(*abuseIpClient, ip, &ipOutput.AbuseipdbCheck)

		if err != nil {
			log.Fatalln(err)
		}

		if csvW != nil {
			outputs.WriteRow(csvW, ipOutput.CreateRecord())
		}

		ipOutput.CreateTableRow(t)
	}
}

func handleIpFile(vtClient *vtsdk.Client, abuseIpClient *abuseipdbsdk.Client, t table.Writer, csvW *csv.Writer) {
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
			_, err := ipaddress.QueryIp(*vtClient, ip, &ipOutput.VtIpAddress)
			if err != nil {
				log.Fatalf("%#v", err)
			}

			_, err = check.QueryCheck(*abuseIpClient, ip, &ipOutput.AbuseipdbCheck)
			if err != nil {
				fmt.Println(err)
			}

			if csvW != nil {
				outputs.WriteRow(csvW, ipOutput.CreateRecord())
			}

			ipOutput.CreateTableRow(t)
		}
	}
}
