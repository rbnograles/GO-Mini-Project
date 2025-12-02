package main

import (
	"fmt"
	"math/rand"
)

func main() {
	
	var randNumb, guessNumb, tries int;
	
	fmt.Println("Guess the correct number between 0 - 100");
	fmt.Println("You will only have 10 tries so think before you type!")

	randNumb = rand.Intn(100)

	fmt.Println("====================================")
	fmt.Println("*****Guess what is the number*****")
	fmt.Println("====================================")
	
	for {
		fmt.Print("Your guess: ")
		fmt.Scanln(&guessNumb);	

		if(guessNumb > randNumb) {
			fmt.Println("Guess lower!!!!")
			fmt.Println("Current tries", tries + 1);
		} else {
			fmt.Println("Guess higher!!!")
			fmt.Println("Current tries", tries + 1);
		}


		if(tries == 9) {
			fmt.Println("You lose!")
			fmt.Println("The random number is", randNumb);
			break;
		}

		if(randNumb == guessNumb) {
			fmt.Println("Congratuations! You win!")
			fmt.Println("The random number is", randNumb);
			break;
		}

		tries += 1;
	}
}
