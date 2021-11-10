package challengelib

import "fmt"

type Data struct {
	Divisor int
	Message string
}

func Run(first int, end int) {
	data := []Data{
		{Divisor: 15, Message: "Linianos"},
		{Divisor: 3, Message: "Linio"},
		{Divisor: 5, Message: "IT"},
	}

MAIN_LOOP:
	for dividend := first; dividend < end+1; dividend++ {
		if message, ok := GetMessage(dividend, data); ok {
			fmt.Println(message)
			continue MAIN_LOOP
		}
		fmt.Println(dividend)
	}

}

func GetMessage(dividend int, data []Data) (string, bool) {
	for _, d := range data {
		if dividend%d.Divisor == 0 {
			return d.Message, true
		}
	}
	return "", false
}
