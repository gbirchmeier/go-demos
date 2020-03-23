package main

import(
  "fmt"
  "grant/helloworld/demopackage"
)

func main() {
  fmt.Println("Hello, world.")
  fmt.Println(demopackage.ReverseRunes("hello world"))
}
