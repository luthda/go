package expenses

import "errors"

type Record struct {
	Day      int
	Amount   float64
	Category string
}

type DaysPeriod struct {
	From int
	To   int
}

func Filter(in []Record, predicate func(Record) bool) []Record {
	var out []Record

	for _, record := range in {
			if predicate(record) {
				out = append(out, record)
			}	
		}

	return out
}

func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		if r.Day >= p.From && r.Day <= p.To {
			return true
		}

		return false
	}
}

func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}

func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	var total float64

	for _, record := range in {
		if record.Day >= p.From && record.Day <= p.To {
			total += record.Amount
		}
	}

	return total
}

func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	recordesByCategory := Filter(in, ByCategory(c))

	if len(recordesByCategory) == 0 {
		return 0, errors.New("no records found for category")
	}

	return TotalByPeriod(recordesByCategory, p), nil
}