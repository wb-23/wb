package main

import "fmt"

func fi(){
	var input string
	var a int
	var b byte
	var C1[]byte
	var C2[]byte
	fmt.Scanln(&input)
	C1 = []byte(input)
	a =len(C1)
	for i:=a;i>0;i--{
		b=C1[i-1]
		C2= append(C2, b)
	}
	for x := range C2 {
		fmt.Printf(string(C2[x]))
	}

}
func main() {
fi()
}