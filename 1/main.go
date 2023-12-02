package main

/*
	Дана структура Human (с произвольным набором полей и методов).
	Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

import (
	"fmt"
	"time"
)

type Human struct {
	Name       string
	Age        int
	Occupation string
}

func (h *Human) SetName(name string) {
	h.Name = name
}

func (h *Human) SetAge(age int) {
	h.Age = age
}

func (h *Human) SetOccupation(occupation string) {
	h.Occupation = occupation
}

type Action struct {
	Name     string
	Duration time.Duration
	Human
}

func (a *Action) SetName(name string) {
	a.Name = name
}

func (a *Action) SetDuration(duration time.Duration) {
	a.Duration = duration
}

func main() {
	var action Action = Action{
		Name:     "Make coffee",
		Duration: 25 * time.Second,
		Human: Human{
			Name:       "Andrew",
			Age:        23,
			Occupation: "Barista",
		},
	}

	/*
		Обращение к полям
	*/

	// При наличии одикаковых полей, приоритет у поля более верхнего уровня
	fmt.Println(action.Name)       // Make coffee
	fmt.Println(action.Human.Name) // John

	// Если поле уникальное, то можно напрямую обращаться к нему из родительской структуры
	fmt.Println(action.Age) // 23

	// При вызове методов с одинаковыми именами идёт вызов вышестоящей структуры
	action.SetName("Make tea") // Make coffee -> Make tea

	/*
		Вызов методов
	*/

	// методы встроенной структуры наследуются родительской
	action.SetAge(15)
	fmt.Println(action.Age)

	// при наличии одинаковых методов, приоритет у метода структуры более верхнего уровня
	action.SetName("Make hot chocolate")
	fmt.Println(action.Name)       // Make tea
	fmt.Println(action.Human.Name) // Jane

	// чтобы вызвать метод встроенной структуры необходимо обратиться к ней напрямую
	action.Human.SetName("John")
	fmt.Println(action.Human.Name) // John
}
