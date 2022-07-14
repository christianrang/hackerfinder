package cmd

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	"github.com/christianrang/hackerfinder/internal"
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

	filterCmd = &cobra.Command{
		Use:   "filter [OPTIONS]",
		Short: "filters out domains, IPs and hashes and searches for them",
		Run: func(cmd *cobra.Command, args []string) {
			if !configuration.Api.HasApiKey() {
				fmt.Printf("error: please configure an API key\n")
				os.Exit(2)

				client := internal.Client{
					VirusTotalClient: vtsdk.CreateClient(configuration.Api.VTConfig),
					AbuseipdbClient:  abuseipdbsdk.CreateClient(configuration.Api.Abuseipdb),
				}

				for _, filename := range filenames {
					contents, err := ioutil.ReadFile(filename)
					if err != nil {
						fmt.Printf("error: opening file: %s\n", err)
						os.Exit(3)
					}

					// TODO: look into the run time of this. If we are looping over the file multiple times it may be worth working this into a single loop
					items := []struct {
						query     internal.QueryFunc
						regex     string
						table     table.Writer
						csvWriter *csv.Writer
					}{
						{
							query: client.QueryDomain,
							regex: commonregex.Domain,
							table: domainTable,
						},
						{
							query: client.QueryIp,
							regex: commonregex.Ip,
							table: ipTable,
						},
						{
							query: client.QueryHashes,
							regex: commonregex.Sha1,
							table: hashesTable,
						},
						{
							query: client.QueryHashes,
							regex: commonregex.Sha256,
							table: hashesTable,
						},
						{
							query: client.QueryHashes,
							regex: commonregex.Md5,
							table: hashesTable,
						},
					}

					for _, item := range items {
						compiledRegex, _ := regexp.Compile(item.regex)
						foundItems := compiledRegex.FindAll(contents, -1)
						handleQuery(client, item.table, item.csvWriter, foundItems, item.query)
					}
				}
			}
		},
	}
)

func handleQuery(client internal.Client, t table.Writer, csvW *csv.Writer, values [][]byte, query internal.QueryFunc) {
	for _, value := range values {

		resp, err := client.Query(string(value), query)

		if err != nil {
			log.Fatalln(err)
		}

		if csvW != nil {
			resp.WriteRow(csvW, resp.CreateRecord)
		}

		resp.CreateTableRow(t)
	}
}
