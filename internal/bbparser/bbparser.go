// package bbparser
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"

	"totolab/internal/structures"
)

// type contentData struct {
// 	data string
// }

func main() {
	var tableDataSlice []structures.ContentData
	c := colly.NewCollector()

	c.OnHTML("div.table-responsive", func(e *colly.HTMLElement) {
		text := strings.TrimSpace(e.Text)
		newEntry := structures.ContentData{Data: text}
		tableDataSlice = append(tableDataSlice, newEntry)

		// length := len(tableDataSlice)
		// fmt.Println("Length of secondTablesSlice:", length)
		// fmt.Printf("Appended text: %s (Type: %s)\n", text, reflect.TypeOf(text))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://totobrief.ru/results/baltbet/10377")

	firstTable := tableDataSlice[0]
	firstTableSlice := strings.Fields(strings.TrimSpace(firstTable.Data))

	for index, value := range firstTableSlice {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}
	combinedAmount := firstTableSlice[9] + firstTableSlice[10] + firstTableSlice[11]
	floatValue, _ := strconv.ParseFloat(combinedAmount, 32)
	fmt.Printf("Value: %f\n", floatValue)

	// fmt.Printf("firstTable text: %s (Type: %s)\n", firstTable, reflect.TypeOf(firstTable.data))
}
