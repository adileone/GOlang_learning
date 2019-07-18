package main

import "fmt"

/*
Go ha solo un costrutto ciclico, il ciclo for.

Il for ha tre componenti separati da punto e virgola:

	l'inizializzazione: eseguita una volta sola prima del ciclo
	la condizione (booleana): valutata all'inizio di ogni ripetizione del ciclo
	l'incremento/decremento: eseguito alla fine di ogni ripetizione del ciclo

	L'inizializzazione è spesso fatta con una dichiarazione breve di variabile, e le variabili create in questo componente del ciclo sono visibili solo all'interno del ciclo for.
    Il ciclo si fermerá solo quando la condizione booleana sará valutata falsa.

Nota: diversamente da altri linguaggi come C, Java, o Javascript i tre componenti del costrutto ciclico non sono tra parentesi ( ) e le parentesi { } sono obbligatorie.
*/

func main() {

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
