package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	pages := getPages()
	for i := 0; i < pages; i++ {
		getPage(i)
	}
}

func getPage(page int) {
	var jobs []extractedJob
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkStatusCode(res)
	// jobsearch-SerpJobCard
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	searchCards := doc.Find(".jobsearch-SerpJobCard").Each(func(i int, card *goquery.Selection) {
		job := extractJob(card)
		jobs = append(jobs, job)
	})
	fmt.Println(searchCards)
}

func extractJob(card *goquery.Selection) extractedJob {
	id, _ := card.Attr("data-jk")
	title := cleanText(card.Find(".title>a").Text())
	location := cleanText(card.Find(".sjcl").Text())
	salary := cleanText(card.Find(".salaryText").Text())
	summary := cleanText(card.Find(".summary").Text())
	return extractedJob{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary,
	}
}

func getPages() int {
	res, err := http.Get(baseURL)
	pages := 0
	checkErr(err)
	checkStatusCode(res)

	// prevent memory leak
	defer res.Body.Close()
	doc, err2 := goquery.NewDocumentFromReader(res.Body)
	checkErr(err2)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})
	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkStatusCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with status code: ", res.StatusCode)
	}
}

func cleanText(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
