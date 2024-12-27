package main

import "os"
import "log"
import "bufio"
import "fmt" 
import "strings"
import "math"
import "strconv"
import "sort"


func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    var distance = 0.0
    var listOne []float64
    var listTwo []float64
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        var itemsInOneLine = strings.Split(scanner.Text(), "   ")
        oneLine, _ := strconv.ParseFloat(itemsInOneLine[0], 64)
        twoLine, _ := strconv.ParseFloat(itemsInOneLine[1], 64)
        listOne = append(listOne, oneLine)
        listTwo = append(listTwo, twoLine)
    }
    
    sort.Float64s(listOne)
    sort.Float64s(listTwo)
    for i := 0; i < len(listOne); i++ {
        distanceToAdd := math.Abs(listOne[i] - listTwo[i])
        distance += distanceToAdd
    }
    s := strconv.FormatFloat(distance, 'f', -1, 64)
    fmt.Println("total distance is: ", s)
    
    var similarityScore = 0.0
    //some dictionary/ keyvault pair like val, number of times with counter
    counterSlice := make(map[float64]int)
    for i := 0; i < len(listTwo); i++ {
        counterSlice[listTwo[i]]++
    }
    for i := 0; i < len(listOne); i++ {
       similarityScore += listOne[i] * float64(counterSlice[listOne[i]])
    }
    
    g := strconv.FormatFloat(similarityScore, 'f', -1, 64)
    fmt.Println("total similiary score is: ", g)
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    //foreach line split ('  '), abs add 2 both values, save to some variable print result


}

