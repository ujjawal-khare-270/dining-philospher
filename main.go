package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

var noOfEats = 4
var eatTime = 3 * time.Second
var thinkTime = 2 * time.Second
var numOfForks = 5

var philosophers = []Philosopher{
	{name: "A", leftFork: 4, rightFork: 0},
	{name: "B", leftFork: 0, rightFork: 1},
	{name: "C", leftFork: 1, rightFork: 2},
	{name: "D", leftFork: 2, rightFork: 3},
	{name: "E", leftFork: 3, rightFork: 4},
}

func DinePhilosopher(philosopher Philosopher, wg *sync.WaitGroup, forksMutexMap map[int]*sync.Mutex) {
	defer wg.Done()

	for i := 0; i < noOfEats; i++ {
		if philosopher.rightFork < philosopher.leftFork {
			l := philosopher.leftFork
			r := philosopher.rightFork

			forksMutexMap[l].Lock()
			forksMutexMap[r].Lock()
			fmt.Println(philosopher.name, "has started eating for the ", i+1, " time")
			time.Sleep(eatTime)

			forksMutexMap[l].Unlock()
			forksMutexMap[r].Unlock()
		} else {
			l := philosopher.leftFork
			r := philosopher.rightFork

			forksMutexMap[r].Lock()
			forksMutexMap[l].Lock()
			fmt.Println(philosopher.name, "has started eating for the ", i+1, " time")
			time.Sleep(eatTime)

			forksMutexMap[l].Unlock()
			forksMutexMap[r].Unlock()
		}
		time.Sleep(thinkTime)
		fmt.Println(philosopher.name, "has done thought")
	}
	fmt.Println("Hurray", philosopher.name, "has done eating")
}
func main() {
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	var forksMutexMap = make(map[int]*sync.Mutex)
	for i := 0; i < numOfForks; i++ {
		forksMutexMap[i] = &sync.Mutex{}
	}

	for i := 0; i < len(philosophers); i++ {
		go DinePhilosopher(philosophers[i], wg, forksMutexMap)
	}

	wg.Wait()
}
