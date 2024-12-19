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
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
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
func clean_input(text string) string, []string  {
	//convert to lower case
	input_lower_case := strings.toLower(text)
	//arguments
	split_text := strings.fields(input_lower_case)
	//get the command
	command := split_text[0]
	//get the arguments
	arguments := split_text[1:]
	
	//print the split arguments
	fmt.Println("Command: ", command)
	fmt.Println("Argument(s): ", arguments)	
	
	//split up the string into seperate strings using whitespaces
	return command, arguments
}
