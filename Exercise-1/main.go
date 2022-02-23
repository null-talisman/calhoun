// Exercise 1 from Calhoun.io
// Not exactly what it's suppose to do but still something.

package main

import "fmt"
import "encoding/csv"
import "os"
import "io"
import "log"
//import "strings"
import "strconv"

func main() {
	// just do something lol
	fmt.Println("---TESTING---")

	// create new file object called f
	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program ? 
	defer f.Close()


	index := 1

	// new csvReader with file f
	csvReader := csv.NewReader(f)
	for {
		// csv is read in as []strings
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// so [5+5, 10] where rec[0] == '5+5' and rec[1] is 10
		question := rec[0]
		answer   := rec[1]
		result   := ""

		// get the digits 
		num_one, err := strconv.Atoi(question[0:1])
		num_two, err := strconv.Atoi(question[2:3])
		ans_str, err := strconv.Atoi(answer)
		if err != nil {
			log.Fatal(err)
		} 
		sum := num_one + num_two

		// is it correct
		if sum == ans_str {
			result = "Correct"
		}	else {
			result = "Incorrect"
		}
		
		// question counter
		fmt.Printf("Q.%v: %v (%v/%v)\n", index, result, sum, ans_str)
		index++

	}
}
