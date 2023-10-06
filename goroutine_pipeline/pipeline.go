package main

import (
	"fmt"
)

// CHANNEL
// func main() {
// 	charChannel := make(chan string, 3)
// 	chars := []string{"a", "b", "c"}

// 	for _, char := range chars {
// 		select {
//		// pq o canal está dentro de um select?
// 		case charChannel <- char:
// 		}
// 	}
// 	close(charChannel)

// 	for result := range charChannel {
// 		fmt.Println(result)
// 	}
// }


// // RUNNING INIFITE
// func main() {
// 	go func() {
// 		for {
// 			select {
// 			default: 
// 				fmt.Println("DOING WORK")
// 			}
// 		}
// 	}()
// 	time.Sleep(time.Duration(5) * time.Second)
// }

// //FOR SELECT PATTERN
// func doWork(done <- chan bool) {
// 	for {
// 		select {
// 		case <- done:
// 			return
// 		default:
// 			fmt.Printf("DOING WORK\n")
// 		}
// 	}
// }

// func main() {
// 	done := make(chan bool)
// 	go doWork(done)
// 	time.Sleep(time.Duration(3) * time.Second)
// 	close(done)
// }



// GOLANG GOROUTINES PIPELINES

func sliceToChannel(nums []int) <-chan int{
	out := make(chan int, 50)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <- chan int) <- chan int {
	out := make(chan int, 50)
	go func() {
		for n := range in {
			out <- n*n
		}
		close(out)
	}()
	return out
}

func main() {
	// input
	nums := []int{2,3,5,8,6}
	// stage1
	dataChannel := sliceToChannel(nums)
	// stage2
	finalChannel := square(dataChannel)
	// stage3
	for n := range finalChannel {
		fmt.Println(n)
	}
}