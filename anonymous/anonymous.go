package anonymous

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	double:= createTransformers(2)
	triple:= createTransformers(3)

	transformed := transformNumbers(&numbers, func (num int)int{return num*2})

	doubled:=transformNumbers(&numbers, double)
	tripled:= transformNumbers(&numbers, triple)

	fmt.Println(transformed)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

// Closures
func createTransformers(factor int) func(int)int{
	return  func(number int)int{
		return  number * factor
	}
}