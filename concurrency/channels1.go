package main

import (
	"fmt"
	"sync"
	"time"
)

//buffering
const limit = 10
const buff = 0 //0=unbuffered -> simplifies synchronization as the channel will be waiting for a sender and receiver to be ready

//synchronisation
var wg sync.WaitGroup //the waitgroup assures that the program is blocked until all declared goroutines are done
var mux sync.Mutex    //for blocking simultaneous access to a variable (here the counter, used in message generation)
var counter = 0 //initial value

func main() {
	defer fmt.Println("fin du programme")
	//declare the channel that is used for passing msg from the generator to the printer
	chanmsg := make(chan string, buff)
	//declare the channel that alerts chanmsg when all msg have been sent
	done :=make(chan bool,1)
	//declare the go routines that will be started to the waitgroup. If know by advance, it is
	//better to set it early and avoid problems if wg.Add() arrives later that wg.Done() and the waitgroup falls to 0 or -1.
	wg.Add(limit)
	//send msg through the pipe (chanmsg), non blocking because of "go" instruction for goroutine launch
	for i := 0; i < limit; i++ {
		go GenerateMsg("coucou", chanmsg, &wg)
	}
	//receive from the pipe and print
	for i := 0; i < limit; i++ {
		go PrintMsg(chanmsg, &wg)
	}
	//wait until all goroutines have been released and are done with their task
	wg.Wait()
	done <- true
	time.Sleep(1*time.Second)
	<- done
	//go on when all done

}

func GenerateMsg(template string, channel chan string, wg *sync.WaitGroup) {
	//block counter from read/write by other goroutines
	mux.Lock()
	counter++
	//release counter for read/write
	mux.Unlock()
	//send the formatted message through the pipe (chanmsg)
	channel <- fmt.Sprintf("%s pour la %de fois", template, counter)
}

func PrintMsg(channel chan string, wg *sync.WaitGroup) {
	//declare to the waitgroup that the routine is done
	defer wg.Done()
	//read and print msg from the pipe
	fmt.Printf("message reçu:'%s'\n", <-channel)
}
