package main

import "fmt"

func main() {

	var num1, num2, result int
	var operator string

	fmt.Println("============================================")
	fmt.Println("           Simple Calculator               ")
	fmt.Println("============================================")

	fmt.Print("Enter First number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter Second number: ")
	fmt.Scanln(&num2)

	for {
		fmt.Print("Enter Operator (+ - * /): ")
		fmt.Scanln(&operator)

		switch operator {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "/":
			if num2 == 0 {
				fmt.Println("Error: Division by zero")
				continue
			}
			result = num1 / num2
		default:
			fmt.Println("Invalid Operator. Please try again.")
			continue
		}

		break;
	}

	fmt.Println("Result:", result)
}
