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
var wg sync.WaitGroup //the waitgroup assures that the program is block until all declared goroutines are done
var mux sync.Mutex    //for blocking simultaneous access to the counter, used in message generation
var counter = 0

func main() {
	defer fmt.Println("fin du programme")
	//declare the channel that is used for passing msg from the generator to the printer
	chanmsg := make(chan string, buff)
	done :=make(chan bool,1)
	//declare the go routines that will be started to the waitgroup
	wg.Add(limit)
	//send msg through the pipe (chanmsg), non blocking
	for i := 0; i < limit; i++ {
		go GenerateMsg("coucou", chanmsg, &wg)
	}
	//receive from the pipe and print
	for i := 0; i < limit; i++ {
		go PrintMsg(chanmsg, &wg)
	}
	//wait untill all goroutines have been release and are done
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
	//send the message through the pipe (chanmsg)
	channel <- fmt.Sprintf("%s pour la %de fois", template, counter)
}

func PrintMsg(channel chan string, wg *sync.WaitGroup) {
	//declare to the waitgroup that the routine is done
	defer wg.Done()
	//print msg from the pipe
	fmt.Printf("message reçu:'%s'\n", <-channel)
}
