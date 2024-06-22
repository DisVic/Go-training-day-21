n := 10
	numbers := make([]uint8, n)
	for i := 0; i < n; i++{
		_, _ = fmt.Scan(numbers[i])
	}
	counter := 0
	for counter < 3{
		var a, b int
		_, _ = fmt.Scan(&a, &b)
		temp := numbers[a]
		numbers[a] = numbers[b]
		numbers[b] = temp
		counter++
	}
	for k := 0; k < n; k++{
		fmt.Print(numbers[k], " ")
	}