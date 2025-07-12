package variadic

import "fmt"

func main(){
	numbers:= []int{1,2,2,2}
	fmt.Println(sumUp(numbers)) //7

	// variadic fx - Variable no. of params.
	fmt.Println(sumUpVariadic(1,2,3,4,5,6)) //21
	fmt.Println(sumUpVariadic(2,55,66,7,8,1,1,1,25)) //166
}

// 1st scenario..
func sumUp(nums []int)int{
	sum := 0

	for _,val := range nums{
		sum += val
	}

	return sum
}

// 2nd scenario.. Variadic fx - any amount of params.
func sumUpVariadic(nums ...int)int{
	sum := 0

	for _,val := range nums{
		sum += val
	}

	return sum
}
