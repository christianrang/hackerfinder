package cmd

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/christianrang/hackerfinder/internal"
	outputDomain "github.com/christianrang/hackerfinder/internal/outputs/domain"
	outputHashes "github.com/christianrang/hackerfinder/internal/outputs/hashes"
	outputIp "github.com/christianrang/hackerfinder/internal/outputs/ip"
	outputTypes "github.com/christianrang/hackerfinder/internal/outputs/types"
	"github.com/christianrang/hackerfinder/internal/outputs/ui"
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

	results = make([]outputTypes.Output, 0)

	filterCmd = &cobra.Command{
		Use:   "filter [OPTIONS]",
		Short: "filters and searches for domains, IPs and hashes",
		Run: func(cmd *cobra.Command, args []string) {
			go func() {
				if !configuration.Api.HasApiKey() {
					fmt.Printf("error: please configure an API key\n")
					os.Exit(2)
				}

				client := internal.Client{
					VirusTotalClient: vtsdk.CreateClient(configuration.Api.VTConfig),
					AbuseipdbClient:  abuseipdbsdk.CreateClient(configuration.Api.Abuseipdb),
				}

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

					items := []struct {
						query     internal.QueryFunc
						regex     string
						table     table.Writer
						csvWriter *csv.Writer
						validator func(string) bool
					}{
						{
							query:     client.QueryDomain,
							regex:     commonregex.Domain,
							table:     domainTable,
							csvWriter: domainCsvWriter,
							validator: ValidateDomain,
						},
						{
							query:     client.QueryIp,
							regex:     commonregex.Ip,
							table:     ipTable,
							csvWriter: ipCsvWriter,
							validator: ValidateIp,
						},
						{
							query:     client.QueryHashes,
							regex:     commonregex.VirusTotalHashes,
							table:     hashesTable,
							csvWriter: hashesCsvWriter,
							validator: ValidateHash,
						},
					}

					for _, item := range items {
						compiledRegex, _ := regexp.Compile(item.regex)
						foundItems := compiledRegex.FindAll(contents, -1)
						for _, foundItem := range foundItems {
							foundItemString := string(foundItem)
							if item.validator(foundItemString) {
								handleQuery(client, item.table, item.csvWriter, foundItemString, item.query, p)
							}
						}
					}
				}

				p.Quit()
			}()

			if err := p.Start(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			n := tea.NewProgram(ui.InitTableModel(results))
			if err := n.Start(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
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

func handleQuery(client internal.Client, t table.Writer, csvW *csv.Writer, value string, query internal.QueryFunc, p *tea.Program) {
	p.Send(ui.QueryMsg{Target: value})

	resp, err := client.Query(value, query)

	if err != nil {
		log.Fatalln(err)
	}

	if csvW != nil {
		resp.WriteRow(csvW, resp.CreateRecord)
	}

	results = append(results, resp)
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

func ValidateDomain(domain string) bool {
	ipCompiledRegex, _ := regexp.Compile(commonregex.Ip)
	if ipCompiledRegex.MatchString(domain) {
		return false
	}

	return true
}

func ValidateIp(ip string) bool {
	return true
}

func ValidateHash(hash string) bool {
	return true
}
