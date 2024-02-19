package cal



func MaxDiff(n []int)int  {
	if len(n)<2{
		return 0
	}

	min ,max := n[0],n[0]

	for i := 0; i < len(n); i++{
		if n[i] < min{
			min = n[i]
		}
		if(n[i] > max) {
			max = n[i]
		}
	}
	return max - min
}