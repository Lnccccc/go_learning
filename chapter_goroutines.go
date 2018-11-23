package main

import "fmt"

func f(from string) {
	for i := 0;i<3;i++ {
		fmt.Println(from,":",i)
	}
}

func main () {

	go func() {
		for i := 0;i<100;i++ {
			fmt.Println("go1:",i)
		}
	}()
	go func() {
		for j :=0;j<100;j++ {
			fmt.Println("go2:",j)
		}
	}()
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}



