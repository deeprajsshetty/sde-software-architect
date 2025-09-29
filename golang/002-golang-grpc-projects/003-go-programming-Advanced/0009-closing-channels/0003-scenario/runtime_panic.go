package main

// Closing channel multile times by mistake.
// Runtime error.
func main() {
	ch := make(chan int)

	ch <- 1

	go func(){
		close(ch)
		close(ch)
	}()

}

/*
------------------------------Output------------------------------------
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
	/Users/deeprajsshetty/go/src/github.com/deeprajsshetty/sde-software-architect/golang/002-golang-grpc-projects/003-go-programming-Advanced/0009-closing-channels/0003-scenario/runtime_panic.go:8 +0x38
exit status 2
------------------------------Output------------------------------------
*/