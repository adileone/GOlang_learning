package main

import (
	"fmt"
	"math"
	"strconv"
)

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

func forCycle() string {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	return strconv.Itoa(sum)

}

//Ricettori puntatori

/*
Metodi con ricettori puntatori possono modificare il valore al quale puntano.
Con un ricettore di tipo valore, il metodo Scale opera su una copia del valore originale Vertex.
*/

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//Vertex con e senza puntatore comportamenti differenti
func (v Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	fmt.Println(forCycle())
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
