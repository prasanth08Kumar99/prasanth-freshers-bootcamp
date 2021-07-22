package main
import (
	"fmt"
	"sync"
)
var (
	mutex sync.Mutex
	balance int
)
func Deposit (amount int , wg *sync.WaitGroup){
	defer wg.Done()
	mutex.Lock()
	fmt.Println(amount,"amount is being deposited")
	balance += amount
	mutex.Unlock()
}
func Withdraw (amount int , wg *sync.WaitGroup){
	defer wg.Done()
	mutex.Lock()
	if amount > balance {
		fmt.Println("Insufficient Balance")
	} else {
		fmt.Println(amount,"amount is withdrawn")
		balance -= amount
	}
	mutex.Unlock()
}
func main(){
	balance = 5000
	var wg sync.WaitGroup
	wg.Add(3)
	Withdraw(3000,&wg)
	Deposit(1000,&wg)
	Withdraw(4000,&wg)
	wg.Wait()
	fmt.Println("The current balance is ",balance)
}
