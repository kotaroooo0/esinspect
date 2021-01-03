package cmd

import (
	"encoding/csv"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(crawlCmd)
}

var crawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Collect data to add Elasticsearch",
	Args:  cobra.ExactArgs(1),
	RunE:  crawl,
}

func crawl(cmd *cobra.Command, args []string) error {
	records := [][]string{
		{"name", "description"},
	}
	doc, err := goquery.NewDocument("https://github.com/" + args[0] + "?tab=repositories")
	if err != nil {
		return err
	}

	ul := doc.Find("div#user-repositories-list li")
	ul.Each(func(i int, li *goquery.Selection) {
		name := li.Find("h3.wb-break-all > a").Text()
		name = strings.Replace(name, " ", "", -1)
		name = strings.Replace(name, "\n", "", -1)

		description := li.Find("p").Text()
		description = strings.Replace(description, " ", "", 10)
		description = strings.Replace(description, "\n", "", -1)

		records = append(records, []string{
			name, description,
		})
	})

	w := csv.NewWriter(os.Stdout)
	if err := w.WriteAll(records); err != nil {
		return err
	}
	return nil
}
