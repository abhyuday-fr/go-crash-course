package main

// goroutines are a way to launch multiple functions and have them execute concurrently
// concurrency              !=    parallel execution
// (one core - switches fast)    (multiple core - each execute their own tasks)


import (
	"fmt"
	"time"
	"sync"
)

var m =  sync.RWMutex{} // Multiple threads accessing and modifying the same memory block at the same time can lead to a bunch of issues so we use Mutex to fix avoid that
var wg = sync.WaitGroup{} // Waitgroups are like counters, whenever we spawn a new go routine we make sure to .Add(1) to the counter and inside the function we call the .Done() fnctn at the end and then add the .Wait() method after all the callings i.e loop
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main(){
	t0 := time.Now()
	for i:=0; i<len(dbData); i++{ // normally we are calling the data sequentially one by one which will take approx 5 seconds
		// better way to do this is by calling concurrently, so we add the 'go' keyword in front of the function calling
		
		wg.Add(1) //counter
		
		go dbCall(i) // when we add 'go' keyword, our program won't wait for this fnctn to get complete rather it will keep moving on to the next step in the loop, so it it just ends without doing anything.. to fix that we use sync package and its functions
	}
	wg.Wait() // the counters has become 0
	
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("The results are %v", results)
}

func dbCall(i int){
	// Simulate DB Call Delay
	var delay float32 = 2000 // 2 seconds delay
	time.Sleep(time.Duration(delay)*time.Millisecond)
	
	fmt.Println("\nThe result from the database is:", dbData[i])
	
	save(dbData[i])

	log()

	wg.Done() // Decrements the code
}

func save(result string){
	m.Lock() // when a go routine reaches this lock method a check is performed to see if a lock has already been set by another go routine, if it has it will wait here until the lock is released and then set the lock itself once
	results = append(results, result) // placing mutex's lock and unlock around the part of our code which accesses the result slice
	m.Unlock() // the lock is realsed after the accessing the result slice using the unlock method and now go routined can obtain a lock as neededk
}

func log(){
	m.RLock()
	fmt.Printf("\nThe Current Results are: %v", results)
	m.RUnlock()
}
/*
TLDR;
1. Add go keyword before the db calling
2. use Waitgroup from sync package, put .Add(1), .Wait(), ,.Done() methods at appropriate locations
3. use RWMutext from sync package, put .Lock() and .Unlock() around the code part where we access the results slice
4. use .Rlock() and .RUnlock() to log
*/