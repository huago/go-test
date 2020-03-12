package test

func SomeRand() []int {
	//var sr = []int{1, 5, 3, 7, 2, 9, 4, 6, 8}
	//rand.Seed(time.Now().Unix())

	//for i := 0; i < 10; i++ {
	//
	//	num := rand.Int() % 100
	//	sr = append(sr, num)
	//}

	sr := []int{7, 5, 2, 9, 1, 3}
	sr = []int{2, 1, 7, 8, 9, 3, 4, 5, 6}

	return sr
}

func BubbleSort(bs []int) []int {
	for i := 0; i < len(bs)-1; i++ {
		for j := i + 1; j < len(bs); j++ {
			if bs[i] > bs[j] {
				tmp := bs[j]
				bs[j] = bs[i]
				bs[i] = tmp
			}
		}
	}

	return bs
}

// 无效算法
func BubbleSortWithFlag(bs []int) []int {
	var flag = true

	for i := 0; i < len(bs)-1 && flag; i++ {
		flag = false
		for j := i + 1; j < len(bs); j++ {
			if bs[i] > bs[j] {
				tmp := bs[j]
				bs[j] = bs[i]
				bs[i] = tmp

				flag = true
			}
		}
	}

	return bs
}

func InsertSort(is []int) []int {
	// 从第二个元素开始遍历，与前面的子序列作对比
	for i := 1; i < len(is); i++ {
		// 当前节点值
		tmp := is[i]
		var j int
		// 元素i之前的子序列
		for j = i - 1; j >= 0; j-- {
			if is[j] > tmp {
				// j元素值后移
				is[j+1] = is[j]
			} else {
				break
			}
		}

		// 找到合适的位置，插入进去
		is[j+1] = tmp
	}

	return is
}

func QuitSort(qs *[]int, left int, right int) []int {
	if left < right {
		k := partition(qs, left, right)

		QuitSort(qs, left, k-1)
		QuitSort(qs, k+1, right)
	}

	return nil
}

func partition(qs *[]int, left int, right int) int {
	// 以最后一个元素为基准点
	arr := *qs
	key := arr[right]

	i := left - 1
	var tmp int

	// 开始以基准点为标准分割序列
	for j := left; j < right; j++ {
		if arr[j] < key {
			i++

			tmp = arr[j]
			arr[j] = arr[i]
			arr[i] = tmp
		}
	}

	return 0
}
