package cmd

import (
	"encoding/csv"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(crawlCmd)
}

var crawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Crawling data to add Elasticsearch",
	Args:  cobra.NoArgs,
	RunE:  crawl,
}

func crawl(cmd *cobra.Command, args []string) error {
	records := [][]string{
		{"entry_title", "blog_title"},
	}
	doc, err := goquery.NewDocument("https://hatenablog.com/")
	if err != nil {
		return err
	}

	selections := doc.Find("div.serviceTop-recommend-grid")
	selections.Each(func(i int, selection *goquery.Selection) {
		entryTitle := selection.Find("div.serviceTop-entry-title a").Text()
		blogTitle := selection.Find("div.serviceTop-blog-title a").Text()
		records = append(records, []string{
			entryTitle, blogTitle,
		})
	})

	w := csv.NewWriter(os.Stdout)
	if err := w.WriteAll(records); err != nil {
		return err
	}
	return nil
}
