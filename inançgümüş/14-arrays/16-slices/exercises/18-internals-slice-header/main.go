// EXERCISE: Observe the memory allocations
//
//  In this exercise, your goal is to observe the memory allocation
//  differences between arrays and slices.
//
//  You will create, assign arrays and slices then you will print
//  the memory usage of your program on each step.
//
//  Please follow the instructions inside the code.
//
//
// EXPECTED OUTPUT
//
//  Note that, your memory usage numbers may vary. However, the size of the
//  arrays and slices should be the same on your own system as well
//  (if you're on a 64-bit machine).
//
//
//  [initial memory usage]
//          > Memory Usage: 104 KB
//  [after declaring an array]
//          > Memory Usage: 78235 KB
//  [after copying the array]
//          > Memory Usage: 156365 KB
//  [inside passArray]
//          > Memory Usage: 234495 KB
//  [after slicings]
//          > Memory Usage: 234497 KB
//  [inside passSlice]
//          > Memory Usage: 234497 KB
//
//  Array's size : 80000000 bytes.
//  Array2's size: 80000000 bytes.
//  Slice1's size: 24 bytes.
//  Slice2's size: 24 bytes.
//  Slice3's size: 24 bytes.
//
//
// HINTS
//
//  I've declared a few functions to help you.
//
//    report function:
//    - Prints the memory usage.
//    - Just call it with a message that matches to the expected output.
//
//    passArray function:
//    - It automatically prints the memory usage.
//    - Accepts a [size]int array, so you can pass your array to it.
//
//    passSlice function:
//    - It automatically prints the memory usage.
//    - Accepts an int slice, so you can pass it one of your slices.
//
// ---------------------------------------------------------

package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"unsafe"
)

const size = 1e7

func main() {
	debug.SetGCPercent(-1)

	report("initial memory usage")

	var array [size]int
	report("after declaring an array")

	array2 := array
	report("after copying the array")

	passArray(array)

	slice1 := array[:]
	slice2 := array[:1e3]
	slice3 := array[1e3:1e4]
	report("after slicings")

	passSlice(slice3)

	fmt.Println()
	fmt.Printf("Array's size : %d bytes.\n", unsafe.Sizeof(array))
	fmt.Printf("Array2's size: %d bytes.\n", unsafe.Sizeof(array2))
	fmt.Printf("Slice1's size: %d bytes.\n", unsafe.Sizeof(slice1))
	fmt.Printf("Slice2's size: %d bytes.\n", unsafe.Sizeof(slice2))
	fmt.Printf("Slice3's size: %d bytes.\n", unsafe.Sizeof(slice3))
}

func passArray(items [size]int) {
	items[0] = 100
	report("inside passArray")
}

func passSlice(items []int) {
	report("inside passSlice")
}

func report(msg string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("[%s]\n", msg)
	fmt.Printf("\t> Memory Usage: %v KB\n", m.Alloc/1024)
}
