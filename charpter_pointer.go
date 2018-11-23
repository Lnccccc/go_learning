package main

import "fmt"

func zeroval(ival int) { //变量
	ival = 0
}

func zeroptr(iptr *int) { //指针
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:",i)

	zeroval(i)
	fmt.Println("zeroval:",i)

	zeroptr(&i)
	fmt.Println("zeroptr:",i)

	fmt.Println("pointer:",&i)

}
