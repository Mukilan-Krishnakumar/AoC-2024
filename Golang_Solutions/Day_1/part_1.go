package main

import(
  "fmt"
  "os"
  "bufio"
  "regexp"
  "strconv"
  "slices"
  "math"
)

func main(){
  re := regexp.MustCompile(`(\d+)   (\d+)`)
  file, err := os.Open("./input.txt")
  if err != nil{
    panic(err)
  }
  defer file.Close()

  var list_1 []int
  var list_2 []int
  list_1 = make([]int, 10, 10)
  list_2 = make([]int, 10, 10)
  scanner := bufio.NewScanner(file)
  for scanner.Scan(){
    text := scanner.Text()
    _, err := regexp.Match(`(\d+)   (\d+)`, []byte(text))
    if err != nil{
      panic(err)
    }
    submatches := re.FindSubmatch([]byte(text))
    l1, err := strconv.Atoi(string(submatches[1]))
    if err != nil{
      panic(err)
    }
    l2, err := strconv.Atoi(string(submatches[2]))
    if err != nil{
      panic(err)
    }
    list_1 = append(list_1, l1)
    list_2 = append(list_2, l2)
    
  }
  slices.Sort(list_1)
  slices.Sort(list_2)
  var diff int
  for i:= 0; i < len(list_1); i++{
    diff += int(math.Abs(float64(list_1[i] - list_2[i])))
  }
  fmt.Println(diff)
}
