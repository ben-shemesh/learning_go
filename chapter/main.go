package main

import (
	"fmt"
)
func averag(xs[]float64) float64 {
	total2 := 0.0;
	for _, v := range xs {
		total2 += v
	}
	return total2
}
func timesTwo() func() uint{
	i := uint(0)
	return func() (picked uint) {
		picked = i
		i += 2
		return
	}
}
func factorial(x int) int {
	if x == 0 {
		return 1
	}
	fmt.Println(x)
	return factorial(x -1)
}
func main(){
	fmt.Println(factorial(12))
	// theFunc := timesTwo()
	// fmt.Println(theFunc())
	// fmt.Println(theFunc())

	// theSumUp := func(nums ...int) int {
	// 	sums := 0;
	// 	for _, theOnes := range nums {
	// 		sums += theOnes
	// 	}
	// 	return sums;
	// }
	// myNewNum := 0
	// incrementIt := func() int {
	// 	myNewNum ++
	// 	fmt.Println("the number is", myNewNum)
	// 	return myNewNum
	// }
	// s2 := []float64{23,45,35,232,54,43,}
	// totalReturn := averag(s2)
	// fmt.Println(totalReturn)
	// voradicFun := []int{23,34,10}
	// myInts := []int{4,3,5}
	// fmt.Println(theSumUp(myInts...))
	// fmt.Println(incrementIt())
	// fmt.Println(incrementIt())
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
	// for i := 0; i <= 10; i++ {
	// 	switch i {
	// 	case 0: fmt.Println("Zero")
	// 	case 1: fmt.Println("One")
	// 	case 2: fmt.Println("Two")
	// 	case 3: fmt.Println("Three")
	// 	default: fmt.Println("Im not printing the rest..")
	// 	}
	// 	var i = 1
	// 	for i < 100 {
	// 		if i % 3 == 0 && i % 5 == 0 {fmt.Println(i,":FizzBuzz")
	// 		} else if i % 3 == 0 {fmt.Println(i,":Fizz")
	// 		} else if i % 5 == 0 {fmt.Println(i,":Buzz")
	// 		} else {fmt.Println(i)}
	// 		i++
	// 	}
	// }
		//Array, Slices and Maps
		// var x[5]int;
		// x[4] = 100;
		// fmt.Println(x[4])

		// scores :=[5]float32{34,454,56,234,54}
		// var total float32 = 0
		// for _, value := range scores {
		// 	fmt.Println(value)
		// 	total += value
		// }
		// fmt.Println(total / float32(len(scores)))
		//  newScores := make([]float64,5,10)
		// fmt.Scanf("%f", &newScores)
		// fmt.Println(newScores[0])
		// arr := []int{2,4,56,7,7,}
		// fmt.Println(arr)
		// y := append(arr,3,4,5)
		// fmt.Println(y)
		// 	slice1 := []int{1,2,3,}
		// 	slice2 := make([]int, 2)
		// 	slice3 := copy(slice2,slice1)
		// 	fmt.Println(slice1,slice2)
		// 	fmt.Println(slice3)
		// x := make(map[int]int)
		// x[5] = 10;
		// fmt.Println(x)
		// maps are a way to connect a value with a key
			// the key can be of one type and the value can be of another or same type
		// elements := make(map[string]string)
		// elements["Hi"] = "Hydrogen";
		// elements["He"] = "Helium"
		// fmt.Println(elements["He"])
		// name, ok := elements["Sh"];
		// fmt.Println(name,ok)
			// elements := map[string]map[string]string{
			// 	"Hi": map[string]string{
			// 		"name": "Hydrogen",
			// 		"state": "gas",
			// 	},
			// 	"He": map[string]string{
			// 		"name": "Helium",
			// 		"state": "gas",
			// 	}, "Li": map[string]string{
			// 		"name": "Litihum",
			// 		"state": "gas",
			// 	},
			// }
			// if ele, ok := elements["He"]; ok {
			// 	fmt.Println(ele["name"], ele["state"])
			// }
			// x := []int{ 48,96,86,68,
			// 	57,82,63,70,
			// 	37,34,83,3,27,
			// 	19,97,9,17,4,
			// }
			// // // finiding the smallest number
			// for i := 0; i < len(x); i++ {
			// 	smallest := x[i]
			// 	for j:= 0; j < len(x); j++{
			// 		if x[j] < x[i] {smallest = x[j]} 
			// 		fmt.Println("the smallest number is", smallest)
			// 	}
			// }

	// xs := []float64{23,34,54,98,75,78,88,}
	// total2 := 0.0;
	// for _, v := range xs {
	// 	total2 += v
	// }
	// fmt.Println(total2)
	// dividedBy := float64(len(xs))
	// fmt.Println(dividedBy)
	// fmt.Println(total2 / dividedBy)
}
