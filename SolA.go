
package main

import (
	"fmt"
	"strconv"
)

const (
	PHILOSOPHERS = 5
	FORKS        = 5
)

func philosopher(id int, left_fork chan struct{}, right_fork chan struct{}) {
	// Define o filósofo canhoto
	isLefty := id == 0

	for {
		fmt.Println(strconv.Itoa(id) + " senta")
		
		// Tenta pegar o primeiro garfo na ordem adequada
		if isLefty {
			<-right_fork
			fmt.Println(strconv.Itoa(id) + " pegou direita")
			<-left_fork
			fmt.Println(strconv.Itoa(id) + " pegou esquerda")
		} else {
			<-left_fork
			fmt.Println(strconv.Itoa(id) + " pegou esquerda")
			<-right_fork
			fmt.Println(strconv.Itoa(id) + " pegou direita")
		}
		
		fmt.Println(strconv.Itoa(id) + " come")
		left_fork <- struct{}{} // devolve
		right_fork <- struct{}{}
		fmt.Println(strconv.Itoa(id) + " levanta e pensa")
	}
}

func main() {
	var fork_channels [FORKS]chan struct{}
	for i := 0; i < FORKS; i++ {
		fork_channels[i] = make(chan struct{}, 1)
		fork_channels[i] <- struct{}{} // no inicio garfo esta livre
	}
	for i := 0; i < (PHILOSOPHERS); i++ {
		fmt.Println("Filosofo " + strconv.Itoa(i))
		go philosopher(i, fork_channels[i], fork_channels[(i+1)%PHILOSOPHERS])
	}
	var blq chan struct{} = make(chan struct{})
	<-blq
}
