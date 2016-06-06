package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("started...")
	reschan := make(chan float64)
	go calculator(1, 1, 1,reschan)
	select {
	case res := <-reschan:
		fmt.Println("result:", res)
	case <-time.After(1 * time.Second):
		fmt.Println("timed out: 1s")
	}

}

func calculator(a float64, b float64, c float64,channel chan float64){
	channel <-(a + 15/b*6 - 0.002) / ((b/2+1)*6/13 + 1)
}
