package main

import "fmt"

func main(){
	var history [] int
	var Numb1,Numb2,result int
	var a string
	for true {
		fmt.Println("请输入算式")
		fmt.Scanln(&Numb1, &a, &Numb2)
		switch a {
		case "+":result=Numb1+Numb2
		case "-":result=Numb1-Numb2
		case "*":result=Numb1*Numb2
		case "/":result=Numb1/Numb2
		default:
			fmt.Println("运算符不正确")
		}
		history = append(history, result)
		fmt.Println("结果是:",result)
		for _,number:= range history {
			fmt.Print(number," ")
		}
		fmt.Println()
	}
}

