package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup // 1

//canale singolo, goroutine singola, una scrittura, una lettura.

//DoMultiply : a simple func
func DoMultiply(x, y int) {
	defer wg.Done() // 3 <---- make the main to wait
	fmt.Printf("Result: %d\n", x*y)
}

//I programmi Go terminano quando termina la funzione main
//è prassi comune aspettare che tutte le goroutine finiscano.
//Una soluzione comune per questo è utilizzare un oggetto sync.WaitGroup .

/* Uso WaitGroup in ordine di esecuzione:

    1.	Dichiarazione di variabile globale: Rendendolo globale è il modo più semplice per renderlo visibile a tutte le funzioni e i metodi.
	2.	Aumentare il contatore. Questo deve essere fatto nella goroutine principale perché non vi è alcuna garanzia che una goroutine appena
		avviata venga eseguita prima delle 4 a causa delle garanzie del modello di memoria.
	3.	Diminuendo il contatore. Questo deve essere fatto all'uscita di una goroutine. Usando una chiamata differita, ci assicuriamo che venga chiamata
		ogni volta che la funzione termina , indipendentemente da come finisce.
	4.  Aspettando che il contatore raggiunga 0. Questo deve essere fatto nella goroutine principale per impedire che il programma esca prima
		che tutte le goroutine siano finite. I parametri vengono valutati prima di iniziare una nuova goroutine .
		Quindi è necessario definire i loro valori esplicitamente prima di wg.Add(10) modo che il codice eventualmente panico non aumenti il ​​contatore.
		Aggiungendo 10 elementi al WaitGroup, quindi attenderà 10 elementi prima che wg.Wait restituisca il controllo alla goroutine main() .
		Qui, il valore di i è definito nel ciclo for. */

func main() {

	wg.Add(2) // 2

	// create new channel of type string
	ch := make(chan string)

	// start new anonymous goroutine
	go func() {
		time.Sleep(time.Second)
		// send "Hello World" to channel
		ch <- "Hello World"
	}()
	// read from channel
	msg, ok := <-ch
	fmt.Printf("msg='%s', ok='%v'\n", msg, ok)

	go DoMultiply(1, 2) // first execution, non-blocking
	go DoMultiply(3, 4) // second execution, also non-blocking

	wg.Wait() // 4
}
