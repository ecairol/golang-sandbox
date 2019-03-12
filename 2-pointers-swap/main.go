package main

import (
	"fmt" 
	"bufio" 
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter 2 values:")
	
	// The blank identifier _ is used whenever syntax requires a variable name but program logic does not
	var v1, _ = reader.ReadString('\n') 
	var v2, _ = reader.ReadString('\n')

	// Trim the line break from ReadString()
	v1 = strings.TrimSuffix(v1, "\n")
	v2 = strings.TrimSuffix(v2, "\n")

	fmt.Printf("Value 1 is %s, Value 2 is %s \n", v1, v2)
	
	swap(&v1, &v2)

	fmt.Println("...swapping")
	fmt.Printf("Value 1 is now %s, Value 2 is %s \n", v1, v2)
}

// Swap pointer *v1 so it points to the address of v2 and viceversa
func swap(v1 *string, v2 *string) {
	pointerv1 := *v1
	*v1 = *v2
	*v2 = pointerv1
}
