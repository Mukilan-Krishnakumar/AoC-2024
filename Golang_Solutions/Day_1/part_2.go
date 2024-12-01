package main

import(
  "fmt"
  "os"
  "bufio"
  "regexp"
  "strconv"
)

func get_frequency_map(og_list []int) map[int]int {
  var frequency map[int]int
  frequency = make(map[int]int)
  for _, number := range og_list{
    value, _ := frequency[number]
    value += 1
    frequency[number] = value
  }
  return frequency
}

func get_similarity_score(og_list []int, frequency_map map[int]int) int{
  var similarity int = 0
  for _, number := range og_list{
    value, ok := frequency_map[number]
    if ok{
      similarity += value * number
    }
  } 
  return similarity
}

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
  var frequency map[int]int
  frequency = get_frequency_map(list_2)
  var similarity int = 0
  similarity = get_similarity_score(list_1, frequency)
  fmt.Println(similarity)
}