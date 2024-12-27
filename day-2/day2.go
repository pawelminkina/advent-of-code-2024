package main
import "log"
import "bufio"
import "os"
import "fmt" 
import "strings"
import "math"
import "strconv"


func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
   
    var safeReports = 0
    for scanner.Scan() {
        var text = scanner.Text()
        var itemsInOneLine = strings.Split(text, " ")
        fmt.Println("current report: ", text)
        var isSafe = IsSafe(itemsInOneLine)
        if isSafe {
            fmt.Println("Safe report: ", text)
            safeReports++
        } else {
            for i := 0; i < len(itemsInOneLine); i++ {
            //w petli chce zwyczjanie robic remove at
                var slicedText = RemoveIndex(itemsInOneLine, i)
                isSafe = IsSafe(slicedText)
                if isSafe {
                    fmt.Println("Safe report: ", text)
                    safeReports++
                    break
                }
            }
        }
    }
    fmt.Println("Safe reports: ", safeReports) 
 
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func RemoveIndex(s []string, index int) []string {
    ret := make([]string, 0)
    ret = append(ret, s[:index]...)
    return append(ret, s[index+1:]...)
}

func IsSafe(itemsInOneLine []string) bool {
    var processedNumbers []int
    var lastNumber = math.MinInt
    for i := 0; i < len(itemsInOneLine); i++ {
            currentNumber, _ := strconv.Atoi(itemsInOneLine[i])
            var isIncreasing = IsIncreasingFunc(currentNumber, lastNumber) 
            var isDecreasing = IsDecreasingFunc(currentNumber, lastNumber)
            if i > 0 && !isIncreasing && !isDecreasing {
                return false
            }
    processedNumbers = append(processedNumbers, currentNumber)
    lastNumber = currentNumber
    }
 
    if !(areAllDecreasing(processedNumbers) || areAllIncreasing(processedNumbers)) {
        return false
    }
    return true
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
