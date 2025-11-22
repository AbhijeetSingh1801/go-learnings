package slices_tut

import (
	"fmt"
)


type Report struct{
	ID		int	
	Data	[5000]byte 
	Status	string
}

func printReportValue(r Report) {
	fmt.Printf(
		"Status: %s (Copy)\n", r.Status,
	)

	fmt.Println("----------")
	fmt.Println(&r.ID)
}

func printReportPrinter(rPtr *Report){
	fmt.Printf("Status: %s (Original)\n", rPtr.Status)

	fmt.Println("----------")

	fmt.Println(&rPtr)
}


func zero(xPtr *int){
	*xPtr = 0
}

func Indexfunc() {
	x:=5
	zero(&x)

	fmt.Println(x)

	hugeReport:= Report{ID: 1, Status: "Draft"}

	printReportValue(hugeReport)
	printReportPrinter(&hugeReport)

}