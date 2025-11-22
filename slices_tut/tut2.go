package slices_tut


import (
	"fmt"
)
func tut2(){
	users:= []string{"Alices", "Bob", "charlie"}

	for index, value:= range users {
		fmt.Printf("Index: %d -> used User: %s", index, value)
		fmt.Println()
	}

	fmt.Println("---------")


	for _, value := range users {
		fmt.Println("User: ", value)
	}

	// maps 

	scores:= make(map[string] int)

	scores["Alices"] = 95 
	scores["Bob"] =  82


	fmt.Println("Alice's Score", scores[users[0]])

	score, exists := scores[users[2]]

	if exists {
		fmt.Printf("Score for %s is %d", users[1], score)
		fmt.Println()
	} else {
		fmt.Println("User does not have scores")
	}


	delete(scores, users[2])
}