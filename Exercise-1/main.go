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
	"bufio"
)

var status bool = false
var correct float64 = 0
var incorrect float64 = 0
var questions float64 = 0

// quiz has 30s timer unless specified by the user
func timer() {
	// get specified time; else use default time
	timer := 30
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
	fmt.Printf("\n---PENCILS DOWN---\n") // timer is up
	time.Sleep(3)
	// check if the score is 0; 
	if correct == 0 {
		fmt.Printf("Score: 0%")
		os.Exit(0)
	} else {
		result := (correct/questions) * 100
		//fmt.Printf("you got %v correct out of %v questions.", correct, questions)
		fmt.Printf("Score: %.2f%", result)
		os.Exit(0)
	}
	
	// finish
	os.Exit(0)
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

	f, err := os.Open(filename) // create new file object called f
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close() // close the file at the end

	csvReader := csv.NewReader(f) // new csvReader with file f

	for status == false {
		// for each line in csv file f
		for {
			// fmt.Printf("%v", status)
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
			fmt.Printf("%v ", question)
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
}
