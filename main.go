package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatalf("Error: Only one argument is required")
	}
	inputFile := os.Args[1]

	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		log.Fatalf("Error: File %s does not exist!", inputFile)
	}

	db, err := sql.Open("sqlite3", inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	outputFile := "Kobo_Wordlist.txt"

	if _, err := os.Stat(outputFile); err == nil {
		var userInput string
		fmt.Printf("Warning! File: %s already exists. Append to end of file? (yes or no): ", outputFile)
		for {
			_, err := fmt.Scanln(&userInput)
			if err != nil {
				fmt.Println("Something went wrong... Try again")
				fmt.Scanf("%v", &userInput)
				continue
			}
			userInput = strings.TrimSpace(strings.ToLower(userInput))

			if userInput == "yes" || userInput == "y" {
				fmt.Println("Roger that. Appending new words to end of file!")
				break
			} else if userInput == "no" || userInput == "n" {
				fmt.Printf("Quitting! Please delete/remove %s and try again\n", outputFile)
				os.Exit(1)
			} else {
				fmt.Println("Invalid input. Please respond with yes or no.")
			}
		}
	}

	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	theWords, err := db.Query("SELECT Text FROM WordList")
	if err != nil {
		log.Fatal(err)
	}
	defer theWords.Close()

	for theWords.Next() {
		var name string
		err = theWords.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(name)
		_, err := writer.WriteString(name + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	if err = theWords.Err(); err != nil {
		log.Fatal(err)
	}
	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Finished writing to Kobo_Wordlist.txt! Enjoy your fancy words :)")

}
