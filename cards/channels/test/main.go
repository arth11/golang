package main

import (
	    "fmt"
	    "sync"
	)
	
func visit(numbers []int, callback func(number int, arg []int, wg *sync.WaitGroup, c chan []int), goRoutinesNumber int) {
    var wg sync.WaitGroup
    c := make(chan []int, 100)
    wg.Add(goRoutinesNumber)
    for i := 0; i < goRoutinesNumber; i++ {
        slice := make([]int, len(numbers))
        copy(numbers, slice)
        go callback(i, slice, &wg, c)
    }
    wg.Wait()
    close(c)
    for i := range c {
        fmt.Println(i)
    }
}

func test(number int, arg []int, wg *sync.WaitGroup, c chan []int) {
    defer wg.Done()
    for i := 0; i < len(arg); i++ {
        arg[i] = number
    }
    c <- arg
    fmt.Println("Gotoutine number ", number, " DONE")
}

func main() {
    length := 100
    goRoutinesNumber := 10
    slice := make([]int, length)
    for i := 0; i < length; i++ {
        slice[i] = i
    }
    visit(slice, test, goRoutinesNumber)
} 