package main

import (
	"encoding/csv"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	user := flag.String("u", "kotaroooo0", "GitHub user name(required)")
	flag.Parse()
	gitHubScriper := NewGitHubScriper(*user)

	if err := gitHubScriper.scripe(); err != nil {
		log.Fatal(err)
	}
}

type Scriper interface {
	scripe() error
}

// sample scriper
type GitHubScriper struct {
	User string
}

func NewGitHubScriper(user string) GitHubScriper {
	return GitHubScriper{
		User: user,
	}
}

func (s GitHubScriper) scripe() error {
	records := [][]string{
		{"name", "description"},
	}
	doc, err := goquery.NewDocument("https://github.com/" + s.User + "?tab=repositories")
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
