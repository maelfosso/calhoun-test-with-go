package example

import "fmt"

var out string = "this is not used"

func ExampleHello() {
	greeting, err := Hello("Jon")
	if err != nil {
		panic(err)
	}
	fmt.Println(greeting)

	// Output:
	// Hello, Jon
}

func ExamplePage() {
	checkIns := map[string]bool{
		"Bob":   true,
		"Alice": false,
		"Eve":   true,
		"Janet": true,
		"Susan": false,
		"John":  true,
	}
	Page(checkIns)

	// Unordered Output:
	// Paging Alice; please see the front desk to check in.
	// Paging Susan; please see the front desk to check in.
}
