package question567

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	n1, n2 := len(s1), len(s2)
	if n1>n2 {
		return false
	}
	myMap := make(map[byte]int)
	for i := 0; i < n1; i++ {
		myMap[s1[i]]++
		myMap[s2[i]]--
	}
	var diff int = 0
	for _,k :=range myMap{
		if k!=0 {
			diff++
		}
	}
	if diff==0 {
		return true
	}

	for i := n1; i < n2; i++ {
		if myMap[s2[i-n1]]==0 {
			diff++
		}
		myMap[s2[i-n1]]++
		if myMap[s2[i-n1]]==0 {
			diff--
		}
		if myMap[s2[i]]==0 {
			diff++
		}
		myMap[s2[i]]--
		if myMap[s2[i]]==0 {
			diff--
		}
		fmt.Println(diff)
		if diff == 0 {
			return true
		}
	}
	return false

}
/*func checkInclusion(s1 string, s2 string) bool {
	n1, n2 := len(s1), len(s2)
    if n1>n2 {
        return false
    }
	map1 := make(map[byte]int)
	map2 := make(map[byte]int)
	for i := 0; i < n1; i++ {
		map1[s1[i]]++
		map2[s2[i]]++
	}
	var fflag bool = true
	for k, v := range map1 {
		if map2[k]!=v {
			fflag = false
		}
	}
	if fflag==true {
		return true
	}
    fmt.Println(map1)
	for i := 1; i < n2-n1+1; i++ {
		map2[s2[i-1]]--
		map2[s2[i+n1-1]]++
        fmt.Println(map1)
		var flag bool = true
		for k, v := range map1 {
			if map2[k]!=v {
				flag = false
			}
		}
		if flag==true {
			return true
		}
	}
	return false

}
*/
