package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
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
	var jobs []extractedJob
	mainCh := make(chan []extractedJob)
	pages := getPages()
	for i := 0; i < pages; i++ {
		go getPage(i, mainCh)
	}
	for i := 0; i < pages; i++ {
		extractedJobs := <-mainCh
		jobs = append(jobs, extractedJobs...)
	}
	writeJobs(jobs)
	fmt.Println("Done, jobs extracted. ", len(jobs))
}

func writeJobs(jobs []extractedJob) {
	// jCh := make(chan error)
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"Link", "Title", "Location", "Salary", "Summary"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.salary, job.summary}
		err := w.Write(jobSlice)
		checkErr(err)
		// go writeJob(*w, jobSlice, jCh)
	}
	// for i := 0; i < len(jobs); i++ {
	// 	err := <-jCh
	// 	checkErr(err)
	// }
}

// func writeJob(w csv.Writer, job []string, ch chan<- error) {
// 	err := w.Write(job)
// 	ch <- err
// }

func getPage(page int, mainCh chan<- []extractedJob) {
	ch := make(chan extractedJob)
	var jobs []extractedJob
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting... ", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkStatusCode(res)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	searchCards := doc.Find(".jobsearch-SerpJobCard")
	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, ch)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-ch
		jobs = append(jobs, job)
	}
	mainCh <- jobs
}

func extractJob(card *goquery.Selection, ch chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := cleanText(card.Find(".title>a").Text())
	location := cleanText(card.Find(".sjcl").Text())
	salary := cleanText(card.Find(".salaryText").Text())
	summary := cleanText(card.Find(".summary").Text())
	ch <- extractedJob{
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
