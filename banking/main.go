package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

/*
MUTEX : Solo una goroutine per volta può accedere a una variabile senza conflitti.

Questo concetto si chiama mutua esclusione, e il nome della struttura che offre questa proprietà è mutex.

La libreria standard di Go fornisce mutua esclusione con sync.Mutex
e i suoi due metodi: Lock e Unlock

Possiamo definire una parte di codice da eseguire in mutua esclusione
semplicemente circondandola con una chiamata a Lock e Unlock.

Possiamo anche usare defer per assicurarci che il mutex verrá sbloccato.
*/

//Account represents the bank account
type Account struct {
	currentMoney int64
	active       bool
	sync.Mutex
}

//OpenAccount : a way to create a bank-account with an initial deposit
func OpenAccount(initialDeposit int64) *Account {

	if initialDeposit < 0 {
		return nil
	}

	return &Account{currentMoney: initialDeposit, active: true}
}

//Close : a way to close the account
func (acc *Account) Close() (payout int64, ok bool) {
	acc.Lock()

	defer acc.Unlock()
	defer wg.Done()

	if acc.active {
		payout, ok = acc.currentMoney, true
		acc.currentMoney, acc.active = 0, false
	}

	return payout, ok
}

//Balance : a way to get the balance
func (acc *Account) Balance() (balance int64, ok bool) {
	acc.Lock()
	defer acc.Unlock()
	defer wg.Done()
	if !acc.active {
		return 0, false
	}
	return acc.currentMoney, true
}

// Deposit is a way to store money in an account
func (acc *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	acc.Lock()

	defer acc.Unlock()
	defer wg.Done()

	newBalance = acc.currentMoney + amount
	if !acc.active || newBalance < 0 {
		return 0, false
	}
	acc.currentMoney = newBalance

	return newBalance, true
}

//TODO : withdraw()

//Withdraw : a way to withdraw money from the bankaccount
func (acc *Account) Withdraw(withdraw int64) (newBalance int64, ok bool) {

	acc.Lock()
	defer acc.Unlock()
	defer wg.Done()

	newBalance = acc.currentMoney - withdraw
	if !acc.active || newBalance < 0 {
		return 0, false
	}
	acc.currentMoney = newBalance

	return newBalance, true
}

/*
//sleep time
func main() {

	acc := OpenAccount(7)

	fmt.Println(acc.currentMoney)

	go acc.Deposit(6)
	go acc.Withdraw(3)

	time.Sleep(time.Second * 1)

	fmt.Println(acc.currentMoney)
}
*/

func main() {

	wg.Add(2)

	acc := OpenAccount(7)

	fmt.Println(acc.currentMoney)

	go acc.Deposit(6)
	go acc.Withdraw(3)

	wg.Wait()

	fmt.Println(acc.currentMoney)
}
