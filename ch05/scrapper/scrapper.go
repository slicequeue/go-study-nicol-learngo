package scrapper

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

var baseURL string = "https://www.saramin.co.kr"

const (
	CSV	= iota
	JSON
)
// TODO 시간 측정 비교
// TODO JSON 쓰기, CSV, JSON 타입 상수 형태
// TODO 정렬해보기 - CSV 순서대로 시행, 고루틴 정렬안하고, 고루팀 정렬 세타입 결과 파일 비교

// Scrape Indeed by a term
func Scrape(term string) {
	var pageSize int = 5
	var targetURL string = "https://www.saramin.co.kr/zf_user/search/recruit?search_area=main&searchword=" + term + "&recruitPageCount=" + strconv.Itoa(pageSize)

	var jobs []extractedJob
	c := make(chan []extractedJob)
	totalPages := getPages(targetURL)
	fmt.Println(totalPages)
	for i := 1; i <= totalPages; i++ {
		go getPage(targetURL, i, c)
	}

	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...) // 여기서도 deconstructor... 를 사용할 줄이야 ㅎㅎ
	}
	writeJobs(jobs)
}

func getPage(url string, page int, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageUrl := url + "&recruitPage=" + strconv.Itoa(page)
	fmt.Println(pageUrl)
	res, err := http.Get(pageUrl)
	checkErr(err)
	checkCode(res)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".item_recruit")
	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})
	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}
	mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("value")
	title := CleanString(card.Find(".job_tit>a").Text())
	location := CleanString(card.Find(".job_condition>span:nth-child(1)").Text())
	deadline := CleanString(card.Find(".job_date>.date").Text())
	summary := CleanString(card.Find(".job_condition>span:not(:nth-child(1))").Text())
	url, _ := card.Find(".job_tit>a").Attr("href")
	// fmt.Println(url)
	c <- extractedJob{
		id:       id,
		title:    title,
		location: location,
		deadline: deadline,
		summary:  summary,
		url:      baseURL + url,
	}
}

// CleanString cleans a string
func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ") // 탁월한 방식! TrimSpace 양쪽 여백 없애고 Fields으로 중간에 공백을 기반으로 나누고 이후 Join 으로 한칸 공백으로 묶기
}

func getPages(url string) int {
	pages := 0
	res, err := http.Get(url)
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

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush() // cool~~

	headers := []string{"ID", "Title", "Location", "Deadline", "Summary", "URL"}
	wErr := w.Write(headers)
	checkErr(wErr)

	// 이 부분도 고루틴으로 만들 수 있음 ㅋㅋ
	for _, job := range jobs {
		jobSlice := []string{job.id, job.title, job.location, job.deadline, job.summary, job.url}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
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
