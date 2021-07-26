package main

import (
"encoding/json"
"fmt"
)

//JSON marshalling

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"MoneyPenny","Age":28}]`
	bs := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nall of the data:", people)

	for i, v := range people {
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age)
	}

}
