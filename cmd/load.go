package cmd

import (
	"context"
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"github.com/olivere/elastic"
	"github.com/spf13/cobra"
)

var (
	path  string
	index string
)

func init() {
	loadCmd.Flags().StringVarP(&path, "path", "p", "", "data source path (required)")
	loadCmd.MarkFlagRequired("path")
	loadCmd.Flags().StringVarP(&index, "index", "i", "", "Elasticsearch index for data insertion (required)")
	loadCmd.MarkFlagRequired("index")
	rootCmd.AddCommand(loadCmd)
}

var loadCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Collect data to add Elasticsearch",
	Args:  cobra.NoArgs,
	RunE:  load,
}

func load(cmd *cobra.Command, args []string) error {
	// データの格納先
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	// Elasticsearchクライアント
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://localhost:9200"),
	)
	bulk := client.Bulk()

	// CSVリーダー
	r := csv.NewReader(f)
	columns, err := r.Read()
	if err != nil {
		return err
	}
	id := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		doc := make(map[string]string)
		for i, c := range columns {
			doc[c] = record[i]
		}
		bulk.Add(elastic.NewBulkUpdateRequest().Index(index).Id(strconv.Itoa(id)).Doc(doc).DocAsUpsert(true))
		id++
	}
	_, err = bulk.Do(context.Background())
	return err
}
