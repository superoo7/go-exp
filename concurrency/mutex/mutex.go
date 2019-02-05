package mutex

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock    sync.Mutex
	counter = 0
)

func Mutex() {
	for i := 0; i < 20; i++ {
		go updateCounter()
		time.Sleep(time.Millisecond * 10)
	}
}

func updateCounter() {
	lock.Lock()
	defer lock.Unlock()
	counter++
	fmt.Printf("%d\n", counter)
}
