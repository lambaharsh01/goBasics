package main

import "fmt"

type MinHeap struct {
	Arr []int
}

func (h *MinHeap) Insert(x int) {
	h.Arr = append(h.Arr, x)
	
	var i = len(h.Arr) - 1

	for i > 0 {
		var mid = (i - 1) / 2 
		if h.Arr[mid] >= h.Arr[i] {
			break
		}
		h.Arr[i], h.Arr[mid] = h.Arr[mid], h.Arr[i]
		i = mid
	}
}


func (h *MinHeap) Extract() {
	n := len(h.Arr)
	if n == 0 {
		return
	}

	h.Arr[0], h.Arr[n-1] = h.Arr[n-1], h.Arr[0]
	h.Arr = h.Arr[:n-1]

	n--
	var i int

	for {

		left := (i * 2) + 1
		right := (i * 2) + 2

		smallest := i

		if left  < n && h.Arr[left] > h.Arr[smallest] {
			smallest = left
		}
		if right  < n && h.Arr[right] > h.Arr[smallest] {
			smallest = right
		}

		if smallest == i {
			break
		}

		h.Arr[i], h.Arr[smallest] = h.Arr[smallest], h.Arr[i] 
		i = smallest
	}
}

func RunHeap() {

	// MIN HEAP = Parent > Child 
	// MAX HEAP = Parent < Child

	var mh = MinHeap{}
	mh.Insert(10)
	mh.Insert(1)
	mh.Insert(3)
	mh.Insert(6)
	fmt.Println(mh.Arr)
	mh.Insert(2)
	fmt.Println(mh.Arr)
	mh.Insert(6)
	fmt.Println(mh.Arr)
	mh.Extract()
	fmt.Println(mh.Arr)
	mh.Extract()
	fmt.Println(mh.Arr)
	mh.Extract()
	fmt.Println(mh.Arr)
	// mh.Extract()
	// fmt.Println(mh.Arr)
}
