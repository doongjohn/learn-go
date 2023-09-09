package main

// creating a project:
// go mod init <project-url>

// build:
// go build -o <output-file-path>

// update dependencies:
// go get -u ./...

// run:
// go run <source-file-path>

// test:
// go test -v ./... -count=1

// TODO:
// - learn error handling
// - learn context package

import (
	"fmt"
	"sync"
	"time"
)

type Position struct {
    x float32
    y float32
}

type GameObject struct {
    Position
    name string
}

func main() {
    hello := 10
    hello2, hello := 20, 30
    fmt.Println(hello, hello2)

    num1, err := errFunc1()
    if err != nil {
        fmt.Println("it's err")
    }
    num2, err := errFunc1()
    if err != nil {
        fmt.Println("it's err")
    }
    _ = num1
    _ = num2

    // struct embeding
    obj := GameObject{}
    fmt.Printf("obj.name = %s\n", obj.name)
    fmt.Printf("obj.x = %f\n", obj.x)
    fmt.Printf("obj.y = %f\n", obj.y)

	// rune
	{
		for pos, rune := range "👍🏻안유찬" {
			fmt.Printf("rune %c starts at byte position %d\n", rune, pos)
		}
	}

	// closure
	{
		intSeq := func() func() int {
			i := 0
			return func() int {
				i++
				return i
			}
		}

		nextInt1 := intSeq()
		fmt.Println(nextInt1())
		fmt.Println(nextInt1())
		fmt.Println(nextInt1())

		nextInt2 := intSeq()
		fmt.Println(nextInt2())
		fmt.Println(nextInt2())
	}

	// wait group
	{
		const works = 5
		fmt.Printf("Do %d works using goroutine\n", works)

		var wg sync.WaitGroup
		wg.Add(works)

		someWork := func() {
			time.Sleep(time.Second * 1)
			fmt.Println("Some work is done!")
		}

		start := time.Now()
		for i := 0; i < works; i++ {
			go func() {
				someWork()
				wg.Done()
			}()
		}

		wg.Wait()
		fmt.Printf("took: %d\n", time.Since(start))
	}

	// channel
	{
		messages := make(chan string)

		sendHello := func(name string, msg chan string) {
			time.Sleep(time.Millisecond * 500)
			msg <- "Hello, " + name
		}

		go func() {
			defer close(messages)

			names := [...]string{"John", "Tom", "Ben"}
			for _, name := range names {
				sendHello(name, messages)
			}
		}()

		// for {
		// 	hello, isOpen := <-messages
		// 	if !isOpen {
		// 		break
		// 	}
		// 	fmt.Println(hello)
		// }

		for hello := range messages {
			fmt.Println(hello)
		}
	}

	// multiple channels
	{
		var wg sync.WaitGroup
		wg.Add(2)

		chDone := make(chan struct{})
		//                  ^^^^^^^^ --> empty struct
		ch1 := make(chan string)
		ch2 := make(chan string)

		defer func() {
			close(chDone)
			close(ch1)
			close(ch2)
		}()

		go func() {
			wg.Wait()
			chDone <- struct{}{}
		}()

		go func() {
			for i := 0; i < 3; i++ {
				time.Sleep(time.Second * 1)
				ch1 <- "1 sec"
			}
			wg.Done()
		}()

		go func() {
			for i := 0; i < 3; i++ {
				time.Sleep(time.Second * 2)
				ch2 <- "2 sec"
			}
			wg.Done()
		}()

		loop := true
		for loop {
			select {
			case msg := <-ch1:
				fmt.Println(msg)
			case msg := <-ch2:
				fmt.Println(msg)
			case <-chDone:
				loop = false
			}
		}
	}
}

func errFunc1() (int, error) {
	return 100, nil
}

func errFunc2() (int, error) {
	return 100, nil
}
