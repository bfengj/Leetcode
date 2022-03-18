package question191

func hammingWeight(num uint32) int {
	res := 0
	for ;num>0;num&=num-1 {
		res++
	}
	return res
}

/*func hammingWeight(num uint32) int {
	return strings.Count(strconv.FormatInt(int64(num),2),"1")
}*/