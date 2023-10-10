package main

import "fmt"

func main() {
	s := make([]string, 0)
	s = append(s, "a", "b", "c")

	for key, _ := range s {
		s[key] = "new" + string(key)
		fmt.Println(s[key])
	}

	fmt.Println(s)

}
