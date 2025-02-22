package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		log.Fatal("error opening file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var txtLines []string

	for scanner.Scan() {
		txtLines = append(txtLines, scanner.Text())
	}

	file.Close()

	var leftList, rightList []int

	for _, line := range txtLines {
		numbers := strings.Fields(line)
		if len(numbers) != 2 {
			log.Fatal("wrong number of fields in line")
		}

    left, err := strconv.Atoi(numbers[0])
		if err != nil {
			log.Fatal("error converting number", err)
		}

		right, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatal("error converting number", err)
		}

    leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

  if(len(leftList) != len(rightList)) {
    log.Fatal("lengths don't match")
  }

  similarityScore := 0

  for _, number := range leftList {
    occurences := 0
    for _, rightNumber := range rightList {
      if number == rightNumber {
        occurences ++
      }
    }
    similarityScore += occurences*number
  }

  fmt.Println(similarityScore)

}
