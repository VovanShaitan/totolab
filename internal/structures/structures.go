package structures

type ContentData struct {
	Data string
}

type CurrencyRUR float32

type CurrencyEUR float32

type DrawInfo struct {
	DrawNumber int
	Status     string
	Date       string
	Time       string
	Jackpot    CurrencyRUR
	Pool       CurrencyRUR
}

type MatchInfo struct {
	MatchNumber int
	HomeTeam    string
	AwayTeam    string
	Event       string
	Way1Book    float32
	WayXBook    float32
	Way2Book    float32
	Way1Pool    float32
	WayXPool    float32
	Way2Pool    float32
	Score       string
	Way         string
}
