package main
import "log"
import "bufio"
import "os"
import "fmt" 
import "strings"
import "math"
import "strconv"


func main() {
    file, err := os.Open("testinput.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
   
    var safeReports = 0
    for scanner.Scan() {
        var text = scanner.Text()
        fmt.Println("current report: ", text)
        var itemsInOneLine = strings.Split(text, " ")
        var lastNumber = math.MinInt
        var isSafe = false
        var alreadyFailedForLine = false
        var processedNumbers []int
        for i := 0; i < len(itemsInOneLine); i++ {
            currentNumber, _ := strconv.Atoi(itemsInOneLine[i])
            var isIncreasing = IsIncreasingFunc(currentNumber, lastNumber) 
            var isDecreasing = IsDecreasingFunc(currentNumber, lastNumber)
            if i > 0 && !isIncreasing && !isDecreasing {
                if alreadyFailedForLine  {
                    fmt.Println("Failed line: ", text)
                    isSafe = false
                    break
                }
                if len(processedNumbers) > 2 && !IsIncreasingFunc(currentNumber, processedNumbers[i-2]) && !IsDecreasingFunc(currentNumber, processedNumbers[i-2]) {
                    fmt.Println("Happened here for text: ", text)
                    alreadyFailedForLine = true
                    continue
                }
                if (len(itemsInOneLine) > 2 && i == 1){
                    nextItem, _ := strconv.Atoi(itemsInOneLine[i+1])
                    if IsIncreasingFunc(nextItem, currentNumber) || IsDecreasingFunc(nextItem, currentNumber){
                        alreadyFailedForLine = true

                        processedNumbers = append(processedNumbers, currentNumber)
                        lastNumber = currentNumber
                        isSafe = true
                        continue
                    }
                } 
                alreadyFailedForLine = true
                continue
            }
          processedNumbers = append(processedNumbers, currentNumber)
          lastNumber = currentNumber
          isSafe = true
          if i > 0 && !(areAllDecreasing(processedNumbers) || areAllIncreasing(processedNumbers)) {

                fmt.Println("failed for number : ", currentNumber)
                fmt.Println("in text: ", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(processedNumbers)), ","), "[]"))
            if alreadyFailedForLine {
                fmt.Println("Unsafe report, isIncreasing: ", text)
                isSafe = false
                break
                }
            if len(processedNumbers) == 3 {
                var newNumbers []int 
                    for j := 1; j < len(itemsInOneLine); j++ {
                        val, _ := strconv.Atoi(itemsInOneLine[j])
                        newNumbers = append(newNumbers, val)
                    }
                    fmt.Println("newnumbers: ", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(newNumbers)), ","), "[]"))
                    var removingFirstWorks = areAllDecreasing(newNumbers) || areAllIncreasing(newNumbers)
                fmt.Println("removing first works: ", removingFirstWorks)
                    if (removingFirstWorks) {
                        processedNumbers = processedNumbers[1:] 
                fmt.Println("fater removal: ", processedNumbers)
                        alreadyFailedForLine = true
                        continue
                    }
                }
            processedNumbers = processedNumbers[:len(processedNumbers)-1]
            processedNumbers[i-1] = currentNumber
            alreadyFailedForLine = true
          }
        }
        if isSafe && len(itemsInOneLine) > 1 {
            fmt.Println("Safe report: ", text)
            safeReports++
        }
    }
    fmt.Println("Safe reports: ", safeReports) 
 
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func IsIncreasingFunc(currentNumber int, lastNumber int) bool{
            return currentNumber  >= lastNumber + 1 && currentNumber <= lastNumber + 3
}


func IsDecreasingFunc(currentNumber int, lastNumber int) bool{
           return currentNumber >= lastNumber - 3 && currentNumber <= lastNumber - 1
}

func areAllDecreasing(numbers []int) bool {
    var lastNumber = 0
    var result = true
    for i := 0; i < len(numbers); i++ {
       if i == 0{
        lastNumber = numbers[i]
            continue
        }
       if numbers[i] < lastNumber {
            result = result && true 
            lastNumber = numbers[i]
        } else {
        result = false
        break
        }
    }
    return result
}

func areAllIncreasing(numbers []int) bool {
    var lastNumber = 0
    var result = true
    for i := 0; i < len(numbers); i++ {
       if i == 0{
        lastNumber = numbers[i]
            continue
        }
       if numbers[i] > lastNumber {
            result = result && true 
            lastNumber = numbers[i]
        } else {
        result = false
        break
        }
    }
    return result
}
