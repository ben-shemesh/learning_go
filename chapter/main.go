package main

import (
	"fmt"
)

func main(){
	// var name string;
	// name = "Avin"
	// fmt.Println("Hello,", name);
	// lengthName := len(name);
	// anotherNumer := 4;
	// fmt.Println(lengthName + anotherNumer);
	// trueOrfalse := (true && false) || (false && true) || !(false && false)
	// fmt.Println(trueOrfalse);
	// var x string;
	// fmt.Println(x)
	// x = x + "second";
	// fmt.Println(x)

	// fmt.Println("enter a number ");
	// var input int;
	// // the %d represents a decimal number
	// 	// the user inpute read from the standard input
	// 	// must be directed to the address of the variable
	// 	// to display the results.
	// fmt.Scanf("%d", &input);
	// outPut := input * 2;
	// fmt.Println("The number the user selected is: ", input)
	// fmt.Println("The number selected doubled is: ", outPut)
	// // exercises
	// // 1.) what are two ways to create a variable
	// var celcius int;
	// fmt.Println("please enter a number in Celcius to convert to Fahrenheit.")
	// fmt.Scanf("%d", &celcius)
	// var	farenheit int;
	// fmt.Println("Please enter a number in Fahrenheit to convert to Celsius.")
	// fmt.Scanf("%d", &farenheit)
	// farenheitTemp := (celcius * 9/5) + 32;
	// celciusTemp   := (farenheit - 32) *5/9;
	// fmt.Printf("the conversion from Celcius %d to fahrenheit is: %d degrees. \n",celcius, farenheitTemp  ) 
	// fmt.Printf("the conversion from Celsius %d to fahrenheit is: %d degrees. \n",celcius, celciusTemp  	 ) 

	// fmt.Println("Please enter a number to that you want to convert from feet to meters.")
	// var feet float32;
	// fmt.Scanf("%f", &feet);
	// meters := feet * 0.3048;
	// fmt.Println(meters)
	for i := 0; i <= 10; i++ {
		switch i {
		case 0: fmt.Println("Zero")
		case 1: fmt.Println("One")
		case 2: fmt.Println("Two")
		case 3: fmt.Println("Three")
		default: fmt.Println("Im not printing the rest..")
		}
		var i = 1
		for i < 100 {
			if i % 3 == 0 && i % 5 == 0 {fmt.Println(i,":FizzBuzz")
			} else if i % 3 == 0 {fmt.Println(i,":Fizz")
			} else if i % 5 == 0 {fmt.Println(i,":Buzz")
			} else {fmt.Println(i)}
			i++
		}
	}
}
