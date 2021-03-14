package main

import (
	"log"

	"github.com/sclevine/agouti"
)

const scheduleUrl = "https://www.foxsports.com/scores/soccer/"

func main() {
	driver := agouti.ChromeDriver(agouti.ChromeOptions("args", []string{"--headless", "--disable-gpu", "--no-sandbox"}), agouti.Debug)

	err := driver.Start()
	checkErr(err)

	page, err := driver.NewPage()
	checkErr(err)

	err = page.Navigate(scheduleUrl)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(page.FindByID("sched-container").Text())

	defer driver.Stop()

	return
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
