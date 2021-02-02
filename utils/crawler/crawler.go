package crawler

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"smooth-bear.live/utils/data"
)

func CrawlingTodayData() {
	const baseUrl = "https://sports.news.naver.com/wfootball/schedule/index.nhn"

	c := make(chan data.TodaySoccerSchedule)

	go getPages(baseUrl, c)
}

func getPages(baseUrl string, c chan<- data.TodaySoccerSchedule) int {
	res, err := http.Get(baseUrl)

	checkErr(err)
	checkStatus(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkStatus(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request Processing Failed. status: ", res.StatusCode)
	}
}
