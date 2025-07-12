package main

import "fmt"

type transformFn func(int)int

// fx - 1st class values in GoLang
func doubleNums(numbers *[]int)[]int{
	dNums := []int{}
	for _,val :=range *numbers{
		dNums= append(dNums, double(val))
	}
	return  dNums
}

func double(num int)int{
	return num*2
}

func quadrupled(num int)int{
	return num*4
}

func transformNums(nums *[]int, transform transformFn)[]int{
tNums:=[]int{}
for _, val := range *nums{
	tNums = append(tNums, transform(val))
}
return  tNums
}



func main(){
nums:= []int{1,2,3}
doubled:=doubleNums(&nums)
quadrupledNums:=transformNums(&nums,quadrupled)
fmt.Println(nums,"->",doubled)
fmt.Println(nums,"->",quadrupledNums)
}