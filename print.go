package main

import ("fmt"
		"strconv"
	)

func main()  {
	
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[0]) //in byte
	fmt.Println("Hello World"[3:7])
	fmt.Println("Hello World"[3:]) //tak pedulikan yg akhir

	//conversion from string to int

	var x string = "100"

	xInt, _ := strconv.Atoi(x)
	fmt.Println(xInt)

	xString := strconv.Itoa(-100)
	fmt.Println(xString)
	fmt.Println("Jenkins Modified")
}
