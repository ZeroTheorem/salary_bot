package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	tele "gopkg.in/telebot.v4"
)

const resultMsg string = `
Result ✨

All: <b>%2.f₽</b>
Advance: <b>%2.f₽</b>
Basic: <b>%2.f₽</b>

Bonus:
Level 100%%: <b>%2.f₽</b> (total: <b>%2.f₽</b>)
Level 75%%: <b>%2.f₽</b> (total: <b>%2.f₽</b>)
Level 60%%: <b>%2.f₽</b> (total: <b>%2.f₽</b>)
Level 50%%: <b>%2.f₽</b> (total: <b>%2.f₽</b>)
Level 25%%: <b>%2.f₽</b> (total: <b>%2.f₽</b>)
`

const perShift float64 = 1646.04

var shifts float64

func main() {
	pref := tele.Settings{
		Token:     "8185753212:AAES-cij7lkJ2R5XJ-4OKRfHbvPoFoPkDfc",
		Poller:    &tele.LongPoller{Timeout: 10 * time.Second},
		ParseMode: tele.ModeHTML,
	}
	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	selector := &tele.ReplyMarkup{}
	calculateBtn := selector.Data("Calculate", "calculate")
	selector.Inline(selector.Row(calculateBtn))

	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Hi, enter your number of shifts so I can calculate your income!")
	})

	b.Handle(tele.OnText, func(c tele.Context) error {
		num, err := strconv.ParseFloat(c.Message().Text, 64)
		if err != nil {
			return c.Send("Please enter a number")
		}
		shifts = num
		return c.Send(fmt.Sprintf("Number of shifts: <b>%v</b> Calculate?", shifts), selector)
	})

	b.Handle(&calculateBtn, func(c tele.Context) error {
		return c.Send(calculateSalary(shifts))
	})

	b.Start()
}

func calculateSalary(shifts float64) string {
	total := shifts * perShift
	advance := percent(40, total)
	salary := percent(60, total)
	percent75 := percent(75, total)
	percent50 := percent(50, total)
	percent25 := percent(25, total)
	return fmt.Sprintf(resultMsg,
		total,
		advance,
		salary,
		total,
		total*2,
		percent75,
		total+percent75,
		salary,
		salary+total,
		percent50,
		total+percent50,
		percent25,
		total+percent25,
	)
}
func percent(percent, value float64) float64 {
	return value * percent / 100
}
