// package bbparser
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"

	"totolab/internal/structures"
)

const (
	contentURL = "https://totobrief.ru/results/baltbet/10377"
)

func main() {
	var tableDataSlice []structures.ContentData
	c := colly.NewCollector()

	c.OnHTML("div.table-responsive table tbody", func(e *colly.HTMLElement) {
		newEntry := getContentText(e)
		tableDataSlice = append(tableDataSlice, newEntry)

		findElement := e.DOM.Find("tr.line")
		if len(findElement.Text()) != 0 {
			findElement.Each(func(i int, s *goquery.Selection) {
				extractedMatchInfo := extractMatchInfo(s)
				fmt.Printf("Match number %d. %v\n", i+1, extractedMatchInfo)
			})
		} else {
			extractedDrawInfo := extractDrawInfo(tableDataSlice[0])
			fmt.Printf("drawInfo: %v\n", extractedDrawInfo)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(contentURL)
}

func extractMatchInfo(s *goquery.Selection) structures.MatchInfo {
	matchNumberText := s.Find("td.rw").First().Text()
	matchNumberStr := strings.TrimSpace(matchNumberText)
	matchNumber, _ := strconv.Atoi(matchNumberStr)

	matchTeamsText := s.Find("td div.text-nowrap").Text()
	matchTeamsStr := strings.TrimSpace(matchTeamsText)
	matchTeams := strings.Split(matchTeamsStr, " — ")

	eventText := s.Find("td small").Text()
	event := strings.TrimSpace(eventText)

	p1Text := s.Find("td.p.p1").Text()
	p1Values := strings.Fields(strings.TrimSpace(p1Text))

	pxText := s.Find("td.p.px").Text()
	pxValues := strings.Fields(strings.TrimSpace(pxText))

	p2Text := s.Find("td.p.p2").Text()
	p2Values := strings.Fields(strings.TrimSpace(p2Text))

	way1Book, _ := strconv.ParseFloat(p1Values[1], 32)
	wayXBook, _ := strconv.ParseFloat(pxValues[1], 32)
	way2Book, _ := strconv.ParseFloat(p2Values[1], 32)
	way1Pool, _ := strconv.ParseFloat(p1Values[0], 32)
	wayXPool, _ := strconv.ParseFloat(pxValues[0], 32)
	way2Pool, _ := strconv.ParseFloat(p2Values[0], 32)

	scoreText := s.Find("td").Eq(6).Text()
	score := strings.TrimSpace(scoreText)

	wayText := s.Find("td").Last().Text()
	way := strings.TrimSpace(wayText)

	matchInfo := structures.MatchInfo{
		MatchNumber: matchNumber,
		HomeTeam:    matchTeams[0],
		AwayTeam:    matchTeams[1],
		Event:       event,
		Way1Book:    float32(way1Book),
		WayXBook:    float32(wayXBook),
		Way2Book:    float32(way2Book),
		Way1Pool:    float32(way1Pool),
		WayXPool:    float32(wayXPool),
		Way2Pool:    float32(way2Pool),
		Score:       score,
		Way:         way,
	}
	return matchInfo
}

func getContentText(e *colly.HTMLElement) structures.ContentData {
	text := strings.TrimSpace(e.Text)
	newEntry := structures.ContentData{Data: text}
	return newEntry
}

func extractDrawInfo(tableData structures.ContentData) structures.DrawInfo {
	drawInfoSlice := strings.Fields(strings.TrimSpace(tableData.Data))

	drawNumber, _ := strconv.Atoi(drawInfoSlice[2])
	drawStatus := drawInfoSlice[4]
	drawDate := drawInfoSlice[6]
	drawTime := drawInfoSlice[7]

	combinedJackpotStr := drawInfoSlice[9] + drawInfoSlice[10] + drawInfoSlice[11]
	jackpot, _ := strconv.ParseFloat(combinedJackpotStr, 32)

	combinedPoolStr := drawInfoSlice[14] + drawInfoSlice[15] + drawInfoSlice[16]
	pool, _ := strconv.ParseFloat(combinedPoolStr, 32)

	drawInfo := structures.DrawInfo{
		DrawNumber: drawNumber,
		Status:     drawStatus,
		Date:       drawDate,
		Time:       drawTime,
		Jackpot:    structures.CurrencyRUR(jackpot),
		Pool:       structures.CurrencyRUR(pool),
	}
	return drawInfo
}

// fmt.Printf("drawTime: %s\n", drawTime)
// for index, value := range drawInfoSlice {
// 	fmt.Printf("Index: %d, Value: %s\n", index, value)
// }
// fmt.Println("Length of secondTablesSlice:", len(tableDataSlice))
// fmt.Printf("Appended text: %s (Type: %s)\n", text, reflect.TypeOf(text))
