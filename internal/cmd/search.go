package cmd

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/christianrang/hackerfinder/internal"
	outputDomain "github.com/christianrang/hackerfinder/internal/outputs/domain"
	outputIp "github.com/christianrang/hackerfinder/internal/outputs/ip"
	"github.com/christianrang/hackerfinder/pkg/abuseipdbsdk"
	"github.com/christianrang/hackerfinder/pkg/vtsdk"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var (
	ipFile     []string
	domainFile []string
	domains    []string
	ips        []string
	// Used to output data
	csvFilename string
	csvFile     *os.File
	csvWriter   *csv.Writer
	searchCmd   = &cobra.Command{
		Use:   "search [OPTIONS]",
		Short: "searches virustotal and abuseaipdb",
		Run: func(cmd *cobra.Command, args []string) {
			var t table.Writer

			client := internal.Client{
				VirusTotalClient: vtsdk.CreateClient(configuration.Api.VTConfig),
				AbuseipdbClient:  abuseipdbsdk.CreateClient(configuration.Api.Abuseipdb),
			}

			if !configuration.Api.HasApiKey() {
				fmt.Printf("error: please configure an API key\n")
				os.Exit(2)
			}

			fmt.Println("Searching...")

			switch {
			case len(domains) > 0:
				t = outputDomain.InitializeTable()
			case len(ips) > 0:
				t = outputIp.InitializeTable()
			default:
				t = outputIp.InitializeTable()
			}

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
				outputIp.WriteRow(csvWriter, outputIp.CreateHeaders())
			}

			handleIp(client, t, csvWriter)
			handleIpFile(client, t, csvWriter)

			handleDomain(client, t, csvWriter)
			handleDomainFile(client, t, csvWriter)

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
		&domains,
		"domain",
		domainFile,
		// TODO: add a useage
		"sets a file of Domains to search. Each Domain must be on its own line.",
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

func handleIp(client internal.Client, t table.Writer, csvW *csv.Writer) {
	for _, ip := range ips {

		resp, err := client.QueryIp(ip)

		if err != nil {
			log.Fatalln(err)
		}

		if csvW != nil {
			outputIp.WriteRow(csvW, resp.CreateRecord())
		}

		resp.CreateTableRow(t)
	}
}

func handleDomain(client internal.Client, t table.Writer, csvW *csv.Writer) {
	for _, domain := range domains {

		resp, err := client.QueryDomain(domain)

		if err != nil {
			log.Fatalln(err)
		}

		if csvW != nil {
			outputIp.WriteRow(csvW, resp.CreateRecord())
		}

		resp.CreateTableRow(t)
	}
}

func handleIpFile(client internal.Client, t table.Writer, csvW *csv.Writer) {
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
			resp, err := client.QueryIp(ip)
			if err != nil {
				fmt.Printf("error: failed to query ip: %s", err)
			}

			if csvW != nil {
				outputIp.WriteRow(csvW, resp.CreateRecord())
			}

			resp.CreateTableRow(t)
		}
	}
}

func handleDomainFile(client internal.Client, t table.Writer, csvW *csv.Writer) {
	for _, file := range domainFile {
		data, err := os.ReadFile(file)
		if err != nil {
			log.Fatalf("%#v", err)
		}

		contents := strings.Split(string(data), "\n")

		for _, domain := range contents {
			if domain == "" {
				break
			}
			resp, err := client.QueryDomain(domain)
			if err != nil {
				fmt.Printf("error: failed to query domain: %s", err)
			}

			if csvW != nil {
				outputIp.WriteRow(csvW, resp.CreateRecord())
			}

			resp.CreateTableRow(t)
		}
	}
}
