//package main demonstrates simple solution to the given problem
package main

import (
	"fmt"
)

//func indexOf takes two strings and indicates
//whether the substring is the prefix of the given string
//substring has smaller length than string
func indexOf(sub string, s string) int {
	var i int
	for i = 0; i < len(sub); i++ {
		if sub[i] != s[i] {
			return -1
		}
	}

	return i
}

//func autocomplete finds all words, which contain given string from the given set
//using auxiliary function indexOf
//[OPTIONAL] in case the set is really big, there may be no need to give all of the results,
//but the first N (10) for the faster response as it seems to be implemented in Google search
func autocomplete(s string, set map[string]bool) []string {
	L := len(set)
	//if the string or set is empty => there is nothing to complete (from)
	if len(s) == 0 || L == 0 {
		return nil
	}

	//resulting slice
	res := make([]string, 0)
	//to count autocomlete options
	count := 1

	for v := range set {
		//if the substring is exactly the word
		if v == s {
			//add to resulting slice
			res = append(res, v)
			count++
			continue
		}
		//if substring is less than a word
		if len(v) > len(s) {
			//check if word's prefix and substring are the same
			if indexOf(s, v) > 0 {
				//and append to the slice
				res = append(res, v)
				count++
			}
		}

		//[OPTIONAL] stop the loop as mentioned above
		if count > 10 {
			break
		}
	}

	return res
}

func main() {
	//as there's no set structure in golang
	//let's create a map
	set := make(map[string]bool)
	set["dog"] = true
	set["deal"] = true
	set["deer"] = true
	set["deep"] = true

	fmt.Println(autocomplete("de", set))
}
