package main

import "fmt"

func main() {

	bookworms, err := loadBookworms("testdata/bookworms.json")
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(bookworms)

	for i, bookwrm := range bookworms {
		fmt.Println(i, bookwrm)
	}

}
