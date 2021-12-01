package main

import (
	"context"
	"log"
	"math"
	"sync"
)

const nPresents = 36000000

type presentsFunc func(n int) int

func presentsToHousep2(n int) int {
	presents := 0
	for i := 1; i < int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			if i <= 50 {
				presents += n / i
			}
			if n/i <= 50 {
				presents += i
			}
		}
	}

	return presents * 11
}

func presentsToHouse(n int) int {
	presents := 0
	for i := 1; i <= int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 {
			presents += i
			presents += n / i
		}
	}

	return presents * 10
}

func worker(ctx context.Context, input <-chan int, result chan<- int, fn presentsFunc) {
	for {
		select {
		case n, ok := <-input:
			if !ok {
				return
			}

			if fn(n) >= nPresents {
				result <- n
			}

		case <-ctx.Done():
			return
		}
	}
}

func producer(ctx context.Context, start int) <-chan int {
	i := start
	ch := make(chan int, 100)
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			case ch <- i:
			}
			i++
		}
	}()

	return ch
}

func p1() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	workers := 8
	ch := producer(ctx, 1)
	result := make(chan int, workers)
	wg := sync.WaitGroup{}
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			worker(ctx, ch, result, presentsToHouse)
		}()
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	minN := 0
	for n := range result {
		cancel()
		if minN == 0 || n < minN {
			minN = n
		}
	}

	log.Printf("part1: %d", minN)
}

func p2() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	workers := 8
	ch := producer(ctx, 50)
	result := make(chan int, workers)
	wg := sync.WaitGroup{}
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			worker(ctx, ch, result, presentsToHousep2)
		}()
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	minN := 0
	for n := range result {
		cancel()
		if minN == 0 || n < minN {
			minN = n
		}
	}

	log.Printf("part2: %d", minN)
}

func main() {
	p1()
	p2()
}
