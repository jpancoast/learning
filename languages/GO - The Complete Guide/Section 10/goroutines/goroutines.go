package main

import (
	"fmt"
	"time"
)

func greet(greeting string, doneChan chan bool, errorChan chan error) {

	err := fmt.Errorf("error")
	if err != nil {
		errorChan <- err
		return
	}

	fmt.Println(greeting)

	doneChan <- true
}

func slowGreet(greeting string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(greeting)

	doneChan <- true //the <- sends data to a channel
}

/*
 * defer waits to run right before the function completed
 */
func deferFunc() {
	defer fmt.Println("I'm done")

	fmt.Println("In defer function")
}

// the 'go ' means to run as a goroutine, and run concurrently
func main() {
	done := make(chan bool, 4)
	errorChans := make(chan error, 4)

	go greet("first greeting", done, errorChans)
	go greet("second greeting", done, errorChans)
	go slowGreet("slow greeting", done)
	go greet("4th greeting", done, errorChans)

	//select is a control structure for channels
	for i := 0; i < 4; i++ {
		select {
		case doneChan := <-done:
			fmt.Println(doneChan)
		case err := <-errorChans:
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	deferFunc()
	/*
		for doneChan := range done {
			fmt.Println(doneChan)
			// Problem here is this gives an error, but so far he hasn't told us how to fix it, other than to close the channel in the longest running function, ie, 'slowGreet'
		}*/

	// dones := make([]chan bool, 4)
	// dones[0] = make(chan bool)
	// dones[1] = make(chan bool)
	// dones[2] = make(chan bool)
	// dones[3] = make(chan bool)

	// go greet("first greeting", dones[0])
	// go greet("second greeting", dones[1])
	// go slowGreet("slow greeting", dones[2])
	// go greet("4th greeting", dones[3])

	// for _, done := range dones {
	// 	<-done
	// }

	// // //This waits for all 4 to finish.
	// // <-done //This is WAITING for the channel to send a value
	// // <-done
	// // <-done
	// // <-done

	// //This is the same as:
	// //for {
	// //	if val, ok := <-done; ok {
	// //		fmt.Println(val)
	// //		break
	// //	}
	// //}
}
