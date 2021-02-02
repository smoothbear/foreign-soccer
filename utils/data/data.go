package data

type TodaySoccerSchedule []GameSchedule

type GameSchedule struct {
	Date      string `json:"date"`
	Time      string `json:"time"`
	Place     string `json:"place"`
	LeftTeam  string `json:"left-team"`
	RightTeam string `json:"right-team"`
	Score     string `json:"score"`
}

type Data interface {
	GetTodaySoccerSchedule() (TodaySoccerSchedule, error)
}
