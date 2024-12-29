package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var totalResult = 0
	var texts []string
	for scanner.Scan() {
		var text = scanner.Text()
		texts = append(texts, text)
	}
	totalResult = checkItemsPart2(texts)
	fmt.Println("total result", totalResult)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func checkItems(items []string) int {
	var totalCount = 0
	for i := 0; i < len(items); i++ {
		var currentLine = items[i]
		//horizontal
		for j := 0; j < len(currentLine); j++ {
			if j+3 < len(currentLine) {
				totalCount += isXmax(string(currentLine[j]), string(currentLine[j+1]), string(currentLine[j+2]), string(currentLine[j+3]))
			}
			if j-3 >= 0 {
				totalCount += isXmax(string(currentLine[j]), string(currentLine[j-1]), string(currentLine[j-2]), string(currentLine[j-3]))
			}
		}

		//going down
		if i+3 < len(items) {
			for j := 0; j < len(currentLine); j++ {
				//check just going down
				totalCount += isXmax(string(items[i][j]), string(items[i+1][j]), string(items[i+2][j]), string(items[i+3][j]))

				//check going diagonal
				if j+3 < len(currentLine) {
					//diagonal going right
					totalCount += isXmax(string(items[i][j]), string(items[i+1][j+1]), string(items[i+2][j+2]), string(items[i+3][j+3]))
				}
				if j-3 >= 0 {
					//diagonal going left
					totalCount += isXmax(string(items[i][j]), string(items[i+1][j-1]), string(items[i+2][j-2]), string(items[i+3][j-3]))
				}
			}
		}

		//going up
		if i-3 >= 0 {
			for j := 0; j < len(currentLine); j++ {
				//check just going up
				totalCount += isXmax(string(items[i][j]), string(items[i-1][j]), string(items[i-2][j]), string(items[i-3][j]))

				//check going diagonal
				if j+3 < len(currentLine) {
					//diagonal going right
					totalCount += isXmax(string(items[i][j]), string(items[i-1][j+1]), string(items[i-2][j+2]), string(items[i-3][j+3]))
				}
				if j-3 >= 0 {
					//diagonal going left
					totalCount += isXmax(string(items[i][j]), string(items[i-1][j-1]), string(items[i-2][j-2]), string(items[i-3][j-3]))
				}
			}
		}
	}
	return totalCount
}

func isXmax(item1, item2, item3, item4 string) int {
	if item1 != "X" {
		return 0
	}
	if item2 != "M" {
		return 0
	}
	if item3 != "A" {
		return 0
	}
	if item4 != "S" {
		return 0
	}
	return 1
}

func checkItemsPart2(items []string) int {
	var totalCount = 0
	for i := 1; i < len(items)-1; i++ {
		var currentLine = items[i]

		for j := 1; j < len(currentLine)-1; j++ {
			if string(items[i][j]) != "A" {
				continue
			}

			var isMasInXShape = (ismas(string(items[i-1][j-1]), string(items[i][j]), string(items[i+1][j+1])) ||
				ismas(string(items[i+1][j+1]), string(items[i][j]), string(items[i-1][j-1]))) &&
				(ismas(string(items[i-1][j+1]), string(items[i][j]), string(items[i+1][j-1])) ||
					ismas(string(items[i+1][j-1]), string(items[i][j]), string(items[i-1][j+1])))

			if isMasInXShape {
				totalCount++
			}
		}
	}
	return totalCount
}

func ismas(item1, item2, item3 string) bool {
	if item1 != "M" {
		return false
	}
	if item2 != "A" {
		return false
	}
	if item3 != "S" {
		return false
	}
	return true
}
