package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func (b *bill) addTipPrompt() {
	reader := bufio.NewReader(os.Stdin)

	option, _ := getInput("Select the number for how much you would like to tip: \n1: 10%\n2: 15%\n3: 20%\n4: Custom\n", reader)

	switch option {
	case "1":
		b.addTip(0.10)
	case "2":
		b.addTip(0.15)
	case "3":
		b.addTip(0.20)
	case "4":
		b.customTip()
	default:
		fmt.Println("\n\nInvalid Input. Try Again.\n")
		b.addTipPrompt()
	}

}

func (b *bill) promptOptions() {
	reader := bufio.NewReader(os.Stdin)

	option, _ := getInput("Choose option (a - add more items, t - add tip, s - show and save bill): ", reader)
	fmt.Println(option)

	switch option {
	case "a":
		b.selectMenuItemPrompt()
	case "s":
		fmt.Println(b.format())
		b.saveBill()
	case "t":
		b.addTipPrompt()
	default:
		fmt.Println("Invalid option, please pick correctly...")
		b.promptOptions()
	}

}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)

	return b
}

func main() {

	mybill := createBill()
	mybill.selectMenuItemPrompt()

}
