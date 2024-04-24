package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var aStack []int
	var bStack []int
	list := ""
	// input := ""     // Pour Checker
	err := false

	if len(os.Args) > 1 {
		list = os.Args[1]
	}

	aStack, err = InitList(list)

	for i := 1; i < len(aStack)-1; i++ {
		if aStack[i] == aStack[i-1] {
			err = true
			// fmt.Println("doublon trouvé") ***Debug***
		}
	}
	if err {
		fmt.Println("Error: Duplicate found.")
		return
	}

	fmt.Println("Input a: ", aStack) //***Debug***
	fmt.Println("Input b: ", bStack) //***Debug***

	// ******** PUSH-SWAP *********

	// Start sorting :
	for len(aStack) != 3 {
		// Send every element from aStack to bStack (from smallest to Biggest)

		// 1. Find the smallest number in a and put it on top
		aStack = MoveBiggestTop(aStack)

		// 2. Push the smallest number on top of aStack to bStack
		aStack, bStack = TransferElement(aStack, bStack, "pb")

	}
	//aStack = SwapFirstNbs(RotateStack(aStack))
	// From here, the list is sorted in bStack
	/*for len(bStack) > 0 {
		// Now we push eveything back to aStack
		bStack, aStack = TransferElement(bStack, aStack, "pa")
	}*/

	// ******** CHECKER *********

	// To Do :

	// 1: Get the entire line passed as input (scanln, scanf, scan ?)

	// 2: Apply the commands to the stack

	// 3: If the stack is well sorted by the commands, return OK, else return KO

	fmt.Println("Output a: ", aStack)
	fmt.Println("Output b: ", bStack)
}

func InitList(list string) (stack []int, pb bool) {
	// On initialise les tableaux

	for i := range list {
		if list[i] != ' ' {
			// On transforme les elements en integers
			tmp, err := strconv.Atoi(string(list[i]))
			if err != nil {
				pb = true
			}

			stack = append(stack, tmp)
		}
	}
	return stack, pb
}

func SwapFirstNbs(stack []int) (sortedStack []int) {
	// Fonction pour les commandes SA et SB
	sortedStack = append(sortedStack, stack[1])
	sortedStack = append(sortedStack, stack[0])

	for i := 2; i < len(stack); i++ {
		sortedStack = append(sortedStack, stack[i])
	}
	return sortedStack
}

func RotateStack(stack []int) (sortedStack []int) {
	// Fonction pour les commandes RA et RB

	for i := 0; i < len(stack)-1; i++ {
		sortedStack = append(sortedStack, stack[i+1])
	}
	sortedStack = append(sortedStack, stack[0])

	fmt.Println("ra")

	return sortedStack
}

func ReverseRotateStack(stack []int) (sortedStack []int) {
	// Fonction pour les commandes RRA et RRB
	sortedStack = append(sortedStack, stack[len(stack)-1])

	for i := 1; i < len(stack); i++ {
		sortedStack = append(sortedStack, stack[i-1])
	}

	return sortedStack
}

func RotateBothStacks(aStack []int, bStack []int) {
	RotateStack(aStack)
	RotateStack(bStack)
	fmt.Println("rr")
}

func ReverseRotateBothStacks(aStack []int, bStack []int) {
	ReverseRotateStack(aStack)
	ReverseRotateStack(bStack)
	fmt.Println("rrr")
}

func TransferElement(I_origin []int, I_destination []int, message string) (O_origin []int, O_destination []int) {
	// Fonction pour les commandes PA et PB
	O_destination = I_destination

	O_destination = append(O_destination, I_origin[0])

	for i := 1; i < len(I_origin); i++ {
		O_origin = append(O_origin, I_origin[i])
	}

	fmt.Println(message)

	return O_origin, O_destination
}

func MoveBiggestTop(stack []int) (newStack []int) {
	// Fonction qui ramène chaque plus grand élement du tableau tout en haut
	biggest := stack[0]

	for i := range stack {
		if stack[i] > biggest {
			biggest = stack[i]
		}
	}
	for stack[0] < biggest {
		stack = RotateStack(stack)
	}
	newStack = stack
	return newStack
}
