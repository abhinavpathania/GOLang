package main

import(
	"fmt"
	"time"
)

func measure_time(){
	var n int = 1000000000
	myslice:=[]int{}
	myslice_pre:=make([]int,0,n)

	fmt.Printf("Total time without preallocation: %v \n",timeLoop(myslice,n))
	fmt.Printf("Total time with preallocation: %v",timeLoop(myslice_pre,n))
}

func timeLoop(slice []int, n int) time.Duration{
	t0:=time.Now()
	for len(slice)<n{
		slice = append(slice,1)
	}
	return time.Since(t0)

}