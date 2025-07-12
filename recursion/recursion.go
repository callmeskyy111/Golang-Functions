
package recursion
import "fmt"

// Recursion - A fx that calls itself
// Better than loops - More efficient


func factorial(num int)int{
	if num==0{
		return 1
	}
		return num * factorial(num-1)
}

func main(){
	fmt.Println("5:",factorial(5))
	fmt.Println("9:",factorial(9))
}