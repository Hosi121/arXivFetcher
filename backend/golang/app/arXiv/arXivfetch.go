package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Entry struct {
	Title    string   `xml:"title"`
	Summary  string   `xml:"summary"`
	Authors  []Author `xml:"author"`
}

type Author struct {
	Name string `xml:"name"`
}

func fetchArxivData(arxivID string) *http.Response {
	apiURL := fmt.Sprintf("http://export.arxiv.org/api/query?id_list=%s", arxivID)
	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Failed to fetch data:", err)
		return nil
	}
	return response
}

func parseArxivData(response *http.Response) *Entry {
	if response.StatusCode == 200 {
		var feed struct {
			Entry Entry `xml:"entry"`
		}
		data, _ := ioutil.ReadAll(response.Body)
		xml.Unmarshal(data, &feed)
		return &feed.Entry
	}
	return nil
}

func fetchBibtex(arxivID string) string {
	bibtexURL := fmt.Sprintf("https://arxiv.org/bibtex/%s", arxivID)
	response, err := http.Get(bibtexURL)
	if err != nil || response.StatusCode != 200 {
		return "BibTeX information could not be fetched."
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return "Failed to parse BibTeX."
	}
	bibtex := doc.Find("pre").Text()
	return bibtex
}

func printMarkdown(arxivURL string, entry *Entry, bibtex string) {
	fmt.Println("## 書誌情報")
	fmt.Printf("### タイトル\n%s\n \n", entry.Title)
	fmt.Printf("### URL\n%s\n \n", arxivURL)
	authors := make([]string, len(entry.Authors))
	for i, author := range entry.Authors {
		authors[i] = author.Name
	}
	fmt.Printf("### 著者\n%s\n \n", strings.Join(authors, ", "))
	fmt.Printf("### 概要\n%s\n", entry.Summary)
	fmt.Printf("### BibTeX\n```bibtex\n%s\n```\n", bibtex)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the arXiv link: ")
	arxivLink, _ := reader.ReadString('\n')
	arxivLink = strings.TrimSpace(arxivLink)
	arxivID := strings.Split(arxivLink, "/")[len(strings.Split(arxivLink, "/"))-1]
	response := fetchArxivData(arxivID)
	if response != nil {
		defer response.Body.Close()
		entry := parseArxivData(response)
		if entry != nil {
			bibtex := fetchBibtex(arxivID)
			printMarkdown(arxivLink, entry, bibtex)
		} else {
			fmt.Println("Failed to parse data from arXiv.")
		}
	} else {
		fmt.Println("Failed to fetch data from arXiv API.")
	}
}

