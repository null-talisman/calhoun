// Exercise 1 from Calhoun.io

package main

import (
	"fmt"
	"encoding/csv"
	"os"
	"io"
	"log"
	"strconv"
	"time"
)

var status bool = false // status of quiz

// quiz has 30s timer unless specified by the user
func timer() {
	// get specified time; else use default time
	timer := 5
	if len(os.Args) > 2 {
		x, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		timer = x
	}	
	// time.Duration(timer)
	timer1 := time.NewTimer(time.Duration(timer)*time.Second) // create timer
	<-timer1.C // start timer
	fmt.Printf("\n---PENCILS DOWN---") // timer is up
	status = true
	return
}

func main() {

	go timer() // start timer

	// use default filename unless specified
	filename := ""
	if len(os.Args) < 2 {
        filename = "problems.csv"
    }	else {
		filename = os.Args[1]
	}

	correct   := 0 // count number of correct answers
	incorrect := 0 // count number of incorrect answers

	f, err := os.Open(filename) // create new file object called f

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close() // close the file at the end

	csvReader := csv.NewReader(f) // new csvReader with file f

	for status == false {
		// for each line in csv file f
		for {
			fmt.Printf("%v", status)
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
	}

	// print results of the quiz
	fmt.Printf("---RESULTS---\nCorrect: %v\nIncorrect: %v\n", correct, incorrect)

}
