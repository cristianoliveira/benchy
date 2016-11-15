package main

import (
	"encoding/json"
	"fmt"
	elastic "gopkg.in/olivere/elastic.v3"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Sitemap struct {
	Url       string
	UrlMobile string
	Site      string
	Issued    string
	Title     string
}

func render(w http.ResponseWriter, templatePath string, sr *elastic.SearchResult) {
	fmt.Printf("rendering %s \n", templatePath)
	cwd, _ := os.Getwd()

	s := strings.Split(templatePath, "/")
	name := s[len(s)-1]

	t, err := template.New(name).
		ParseFiles(filepath.Join(cwd, templatePath))
	if err != nil {
		panic(err)
	}

	var sitemaps []Sitemap
	for _, hit := range sr.Hits.Hits {
		var sitemap Sitemap
		err = json.Unmarshal(*hit.Source, &sitemap)
		sitemaps = append(sitemaps, sitemap)
	}

	err = t.Execute(w, sitemaps)
	if err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	client, err := elastic.NewClient(
		elastic.SetURL("http://local.elasticsearch.com"),
		elastic.SetSniff(false),
		elastic.SetMaxRetries(5),
	)
	if err != nil {
		panic(err)
	}

	search, err := client.Search().
		Index("sitemap_g1").
		Do()
	if err != nil {
		panic(err)
	}

	render(w, "templates/index.xml", search)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
