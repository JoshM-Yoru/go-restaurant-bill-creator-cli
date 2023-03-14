package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type menuItem struct {
	dish  string
	price float64
}

type bill struct {
	name  string
	items []menuItem
	tip   float64
	total float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: []menuItem{},
		tip:   0,
		total: 0,
	}

	return b
}

func (b *bill) format() string {
	fs := fmt.Sprintf("\n\n%v's Bill Breakdown \n---------------\n", b.name)

	for _, item := range b.items {
		fs += fmt.Sprintf("%-35v ...$%v \n", item.dish+":", item.price)
		b.total += item.price
	}

	fs += fmt.Sprintf("%-35v ...$%0.2f \n", "Tip:", b.tip)

	fs += fmt.Sprintf("%-35v ...$%0.2f", "Total:", b.total)

	return fs
}

func (b *bill) calcSubtotal() float64 {
	var subtotal float64 = 0

	for _, item := range b.items {
		subtotal += item.price
	}

	return subtotal
}

func (b *bill) addTip(tipPercentage float64) {
	b.tip = b.calcSubtotal() * tipPercentage
	b.total = b.calcSubtotal() + b.tip
	b.promptOptions()
}

func (b *bill) customTip() {
	reader := bufio.NewReader(os.Stdin)

	tipAmount, _ := getInput("How much would you like to tip in $:", reader)
	if tip, err := strconv.ParseFloat(tipAmount, 64); err == nil {
		b.tip = tip
		b.total = b.calcSubtotal() + b.tip
		b.promptOptions()
	} else {
		fmt.Println("\n\nInvalid Input. Try Again.\n")
		b.customTip()
	}
}

func (b *bill) addItem(dish string, price float64) {
	b.items = append(b.items, menuItem{dish, price})
}