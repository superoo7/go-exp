package channel

import (
	"fmt"
	"math/rand"
	"time"
)

// Worker to do stuff
type Worker struct {
	id int
}

func (w Worker) process(c chan int) {
	for {
		data := <-c
		fmt.Printf("worker %d  got %d\n", w.id, data)
	}
}

// Channel channel
func Channel() {
	c := make(chan int)
	for i := 0; i < 4; i++ {
		worker := Worker{id: i}
		go worker.process(c)
	}

	// create while loop
	for {
		c <- rand.Int()
		time.Sleep(time.Millisecond * 50)
	}
}
