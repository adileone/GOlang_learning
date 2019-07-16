package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
PUNTATORI * ----
Un puntatore è un riferimento alla posizione in memoria dove il valore è immagazzinato e non al valore stesso, cioè punta a qualcos’altro.
In questo modo si può quindi modificare la variabile.


!!!
Quando chiamiamo una funzione che prende un argomento la funzione riceve una copia dell’argomento,
il valore originale della variabile all’uscita della funzione resta invariato: a meno che non fai ritornare l'oggetto

per modificare il valore dell'oggetto senza return utilizzi come ricettore il puntatore, non passato come valore

In Go i puntatori sono raramente usati con i tipi built-in (predefiniti) ma estremamente utili nelle structs.

!!!
Ci sono due ragioni per scegliere un ricettore puntatore.
Il primo è che il metodo può così modificare il valore al quale il puntatore sta puntando.
Il secondo è per evitare di copiare il valore per ogni chiamata al metodo. Questo può essere più efficiente se il ricettore è una struttura di grandi dimensioni.


new prende un tipo come argomento, alloca la quantità di memoria necessaria a contenere quel tipo e ritorna un puntatore a quel tipo.
In Go, grazie al garbace collector automatico, vengono meno tutti quelle attenzioni richiesta da new in linguaggi come il C.
*/

//WordCount : goTour
func WordCount(s string) map[string]int {

	a := strings.Fields(s)
	m := make(map[string]int)

	for _, word := range a {
		m[word]++
	}
	return m
}

//Fibonacci : closure func
func Fibonacci() func() int {

	a, b := 0, 1

	return func() int {

		sum := a
		a, b = b, a+b
		return sum
	}
}

//Go non ha classi -- simil oggetti e metodi relativi
//Un metodo è una funzione con uno speciale argomento ricettore.
//Il ricettore del metodo appare nella propria lista di argomenti tra la parola chiave func e il nome del metodo.

//Vertex : example of struct  ---- Oggetto?
type Vertex struct {
	X, Y float64
}

//Abs : valore assoluto di un numero
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//IPAddr : array made of 4 bytes
type IPAddr [4]byte

//String : stringer example
func (ip IPAddr) String() string {
	return strconv.Itoa(int(ip[0])) + "." +
		strconv.Itoa(int(ip[1])) + "." +
		strconv.Itoa(int(ip[2])) + "." +
		strconv.Itoa(int(ip[3]))
}

func main() {

	m := WordCount("ciao mondo ciao alessandro hello mondo ciao")
	for el := range m {
		fmt.Printf("word :" + el + ", num : " + strconv.Itoa(m[el]) + "\n")
	}

	f := Fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	point := Vertex{3, 4}
	fmt.Println(point.Abs())

	hosts := map[string]IPAddr{
		"loopback ": {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Println(name, ":", ip)
	}
}
