package main

import (
	"fmt"
	"time"
)

func onetotwenty() {
	fmt.Println("start onetotwenty goroutine!")
	for i := 0; i < 20; i++ {
		fmt.Printf("Goroutine iteration %d\n", i)
	}
	fmt.Println("end onetotwenty goroutine!")
}

func twentytofourty() {
	fmt.Println("start twentytofourty goroutine!")
	for i := 20; i < 40; i++ {
		fmt.Printf("Goroutine iteration %d\n", i)
	}
	fmt.Println("end twentytofourty goroutine!")
}

func main() {
	go onetotwenty()
	fmt.Println("Start Main function ")
	for i := 0; i < 20; i++ {
		fmt.Printf("Main function iteration %d\n", i)
	}
	fmt.Println("Main function completed.")
	// time.Sleep(100 * time.Millisecond)
	go twentytofourty()
	time.Sleep(100 * time.Millisecond)
}
