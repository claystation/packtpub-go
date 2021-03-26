package main

import "fmt"

func main() {
  
  visits := 15
  fmt.Println("First visit: ", visits == 1)

  fmt.Println("Return visits: ", visits != 1)
  fmt.Println("Silver member: ", visits >= 10 && visits <= 20)

  fmt.Println("Gold member: ", visits >= 21 && visits <= 30)
  fmt.Println("Platinum member: ", visits > 30)
}
