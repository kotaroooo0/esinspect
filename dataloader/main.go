package main

import (
	"context"
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/olivere/elastic"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// コマンドライン引数を取得
	path := flag.String("f", "elasticsearch/data.csv", "csv file path (required)")
	index := flag.String("i", "sample", "target index (require)")
	flag.Parse()

	// データの格納先
	f, err := os.Open(*path)
	if err != nil {
		log.Fatal(err)
	}

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
		bulk.Add(elastic.NewBulkUpdateRequest().Index(*index).Id(strconv.Itoa(id)).Doc(doc).DocAsUpsert(true))
		id++
	}
	_, err = bulk.Do(context.Background())
	return err
}
