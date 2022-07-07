package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"log"
	"strings"
	"strconv"
)

func readCsv() {

	f, err := os.Open("example.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)

	correct := 0
	num_questions := 0

	for {
		sum := 0
		x, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		num_questions = num_questions + 1
		question := strings.Split(x[0], "+")
		for i := range question{
			s_int, err := strconv.Atoi(question[i])
			if err != nil {
				log.Fatal(err)
			}
			sum = s_int + sum
		}
		answer_int, err := strconv.Atoi(x[1])
		if err != nil {
			log.Fatal(err)
		}
		if sum == answer_int {
			correct = correct + 1
		}
	}
	fmt.Printf("Number of questions: %v\n", num_questions)
	fmt.Printf("Correct Answers: %v\n", correct)
}


func main() {
	readCsv()
}
