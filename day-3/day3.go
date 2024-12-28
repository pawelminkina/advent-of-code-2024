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
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var totalResult = 0
	var allTextsb strings.Builder
	for scanner.Scan() {
		var text = scanner.Text()
		allTextsb.WriteString(strings.TrimRight(text, "\r\n"))
	}
	totalResult = getMul(allTextsb.String())
	fmt.Println("total result", totalResult)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// jezeli to nie zadziala musze polaczyc wszystkie linijki ze soba
func getMul(lineOfText string) int {

	var splittedByDont = strings.Split(lineOfText, "don't()")
	//first item is ok, what with rest?
	//second has a dont and do's
	var lineOfTextsb strings.Builder

	for i := 0; i < len(splittedByDont); i++ {
		if i == 0 {
			lineOfTextsb.WriteString(splittedByDont[0])
			continue
		}
		var textByDont = splittedByDont[i]
		var doIndex = strings.Index(textByDont, "do()")
		if doIndex == -1 {
			continue
		}
		lineOfTextsb.WriteString(textByDont[doIndex:])

	}
	var muls = strings.SplitAfter(lineOfTextsb.String(), "mul(")

	if len(muls) == 0 {
		return 0
	}
	var result = 0
	for i := 0; i < len(muls); i++ {
		var mul = muls[i]
		var sb strings.Builder
		for _, r := range mul {
			if r < 3 {
				continue
			}
			if string(r) == ")" {
				break
			}
			sb.WriteString(string(r))
		}

		var valuesToOperateOn []int
		var valuesInsideMul = strings.Split(sb.String(), ",")
		for j := 0; j < len(valuesInsideMul); j++ {
			parsedValue, parseError := strconv.Atoi(valuesInsideMul[j])
			if parseError != nil { //error parsing to int
				valuesToOperateOn = nil
				break
			}

			//Empty character was inside strinbg
			valueBackToString := strconv.Itoa(parsedValue)
			if len(valueBackToString) != len(valuesInsideMul[j]) {
				valuesToOperateOn = nil
				break
			}

			if parsedValue < 1 || parsedValue > 999 {
				valuesToOperateOn = nil
				break
			}

			valuesToOperateOn = append(valuesToOperateOn, parsedValue)

		}

		if len(valuesToOperateOn) == 2 {
			result += valuesToOperateOn[0] * valuesToOperateOn[1]
		}

	}
	return result

}
