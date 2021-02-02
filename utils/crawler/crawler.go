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

func getPages(baseUrl string, c chan<- data.TodaySoccerSchedule) {
	res, err := http.Get(baseUrl)

	checkErr(err)
	checkStatus(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".월간 일정/결과").Each(func(i int, s *goquery.Selection) {
		c <- data.TodaySoccerSchedule{
			data.GameSchedule{
				Place: s.Find("#_monthlyScheduleList").Find(".time_place").Text()
			}
		}
	})

	return
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
