package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	nums := []int{}

	for _, i := range numbersToSum {
		if len(numbersToSum) == 0 {
			nums = append(nums, 0)
		} else {
			nums = append(nums, Sum(i))

		}
	}

	return nums
	//panic("unimplemented") )))))
}

/*
for i := 0; i < 5; i++ {
		sum += numbers[i]
	}

размер массива относится к его типу)
Применение указателя на массив оказывается эффективным и позволяет вызываемой функции изменять переменные вызывающей функции, но массивы остаются
негибким решением из-за присущего им фиксированного размера. Например, функция z e ro не примет указатель на переменную [1 6 ] b y te ; нет также никакого способа добавления или удаления элементов массива. По этим причинам, за исключением
особых случаев, таких как хеш SHA256 фиксированного размера, массивы в качестве
параметров функции используются редко; вместо этого обычно используются срезы
*/