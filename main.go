package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

type SoccerDay struct {
	Day string `json:"day"`
	SoccerSchedule []SoccerSchedule `json:"soccer-schedules"`
}

type SoccerSchedule struct {
	Time string `json:"time"`
	Place string `json:"place"`
	LeftTeam string `json:"left-team"`
	RightTeam string `json:"right-team"`
	LeftScore string `json:"left-score"`
	RightScore string `json:"right-score"`
}

const URL = "https://sports.news.naver.com/wfootball/schedule/index.nhn"

func main() {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var soccerDay []SoccerDay = make([]SoccerDay, 31)
	var dayIndex = 0
	var index = 0

	for {
		var day string
		var time string
		var place string
		var leftTeam string
		var rightTeam string
		var leftScore string
		var rightScore string

		err := chromedp.Run(ctx,
			chromedp.Navigate(URL),
			chromedp.WaitVisible("#wrap"),
			chromedp.Text(fmt.Sprintf("#_monthlyScheduleList > tr:nth-child(%d) > th > div > em", index), &day),
			chromedp.Text(fmt.Sprintf("#_monthlyScheduleList > tr:nth-child(%d) > td.time_place > div > span.time", index), &time),
			chromedp.Text(fmt.Sprintf("#_monthlyScheduleList > tr:nth-child(%d) > td.time_place > div > span.place", index), &place),
			chromedp.Text(fmt.Sprintf("#_monthlyScheduleList > tr:nth-child(%d) > td:nth-child(3) > div > span.team_left > span.name", index), &leftTeam),
			chromedp.Text(fmt.Sprintf("#_monthlyScheduleList > tr:nth-child(%d) > td:nth-child(3) > div > span.team_left > span.score", index), &leftScore),
			chromedp.Text(fmt.Sprintf("#_monthlyScheduleList > tr:nth-child(%d) > td:nth-child(3) > div > span.team_right > span.name", index), &leftScore),
			chromedp.Text(fmt.Sprintf("#_monthlyScheduleList > tr:nth-child(%d) > td:nth-child(3) > div > span.team_right > span.score", index), &leftScore),
		)

		if err != nil {
			log.Print(err)
		}

		soccerSchedule := SoccerSchedule{
			Time: time,
			Place: place,
			LeftTeam: leftTeam,
			RightTeam: rightTeam,
			LeftScore: leftScore,
			RightScore: rightScore,
		}

		log.Print(soccerSchedule)

		if index > 0 && day == "" {
			soccerDay[index].SoccerSchedule = append(soccerDay[index].SoccerSchedule, soccerSchedule)
		} else {
			soccerDay = append(soccerDay, SoccerDay{Day: day})
			index++
			soccerDay[index].SoccerSchedule = append(soccerDay[index].SoccerSchedule, soccerSchedule)
		}

		if dayIndex > 30 {
			break
		}
		dayIndex++
	}

	log.Print(soccerDay)
}