package cmd

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/christianrang/hackerfinder/internal"
	hferrors "github.com/christianrang/hackerfinder/internal/errors"
	outputDomain "github.com/christianrang/hackerfinder/internal/outputs/domain"
	outputHashes "github.com/christianrang/hackerfinder/internal/outputs/hashes"
	outputIp "github.com/christianrang/hackerfinder/internal/outputs/ip"
	"github.com/christianrang/hackerfinder/internal/outputs/ui"
	"github.com/christianrang/hackerfinder/pkg/abuseipdbsdk"
	"github.com/christianrang/hackerfinder/pkg/vtsdk"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var (
	ipFile     []string
	domainFile []string
	hashesFile []string

	domains []string
	ips     []string
	hashes  []string

	// Used to output data
	csvFilename string
	csvFile     *os.File
	csvWriter   *csv.Writer

	p = tea.NewProgram(ui.InitialModel())

	searchCmd = &cobra.Command{
		Use:   "search [OPTIONS]",
		Short: "searches virustotal and abuseaipdb",
		Run: func(cmd *cobra.Command, args []string) {

			go func() {
				var t table.Writer

				client := internal.Client{
					VirusTotalClient: vtsdk.CreateClient(configuration.Api.VTConfig),
					AbuseipdbClient:  abuseipdbsdk.CreateClient(configuration.Api.Abuseipdb),
				}

				if !configuration.Api.HasApiKey() {
					fmt.Println(hferrors.ErrNoAPIKeyFound)
					os.Exit(2)
				}

				fmt.Println("Searching...")

				switch {
				case len(domains) > 0:
					t = outputDomain.InitializeTable()
				case len(ips) > 0:
					t = outputIp.InitializeTable()
				case len(hashes) > 0:
					t = outputHashes.InitializeTable()
				}

				if csvFilename != "" {
					if _, err := os.Stat(csvFilename); !errors.Is(err, os.ErrNotExist) {
						fmt.Println(hferrors.NewFileExistsError(csvFilename))
						os.Exit(1)
					}
					csvFile, err := os.Create(csvFilename)
					defer csvFile.Close()
					if err != nil {
						fmt.Println(hferrors.NewFailedFileCreationError(csvFilename).Wrap(err))
						os.Exit(1)
					}

					csvWriter = csv.NewWriter(csvFile)
					defer csvWriter.Flush()
					switch {
					case len(domains) > 0:
						outputDomain.WriteRow(csvWriter, outputDomain.CreateHeaders)
					case len(ips) > 0:
						outputIp.WriteRow(csvWriter, outputIp.CreateHeaders)
					case len(hashes) > 0:
						outputHashes.WriteRow(csvWriter, outputHashes.CreateHeaders)
					}
				}

				handleIp(client, t, csvWriter)
				handleIpFile(client, t, csvWriter)

				handleDomain(client, t, csvWriter)
				handleDomainFile(client, t, csvWriter)

				handleHashes(client, t, csvWriter)
				handleHashesFile(client, t, csvWriter)

				p.Quit()
				t.Render()
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

	searchCmd.Flags().StringSliceVar(
		&hashes,
		"hashes",
		hashes,
		// TODO: add a useage
		"",
	)

	searchCmd.Flags().StringSliceVar(
		&hashesFile,
		"hashes-file",
		hashesFile,
		// TODO: add a useage
		"sets a file of hashes to search. Each hash must be on its own line.",
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
		handleQuery(client, t, csvWriter, ip, client.QueryIp, p)
	}
}

func handleDomain(client internal.Client, t table.Writer, csvW *csv.Writer) {
	for _, domain := range domains {
		handleQuery(client, t, csvWriter, domain, client.QueryDomain, p)
	}
}

func handleHashes(client internal.Client, t table.Writer, csvW *csv.Writer) {
	for _, hash := range hashes {
		handleQuery(client, t, csvWriter, hash, client.QueryHashes, p)
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
			handleQuery(client, t, csvWriter, ip, client.QueryIp, p)
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
			handleQuery(client, t, csvWriter, domain, client.QueryIp, p)
		}
	}
}

func handleHashesFile(client internal.Client, t table.Writer, csvW *csv.Writer) {
	for _, file := range hashesFile {
		data, err := os.ReadFile(file)
		if err != nil {
			log.Fatalf("%#v", err)
		}

		contents := strings.Split(string(data), "\n")

		for _, hash := range contents {
			if hash == "" {
				break
			}
			handleQuery(client, t, csvWriter, hash, client.QueryIp, p)
		}
	}
}
