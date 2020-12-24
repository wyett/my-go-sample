package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/**
 * @author     : wyettLei
 * @date       : Created in 2020/12/24 11:13
 * @description: TODO
 */

// job
type Job struct {
	id        int // id of number
	randomNum int // random num
}

type Result struct {
	job          Job
	sumOfDigests int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

// calcute digest
func cal(num int) int {
	//fmt.Printf("calcute sum of num")
	sum := 0
	n := num
	for n != 0 {
		lastPosNum := n % 10
		sum += lastPosNum
		n /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

// worker
func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, cal(job.randomNum)}
		results <- output
	}
	wg.Done()
}

// create work Pool
func createWorkPool(numOfWorker int) {
	var wg sync.WaitGroup
	for i := 1; i <= numOfWorker; i++ {
		wg.Add(1)
		go worker(&wg)
		//jobs <- Job{i, rand.Intn(1000)}
	}
	wg.Wait()
	close(results)
}

// allocate job
func allocate(numOfJobs int) {
	for i := 1; i < numOfJobs; i++ {
		randomNum := rand.Intn(999)
		jobs <- Job{i, randomNum}
	}
	close(jobs)
}

// result
func result(done chan bool) {
	for result := range results {
		fmt.Printf("id: %v, randomNum: %v, sumOfDigest:%v\n", result.job.id, result.job.randomNum, result.sumOfDigests)
	}
	done <- true
}

// workerPool Func
func WorkerPoolMain() {
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 10
	createWorkPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
