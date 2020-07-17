package main

import (
  "fmt"
  "buffio"
  "os"
)

func main()  {
  scanner := buffio.NewScanner(os.Stdin)
  scanner.Scan()
  input := scanner.Text()

  fmt.Println("you typed: %q", input)

  // var x int = 1

  // fmt.Println(x)
}
