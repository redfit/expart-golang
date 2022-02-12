package main

import (
	"fmt"
)

func main() {
	// len と cap でsliceの長座、容量を取得する
	src := []int{1, 2, 3, 4}
	fmt.Println(src, len(src), cap(src))

	src = append(src, 5)
	fmt.Println(src, len(src), cap(src))

	// make で長さと容量を指定
	sliceMake := make([]int, 2, 3)
	fmt.Println(sliceMake, len(sliceMake), cap(sliceMake))
	// インデックスと値を指定する
	sliceIndex := []int{2: 1, 5: 5, 7: 13}
	fmt.Println(sliceIndex, len(sliceIndex), cap(sliceIndex))

	// sliceのコピー
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Println(dst, len(dst), cap(dst))

	// sliceの結合
	src1, src2 := []int{1, 2, 3}, []int{4, 5}
	dst = append(src1, src2...)
	fmt.Println(dst, len(dst), cap(dst))

	// 3番目の要素を削除
	dst = append(src[:2], src[3:]...)
	fmt.Println(dst, len(dst), cap(dst))

	// appendの代わりにcopyで要素削除
	i := 2
	src = []int{1, 2, 3, 4, 5}
	dst = src[:i+copy(src[i:], src[i+1:])]
	fmt.Println(dst, len(dst), cap(dst))

	// sliceの並び替え
	src = []int{1, 2, 3, 4, 5}
	for i := len(src)/2 - 1; i >= 0; i-- {
		opp := len(src) - 1 - i
		src[i], src[opp] = src[opp], src[i]
	}
	fmt.Println(src)
	for left, right := 0, len(src)-1; left < right; left, right = left+1, right-1 {
		src[left], src[right] = src[right], src[left]
	}
	fmt.Println(src)

	// 偶数のみでフィルタリング
	src = []int{1, 2, 3, 4, 5}
	dst = src[:0]
	for _, v := range src {
		if v%2 == 0 {
			dst = append(dst, v)
		}
	}
	fmt.Println(dst)

	// これでsrcをGCに回収させることが出来る
	for i := len(dst); i < len(src); i++ {
		src[i] = 0
	}

	// 任意の要素数に分割する
	src = []int{1, 2, 3, 4, 5}
	size := 2
	dst1 := make([][]int, 0, (len(src)+size-1)/size)
	for size < len(src) {
		fmt.Println("src: ", src)
		fmt.Println("size: ", size)
		fmt.Println("test: ", src[0:size])
		src, dst1 = src[size:], append(dst1, src[0:size])
	}
	dst1 = append(dst1, src)
	fmt.Println(dst1)
}
