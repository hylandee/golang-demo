package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gocolly/colly"
	"os"
	"strings"
	"time"
)

func main() {
	days := getDays()
	time.Sleep(5000)
	fmt.Println(days)
	fName := "data.csv"
	file, err := os.Create(fName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for i := range days {
		writer := csv.NewWriter(file)
		defer writer.Flush()
		holiday_url := fmt.Sprintf("http://nationaltoday.com/%s/", days[i])
		fmt.Println(holiday_url)
		collector := colly.NewCollector()
		collector.OnHTML("h3.holiday-title", func(e *colly.HTMLElement) {
			holiday := e.Text
			fmt.Println(holiday)
			writer.Write([]string{holiday})
		})
		collector.Visit("https://nationaltoday.com/what-is-today/")
	}
}

func getDays() []string {
	startDate := time.Date(time.Now().Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(time.Now().Year(), time.December, 31, 0, 0, 0, 0, time.UTC)
	var days []string
	for currentDate := startDate; currentDate.Before(endDate) || currentDate.Equal(endDate); currentDate = currentDate.AddDate(0, 0, 1) {
		day := fmt.Sprintf("%s-%d", strings.ToLower(currentDate.Month().String()), currentDate.Day())
		days = append(days, day)
	}
	return days
}
