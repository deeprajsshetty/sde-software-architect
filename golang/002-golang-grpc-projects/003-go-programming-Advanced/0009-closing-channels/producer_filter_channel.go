package main

import "fmt"

// Send only channel.
func producer(data chan <- int)  {
	for i := range 100 {
		data <- i
	}
	close(data)
}

// Receive and send channels.
func filter(data <-chan int, filtered chan <- int)  {
	for d := range data {
		if d%2 == 0 {
			filtered <- d
		}
	}
	close(filtered)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go producer(ch1)
	go filter(ch1, ch2)

	for filterData := range ch2 {
		fmt.Printf("Filtered data is %d\n", filterData)
	}
}

/*
------Output-------
Filtered data is 0
Filtered data is 2
Filtered data is 4
Filtered data is 6
Filtered data is 8
Filtered data is 10
Filtered data is 12
Filtered data is 14
Filtered data is 16
Filtered data is 18
Filtered data is 20
Filtered data is 22
Filtered data is 24
Filtered data is 26
Filtered data is 28
Filtered data is 30
Filtered data is 32
Filtered data is 34
Filtered data is 36
Filtered data is 38
Filtered data is 40
Filtered data is 42
Filtered data is 44
Filtered data is 46
Filtered data is 48
Filtered data is 50
Filtered data is 52
Filtered data is 54
Filtered data is 56
Filtered data is 58
Filtered data is 60
Filtered data is 62
Filtered data is 64
Filtered data is 66
Filtered data is 68
Filtered data is 70
Filtered data is 72
Filtered data is 74
Filtered data is 76
Filtered data is 78
Filtered data is 80
Filtered data is 82
Filtered data is 84
Filtered data is 86
Filtered data is 88
Filtered data is 90
Filtered data is 92
Filtered data is 94
Filtered data is 96
Filtered data is 98
------Output-------
*/