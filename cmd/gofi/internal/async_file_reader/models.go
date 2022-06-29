package asyncfilereader

// YYYY-MM-DD
const TimeFormat = "2009-10-11"

type Transaction struct {
	ID        string
	From      string
	To        string
	Amount    float32
	CreatedAt string
}
