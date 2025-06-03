SLICES

> A slice in Go is a descriptor (small struct) with three fields

type sliceHeader struct {
    Data uintptr // pointer to the underlying array
    Len  int     // current length
    Cap  int     // capacity (max size before reallocation)
}

> Slice is a small descriptor (a small struct) which has 3 fields:
    >> Data: pointer to the underlying array which holds the actual values
    >> Len: the actual number of elements the slice holds (increases when you append)
    >> Cap: the total size of the underlying array (e.g., if [8]int then cap is 8)

> Append & Capacity Growth:
    >> When append reaches the current capacity (e.g., 8), and you add one more element:
    >> A new bigger array is created (usually double size, like [16]int)
    >> All old values from [8]int are copied to [16]int
    >> The old smaller array is discarded
    >> The new array starts half full (len = 9, cap = 16) and stays the same until you append more

> The growth pattern is usually doubling for small slices but may vary slightly for larger ones.

> Len vs Cap:
    >> Cap = Cap is the maximum number of elements the slice can grow to without allocating a new array. It's usually (but not always) the length of the underlying array from the slice's starting point.
    >> Len = actual number of elements currently in use inside the array

> When you loop through a slice, Go uses len to know how many elements to access, ignoring the rest (even if those unused spots have default values like 0)

> How Len is tracked:
    >> len is changed (mutated) internally every time you:
    >> create a slice
    >> append to a slice
    >> reslice an existing slice
    >> len is not dynamic like in JS or Python arrays
    >> It is stored as a value in the slice struct and updated directly (no scanning of the array to count elements)

> Summary:
    >> You cannot access the length of the underlying array directly, only via cap
    >> The underlying array might be larger and partially unused (default zero values for int)
    >> len tracks how many elements are actually “visible” and used in the slice


> Slice acts like a window into the array — len is how wide your window is, cap is the total width of the array



