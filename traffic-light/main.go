package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var passwords []int
var stopProccess chan struct{}

func main() {
	generate_passwords()
	trafficLight := make(chan struct{}, 2000)
	stopProccess = make(chan struct{}, 1)

	for {
		if len(passwords) == 0 {
			break
		}
		password := passwords[0]
		passwords = passwords[1:]
		
		wg.Add(1)
		trafficLight <- struct{}{}
		
		go check_password(password, trafficLight)
		fmt.Println("Cracking password", password)
	}

	wg.Wait()
}

func check_password(password int, trafficLight <-chan struct{}) {
	defer wg.Done()

	select {
	case <-stopProccess:
		return
	default:
		time.Sleep(5 * time.Second) //Simulate response time to resquest.
	}

	if(password == 4321){
		fmt.Println("Password cracked!")

		close(stopProccess)
	}

	<-trafficLight
}

func generate_passwords() {
	for i := 0; i <= 5000; i++ {
		passwords = append(passwords, i)
	}
}