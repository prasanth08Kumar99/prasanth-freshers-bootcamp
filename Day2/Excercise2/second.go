package main
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
var rate int
func computeRating(ch chan int , wg *sync.WaitGroup)  {
	defer wg.Done()
	rating := suckRating(ch, wg)
	time.Sleep(time.Duration(rating*10e7))
	rate += rating
}

func pumpRating(ch chan int) {
	ch <- rand.Intn(10)
}

func suckRating(ch chan int, wg *sync.WaitGroup) int {
	defer wg.Done()
	return <-ch
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	rate = 0
	for i:=1 ; i<=200; i++ {
		wg.Add(2)
		go pumpRating(ch)
		go computeRating (ch,&wg)
	}
	wg.Wait()
	avgrate := rate/200;
	fmt.Println("The average rating of the teacher is ",avgrate)
}
