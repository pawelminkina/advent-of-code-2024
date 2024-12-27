package main
import "log"
import "bufio"
import "os"
import "fmt" 
import "strings"


func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    var allSafeLine []string 
    var allFailedLine []string 
    var allUnsafeIncreasingLine []string 
    for scanner.Scan() {
        var text = scanner.Text()
        if strings.Contains(text, "Happened here for text:  ") {
            allFailedLine = append(allFailedLine, strings.TrimLeft(text, "Happened here for text:  "))
        }
        if strings.Contains(text, "Safe report:  ") {
            allSafeLine = append(allSafeLine, strings.TrimLeft(text, "Safe report:  "))
        }
        if strings.Contains(text, "Unsafe report, isIncreasing:  ") {
            allUnsafeIncreasingLine = append(allUnsafeIncreasingLine, strings.TrimLeft(text, "Unsafe report, isIncreasing:  "))
        }
    } 
    var intersect = intersection(allFailedLine, allSafeLine)

    for i := 0; i < len(intersect); i++ {
        
        fmt.Println("intersect: ", intersect[i])
    }
    for i := 0; i < len(allFailedLine); i++ {
        
        fmt.Println("unsafe: ", allFailedLine[i])
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
func intersection(s1, s2 []string) (inter []string) {
    hash := make(map[string]bool)
    for _, e := range s1 {
        hash[e] = true
    }
    for _, e := range s2 {
        // If elements present in the hashmap then append intersection list.
        if hash[e] {
            inter = append(inter, e)
        }
    }
    //Remove dups from slice.
    inter = removeDups(inter)
    return
}

//Remove dups from slice.
func removeDups(elements []string)(nodups []string) {
    encountered := make(map[string]bool)
    for _, element := range elements {
        if !encountered[element] {
            nodups = append(nodups, element)
            encountered[element] = true
        }
    }
    return
}
