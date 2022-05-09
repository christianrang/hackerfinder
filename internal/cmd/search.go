package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/christianrang/find-bad-ip/pkg/vtsdk"
	"github.com/christianrang/find-bad-ip/pkg/vtsdk/ipaddress"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var (
	ipFile    []string
	ips       []string
	searchCmd = &cobra.Command{
		Use:   "search",
		Short: "searches virustotal",
		Run: func(cmd *cobra.Command, args []string) {
			client := vtsdk.CreateClient(configuration.VTConfig)

			t := initializeTable()

			handleIp(client, t)
			handleIpFile(client, t)

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
}

func initializeTable() table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"IP", "Malicious", "Suspicious", "Harmless", "Threat Level"})
	t.AppendSeparator()
	return t
}

func handleIp(client *vtsdk.Client, t table.Writer) {
	for _, ip := range ips {
		_, result, err := ipaddress.QueryIp(*client, ip)
		if err != nil {
			log.Fatalln(err)
		}
		result.Table(t)
	}
}

func handleIpFile(client *vtsdk.Client, t table.Writer) {
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
			result.Table(t)
		}
	}
}
