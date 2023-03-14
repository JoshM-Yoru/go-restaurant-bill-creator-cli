package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (b *bill) selectMenuItemPrompt() {

	menu := map[string]float64{
		"Chicken and Waffles":            12.87,
		"Ham and Cheese Omlete":          9.32,
		"Bacon, Egg and Cheese Sandwich": 7.54,
		"Side of Fruit":                  3.89,
		"Hashbrowns":                     3.22,
		"Bacon":                          5.21,
	}

	selection := map[int]string{}

	idx := 1
	fmt.Println("Add Menu Items To The Bill")
	fmt.Println("Breakfast Menu Items")
	fmt.Println("--------------------")

	for dish, price := range menu {

		selection[idx] = dish

		fmt.Println(idx, dish, "$", price)

		idx++
	}

	reader := bufio.NewReader(os.Stdin)

	option, _ := getInput("Select a Menu Item By Number on the Menu, to Select Multiple, Seperate Selections with a Comma: ", reader)
	trimmed := strings.TrimSpace(option)
	if strings.HasSuffix(trimmed, ",") {
		trimmed = trimmed[:len(trimmed)-1]
	}

	selections := strings.Split(trimmed, ",")

	var check bool = true

	for _, s := range selections {

		switch s {
		case "1":
			b.addItem(selection[1], menu[selection[1]])
		case "2":
			b.addItem(selection[2], menu[selection[2]])
		case "3":
			b.addItem(selection[3], menu[selection[3]])
		case "4":
			b.addItem(selection[4], menu[selection[4]])
		case "5":
			b.addItem(selection[5], menu[selection[5]])
		case "6":
			b.addItem(selection[6], menu[selection[6]])
		default:
			fmt.Println("")
			fmt.Println("")
			fmt.Println("Your input had incorrect inputs, please try again: ")
			fmt.Println("")
			check = false
			b.selectMenuItemPrompt()
		}

	}

	fmt.Println("Your subtotal is: ", b.calcSubtotal())

	if check {
		b.promptOptions()
	}

}
