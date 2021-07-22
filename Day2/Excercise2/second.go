package main
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
var rate int
func computerating(ch chan int , wg *sync.WaitGroup)  {
	defer wg.Done()
	rating := suckrating(ch, wg)
	time.Sleep(time.Duration(rating*10e7))
	rate += rating
}

func pumprating(ch chan int) {
	ch <- rand.Intn(10)
}

func suckrating(ch chan int, wg *sync.WaitGroup) int {
	defer wg.Done()
	return <-ch
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	rate = 0
	for i:=1 ; i<=200; i++ {
		wg.Add(2)
		go pumprating(ch)
		go computerating (ch,&wg)
	}
	wg.Wait()
	avgrate := rate/200;
	fmt.Println("The average rating of the teacher is ",avgrate)
}
