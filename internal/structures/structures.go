package structures

import "time"

type ContentData struct {
	Data string
}

type RUR struct {
	Amount float32
}

type DrawInfo struct {
	DrawNumber int
	Status     string
	Date       time.Time
	Time       time.Time
	Jackpot    RUR
	Pool       RUR
}
