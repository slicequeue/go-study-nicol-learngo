package main2

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
	deadline string
	summary  string
	url      string
}

var pageSize int = 1
var baseURL string = "https://www.saramin.co.kr"
var targetURL string = "https://www.saramin.co.kr/zf_user/search/recruit?search_area=main&searchword=python&recruitPageCount=5"

func main_hold() {
	var jobs []extractedJob
	totalPages := getPages()
	fmt.Println(totalPages)
	for i := 1; i <= totalPages; i++ {
		extractedJobs := getPage(i)
		jobs = append(jobs, extractedJobs...) // 여기서도 deconstructor... 를 사용할 줄이야 ㅎㅎ
	}
	for _, job := range jobs {
		fmt.Println(job)
	}
	// fmt.Println(jobs)
	writeJobs(jobs)
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w :=  csv.NewWriter(file)
	defer w.Flush() // cool~~

	headers := []string{"ID", "Title", "Location", "Deadline", "Summary", "URL"}
	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{job.id, job.title, job.location, job.deadline, job.summary, job.url}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func getPage(page int) []extractedJob {
	var jobs []extractedJob
	pageUrl := targetURL + "&recruitPage=" + strconv.Itoa(page)
	fmt.Println(pageUrl)
	res, err := http.Get(pageUrl)
	checkErr(err)
	checkCode(res)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".item_recruit")
	searchCards.Each(func(i int, card *goquery.Selection) {
		job := extractJob(card)
		jobs = append(jobs, *job)
	})
	return jobs
}

func extractJob(card *goquery.Selection) extractedJob {
	id, _ := card.Attr("value")
	// fmt.Println(id)
	title := cleanString(card.Find(".job_tit>a").Text())
	// fmt.Println(title)
	location := cleanString(card.Find(".job_condition>span:nth-child(1)").Text())
	// fmt.Println(location)
	deadline := cleanString(card.Find(".job_date>.date").Text())
	// fmt.Println(deadline)
	summary := cleanString(card.Find(".job_condition>span:not(:nth-child(1))").Text())
	// fmt.Println(summary)
	url, _ := card.Find(".job_tit>a").Attr("href")
	// fmt.Println(url)
	return extractedJob{
		id: id,
		title: title,
		location: location,
		deadline: deadline,
		summary: summary,
		url: baseURL + url,
	}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ") // 탁월한 방식! TrimSpace 양쪽 여백 없애고 Fields으로 중간에 공백을 기반으로 나누고 이후 Join 으로 한칸 공백으로 묶기
}

func getPages() int {
	pages := 0
	res, err := http.Get(targetURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length() //
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status", res.StatusCode)
	}
}
