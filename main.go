package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)


func main() {
	//get console arguments
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		//fmt.Println(scanner.Text()) // Println will add back the final '\n'
		input := clean_input(scanner.Text())
		//get the command
		command := input[0]
		//get the arguments
		arguments := input[1:]
		//process the commands
		process_input(command, arguments)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	//read out command and arguements
	//command, arguments := cleanInput(text string)

}


/*
function that sanitized the console command to seperate topics
https://pkg.go.dev/strings#Fields
*/
func clean_input(text string) []string  {
	//convert to lower case
	input_lower_case := strings.ToLower(text)
	//arguments
	split_text := strings.Fields(input_lower_case)
	//split up the string into seperate strings using whitespaces
	return split_text
}


/*

*/
func process_input(command, arguments) {
	//according to the command
	switch command {
		//we want to search the pokemon
		case pokemon: 
			//in case asking for commands
			fmt.Println("Your pokemon is: ", arguments)

		case help:
			if arguments.length == 0 {
				fmt.Println("No arguments provided")
			} else {
				//only get the first arguement
				argument := arguments[0]
				//according to the argument
				switch argument {
					//just print for now
					default:
						fmt.Println("Argument: ", argument)
				}
			}
			
		default:
			//just print the information
			fmt.Println("Command not found!")
			fmt.Println("Command: ", command)
			fmt.Println("Argument(s): ", arguments)	)
	}
}
