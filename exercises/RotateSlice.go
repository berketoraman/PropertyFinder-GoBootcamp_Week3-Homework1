package main

import "fmt"

func rotate(s []int) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
      s[i], s[j] = s[j], s[i]
    } 
}

func main() {
	fmt.Println("For left rotation write 1.For right rotation right 2")
	
	fmt.Printf("State your rotation side: ")
	var side int
	fmt.Scanln(&side)

	/* Rotate s right by "Rotation value" positions.First we reverse the slice and use the same 
       method from rotating left but since we reversed the slice we get a result like right rotation
    */

	if side == 1 { // if the user chooses 1 this is applied for left rotation
     s := []int{0, 1, 2, 3, 4, 5}
	 fmt.Printf("Enter Left Rotation value: ")
     var rotation int
     fmt.Scanln(&rotation) //Ask the user for the rotation value
	 rotate(s[:rotation])
	 rotate(s[rotation:])
	 rotate(s)
	 fmt.Println(s)  
	} else if side == 2 { //if the user chooses 2 this is applied for right rotation
	 s := []int{0, 1, 2, 3, 4, 5}
	 rotate(s) // the code difference of right rotation is first we reverse the slice and then apply the same methods of left rotation
	 fmt.Printf("Enter Right Rotation value: ")
     var rotation int
     fmt.Scanln(&rotation) //Ask the user for the rotation value
	 rotate(s[:rotation])
     rotate(s[rotation:])
	 fmt.Println(s)  
	} else {
        fmt.Println("Error")
    }
} 	




