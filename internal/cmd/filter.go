package cmd

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	"github.com/christianrang/hackerfinder/internal"
	outputDomain "github.com/christianrang/hackerfinder/internal/outputs/domain"
	outputHashes "github.com/christianrang/hackerfinder/internal/outputs/hashes"
	outputIp "github.com/christianrang/hackerfinder/internal/outputs/ip"
	"github.com/christianrang/hackerfinder/pkg/abuseipdbsdk"
	commonregex "github.com/christianrang/hackerfinder/pkg/regex"
	"github.com/christianrang/hackerfinder/pkg/vtsdk"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var (
	filenames []string

	domainTable table.Writer
	ipTable     table.Writer
	hashesTable table.Writer

	domainCsvWriter *csv.Writer
	ipCsvWriter     *csv.Writer
	hashesCsvWriter *csv.Writer

	filterCmd = &cobra.Command{
		Use:   "filter [OPTIONS]",
		Short: "filters and searches for domains, IPs and hashes",
		Run: func(cmd *cobra.Command, args []string) {
			if !configuration.Api.HasApiKey() {
				fmt.Printf("error: please configure an API key\n")
				os.Exit(2)
			}

			client := internal.Client{
				VirusTotalClient: vtsdk.CreateClient(configuration.Api.VTConfig),
				AbuseipdbClient:  abuseipdbsdk.CreateClient(configuration.Api.Abuseipdb),
			}

			domainTable = outputDomain.InitializeTable()
			ipTable = outputIp.InitializeTable()
			hashesTable = outputHashes.InitializeTable()

			if csvFilename != "" {
				// TODO clean up CSV creation and create a csv file for each artifact type
				domainCsvWriter = CreateCsvWriter(csvFilename, ".domain")
				ipCsvWriter = CreateCsvWriter(csvFilename, ".ips")
				hashesCsvWriter = CreateCsvWriter(csvFilename, ".hashes")

				outputDomain.WriteRow(domainCsvWriter, outputDomain.CreateHeaders)
				outputIp.WriteRow(ipCsvWriter, outputIp.CreateHeaders)
				outputHashes.WriteRow(hashesCsvWriter, outputHashes.CreateHeaders)
			}

			for _, filename := range filenames {
				contents, err := ioutil.ReadFile(filename)
				if err != nil {
					fmt.Printf("error: opening file: %s\n", err)
					os.Exit(3)
				}
				fmt.Println("Filtering ", filename)

				items := []struct {
					query     internal.QueryFunc
					regex     string
					table     table.Writer
					csvWriter *csv.Writer
				}{
					{
						query:     client.QueryDomain,
						regex:     commonregex.Domain,
						table:     domainTable,
						csvWriter: domainCsvWriter,
					},
					{
						query:     client.QueryIp,
						regex:     commonregex.Ip,
						table:     ipTable,
						csvWriter: ipCsvWriter,
					},
					{
						query:     client.QueryHashes,
						regex:     commonregex.VirusTotalHashes,
						table:     hashesTable,
						csvWriter: hashesCsvWriter,
					},
				}

				for _, item := range items {
					compiledRegex, _ := regexp.Compile(item.regex)
					foundItems := compiledRegex.FindAll(contents, -1)
					for _, foundItem := range foundItems {
						handleQuery(client, item.table, item.csvWriter, string(foundItem), item.query)
					}
				}
			}

			domainTable.Render()
			ipTable.Render()
			hashesTable.Render()
		},
	}
)

func init() {
	rootCmd.AddCommand(filterCmd)
	filterCmd.Flags().StringSliceVar(
		&filenames,
		"file",
		filenames,
		"specifies a file to filter hashes, IPs, and domains from",
	)

	filterCmd.Flags().StringVar(
		&csvFilename,
		"csv",
		"",
		"--csv [output filename]",
	)
}

func handleQuery(client internal.Client, t table.Writer, csvW *csv.Writer, value string, query internal.QueryFunc) {
	resp, err := client.Query(value, query)

	if err != nil {
		log.Fatalln(err)
	}

	if csvW != nil {
		resp.WriteRow(csvW, resp.CreateRecord)
	}

	resp.CreateTableRow(t)
}

func CreateCsvWriter(filename string, suffix string) *csv.Writer {
	if _, err := os.Stat(filename); !errors.Is(err, os.ErrNotExist) {
		fmt.Printf("error: file %s already exists. Please use a different filename.\n", csvFilename)
		os.Exit(1)
	}
	writer, err := os.Create(filename + suffix)
	defer csvFile.Close()
	if err != nil {
		fmt.Printf("error: failed to create csv file %s: %s\n", csvFilename, err)
		os.Exit(1)
	}

	return csv.NewWriter(writer)
}
