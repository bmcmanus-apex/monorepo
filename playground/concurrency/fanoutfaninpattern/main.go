package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func repeatFunc[T any, K any](done <-chan K, fn func() T) <-chan T {
	stream := make(chan T)

	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}

	}()

	return stream
}

func take[T any, K any](done <-chan K, stream <-chan T, n int) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case out <- <-stream:
			}
		}
	}()

	return out
}

func primeFinder(done <-chan struct{}, randIntStream <-chan int) <-chan int {
	isPrime := func(randomInt int) bool {
		for i := randomInt - 1; i > 1; i-- {
			if randomInt%i == 0 {
				return false
			}
		}
		return true
	}

	primes := make(chan int)

	go func() {
		defer close(primes)
		for {
			select {
			case <-done:
				return
			case randomInt := <-randIntStream:
				if isPrime(randomInt) {
					primes <- randomInt
				}
			}
		}
	}()

	return primes
}

func fanIn[T any](done <-chan struct{}, channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup

	fannedInStream := make(chan T)

	transfer := func(c <-chan T) {
		defer wg.Done()

		for i := range c {
			select {
			case <-done:
				return
			case fannedInStream <- i:
			}
		}
	}

	for _, c := range channels {
		wg.Add(1)
		go transfer(c)
	}

	go func() {
		wg.Wait()
		close(fannedInStream)
	}()

	return fannedInStream
}

func main() {
	start := time.Now()

	done := make(chan struct{})
	defer close(done)

	randNumFetcher := func() int {
		return rand.Intn(500000000)
	}

	randIntStream := repeatFunc(done, randNumFetcher)

	// naive slow approach that requires synchronicity
	//primeStream := primeFinder(done, randIntStream)
	//for rando := range take(done, primeStream, 10) {
	//	fmt.Println(rando)
	//}

	// fan out
	CPUCount := runtime.NumCPU()
	primeFinderChannels := make([]<-chan int, CPUCount)
	for i := 0; i < CPUCount; i++ {
		primeFinderChannels[i] = primeFinder(done, randIntStream)
	}

	// fan in
	fannedInStream := fanIn(done, primeFinderChannels...)

	for rando := range take(done, fannedInStream, 10) {
		fmt.Println(rando)
	}
	fmt.Println(time.Since(start))
}
