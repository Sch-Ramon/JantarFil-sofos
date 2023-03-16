package main

import (
    "fmt"
    "strconv"
)

const (
    PHILOSOPHERS = 5
    FORKS        = 5
)

func philosopher(id int, first_fork chan struct{}, second_fork chan struct{}, second_fork_available chan bool) {
    for {
        fmt.Println(strconv.Itoa(id) + " senta")
        <-first_fork // pega primeiro garfo
        fmt.Println(strconv.Itoa(id) + " pegou direita")
        select {
        case <-second_fork: // pega segundo garfo se disponível
            fmt.Println(strconv.Itoa(id) + " come")
            first_fork <- struct{}{} // devolve primeiro garfo
            second_fork_available <- true // segundo garfo agora está disponível
            second_fork <- struct{}{} // devolve segundo garfo
            fmt.Println(strconv.Itoa(id) + " levanta e pensa")
        default: // se segundo garfo não está disponível
            fmt.Println(strconv.Itoa(id) + " solta direita")
            first_fork <- struct{}{} // devolve primeiro garfo
            <-second_fork_available // espera segundo garfo ficar disponível
        }
    }
}

func main() {
    var fork_channels [FORKS]chan struct{}
    for i := 0; i < FORKS; i++ {
        fork_channels[i] = make(chan struct{}, 1)
        fork_channels[i] <- struct{}{} // no inicio garfo esta livre
    }

    var second_fork_available_channels [PHILOSOPHERS]chan bool
    for i := 0; i < PHILOSOPHERS; i++ {
        second_fork_available_channels[i] = make(chan bool, 1)
        second_fork_available_channels[i] <- true // no inicio segundo garfo está disponível
    }

    for i := 0; i < PHILOSOPHERS; i++ {
        fmt.Println("Filosofo " + strconv.Itoa(i))
        go philosopher(i, fork_channels[i], fork_channels[(i+1)%PHILOSOPHERS], second_fork_available_channels[i])
    }

    var blq chan struct{} = make(chan struct{})
    <-blq
}
