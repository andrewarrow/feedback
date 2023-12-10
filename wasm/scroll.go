package wasm

import "fmt"

func (g *Global) IsBottom() bool {
	de := g.Document.Document.Get("documentElement")
	db := g.Document.Document.Get("body")
	scrollTop := de.Get("scrollTop").Int()
	if scrollTop == 0 {
		scrollTop = db.Get("scrollTop").Int()
	}
	windowHeight := g.Window.GetInt("innerHeight")
	fmt.Println("scrollTop", scrollTop, innerHeight)

	a1 := de.Get("scrollHeight").Int()
	a2 := db.Get("scrollHeight").Int()
	a3 := de.Get("offsetHeight").Int()
	a4 := db.Get("offsetHeight").Int()
	a5 := de.Get("clientHeight").Int()
	a6 := db.Get("clientHeight").Int()
	documentHeight := max(a1, a2, a3, a4, a5, a6)

	return scrollTop+windowHeight >= documentHeight-100
}

func max(numbers ...int) int {
	maxValue := numbers[0]
	for _, num := range numbers[1:] {
		if num > maxValue {
			maxValue = num
		}
	}

	return maxValue
}
