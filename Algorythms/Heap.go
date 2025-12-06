package main

import "fmt"

// There are 3 approaches in heaps
// 1. Top Down/ Shift Up approach
// Insert a new element
// You put the new element at the end and then bubble it upward until the heap property is restored.

// 2. Bottom Up/ Shift Down Approach
// When you are building a heap from an existing array
// When you remove the max/min from heap

type MaxHeap struct {
	Arr []int
}

func (h *MaxHeap) Insert(x int) {
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


func (h *MaxHeap) Extract() {
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

	arr := []int{10, 1, 3, 6, 4, 8, 7, 9, 2, 5}
	Heapify(arr)

	var mh = MaxHeap{
		Arr: arr,
	}
	// mh.Insert(10)
	// mh.Insert(1)
	// mh.Insert(3)
	// mh.Insert(6)
	// mh.Insert(4)
	// mh.Insert(8)
	// mh.Insert(7)
	// mh.Insert(9)
	// mh.Insert(2)
	// mh.Insert(5)

	fmt.Println(mh.Arr)
	mh.Extract()
	fmt.Println(mh.Arr)
	mh.Extract()
	fmt.Println(mh.Arr)
	mh.Extract()
	fmt.Println(mh.Arr)
	mh.Extract()
	fmt.Println(mh.Arr)
	mh.Extract()
	fmt.Println(mh.Arr)
	mh.Extract()
	fmt.Println(mh.Arr)
	mh.Extract()
	fmt.Println(mh.Arr)
	mh.Extract()
	fmt.Println(mh.Arr)
	mh.Extract()
	fmt.Println(mh.Arr)
	// fmt.Println(mh.Arr)
	// mh.Extract()
	// fmt.Println(mh.Arr)
	// mh.Extract()
	// fmt.Println(mh.Arr)
	// mh.Extract()
	// fmt.Println(mh.Arr)


    // nums := []int{3, 2, 3, 6, 7, 2}
    // fmt.Println("Initial array:", nums)
    // fmt.Println("---- Starting Bottom-Up Heapify ----\n")
    // buildMaxHeap(nums)
    // fmt.Println("Final Max-Heap:", nums)
}


// SORTING 

// INTUITION:
//
// If I try to sort an array by first applying top-down heap construction 
// (inserting elements one by one → sift-up), and then doing bottom-up heap work,
// the time becomes unnecessary O(n log n) + O(n log n).
//
// However, if the array is already available, I can directly apply the 
// bottom-up heapify approach to convert it into a valid heap in O(n) time.
//
// This works because during bottom-up heapify most nodes are leaves or near 
// the bottom, so they move very little (cheap iterations), while only a few 
// nodes near the top may move more (expensive iterations). Since the majority 
// of elements are near the bottom, the total average (and actual worst-case)
// time to build the heap is O(n).
//
// Important Clarification:
// Bottom-up heapify is O(n) *only for building the heap from an array*.
// But when we start extracting elements (heap sort), each extraction performs 
// a sift-down from the root, which costs O(log n). Doing this n times gives 
// heap sort a total time of O(n log n).
//
// Summary:
// - Build heap (bottom-up):     O(n)
// - Each extract (sift-down):   O(log n)
// - Full heap sort:             O(n log n)

// Two types of heapify:
// Bottom-up heapify: Start from the last non-leaf node and move upwards. Used in building a heap from an unsorted array.
// Top-down heapify: Start from a node and push it downwards. Used after removing the root of a heap.


func Heapify(arr []int) {
	// EXAMPLE IS TO HEAPIFY THE ARRAY FOR A MIN HEAP
	// CAN BE MODIFIED EASILY ACCORDING TO MAX HEAP
	for i:=(len(arr) / 2) -1; i>=0; i-- {
		ShiftDown(arr, i, len(arr))	
	}
}


func ShiftDown(arr []int, i, n int) {
	for {
		left := (i*2) + 1
		right := (i*2) + 2

		highest := i 

		if left < n && arr[left] > arr[highest] {
			highest = left
		}
		if right < n && arr[right] > arr[highest] {
			highest = right
		}

		if highest == i {
			return
		} 

		arr[highest], arr[i] = arr[i], arr[highest]
		i = highest
	}
}
