// Exercise 1 from Calhoun.io
// Not exactly what it's suppose to do but still something.

package main

import "fmt"
import "encoding/csv"
import "os"
import "io"
import "log"
import "strconv"

func main() {
 
	fmt.Println("---GOLANG QUIZ---") // intro

	// get specified filename; else use default filename
	filename := ""
	if len(os.Args) < 2 {
        filename = "problems.csv"
    }	else {
			filename = os.Args[1]
	}

	correct   := 0 // count number of correct answers
	incorrect := 0 // count number of incorrect answers

	// create new file object called f
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close() // close the file at the end

	csvReader := csv.NewReader(f) // new csvReader with file f

	// for each line in csv file f
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

		// ask the question
		fmt.Printf("%v=", question)
		fmt.Scanln(&result)

		// conversions
		answer_int, err := strconv.Atoi(answer)
		result_int, err := strconv.Atoi(result)

		// check answer; increment counters
		if result_int == answer_int {
			correct++
		}	else {
			incorrect++
		}

	}

	// print results of the quiz
	fmt.Print("---RESULTS---\n")
	fmt.Printf("Correct: %v\n", correct)
	fmt.Printf("Incorrect: %v\n", incorrect)

}
