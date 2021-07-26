package pascal

func Triangle(n int) [][]int {

	t := [][]int{
		{1},
	}

	for i := 1; i < n; i++ {
		tri := []int{1}
		prev := t[i-1]
		for y := 1; y < i+1; y++ {
			cur := prev[y-1]
			if y <= len(prev)-1 {
				cur += prev[y]
			}
			tri = append(tri, cur)
		}
		t = append(t, tri)
	}

	return t
}
