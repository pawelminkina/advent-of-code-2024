package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var totalResult = 0
	rules := make(map[string][]string)
	var orders []string
	for scanner.Scan() {
		var text = scanner.Text()
		if strings.Contains(text, "|") {
			var rule = strings.Split(text, "|")
			rules[rule[0]] = append(rules[rule[0]], rule[1])
		}
		if strings.Contains(text, ",") {
			orders = append(orders, text)
		}
	}
	//actually nope, part 2 requires different ordering, the fact that it contains or not is not enough
	totalResult = getTotalOfMiddleNumbers(rules, orders)
	fmt.Println("total result", totalResult)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
func getTotalOfMiddleNumbers(rules map[string][]string, orders []string) int {
	var totalOfMiddleNumbers = 0
	for i := 0; i < len(orders); i++ {
		var ordersItem = strings.Split(orders[i], ",")
		totalOfMiddleNumbers += getMiddleNumberForOneItem(rules, ordersItem)
	}
	return totalOfMiddleNumbers
}
func getMiddleNumberForOneItem(rules map[string][]string, ordersItem []string) int {
	for j := 0; j < len(ordersItem); j++ {
		for k := j + 1; k < len(ordersItem); k++ {
			if !contains(rules[ordersItem[j]], ordersItem[k]) {
				return 0
			}
		}
	}
	item, _ := strconv.Atoi(ordersItem[int(math.Floor(float64(len(ordersItem))/2.0))])
	return item
}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
