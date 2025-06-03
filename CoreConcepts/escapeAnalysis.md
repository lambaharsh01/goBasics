MEMORY ALLOCATION & ESCAPE ANALYSIS

MEMORY ALLOCATION:
> Reserving space in memory (RAM) to store variables, objects, or data structures. 

> Go uses two main memory areas:

1. Stack: 
    >> Fast
    >> Short lived memory
    >> useful for functional calls and local variables
 
2. Heap: 
    >> Slower
    >> Long lived memory 
    >> Used when variables need to outlive the functional call or are referenced elsewhere 


ESCAPE ANALYSIS:
> A process go uses at compile time 
> Analyze variables and decides:
    >> will this variable safely live in stack ? true/false
    >> will this variable escape current functional scope (e.g. via goroutine, closures and pointer return) ? true/false

if "can live on stack" = true && "does it escape or outlives functional call" = false {
    stack memory allocation will be used for the variable
} else {
    heap memory allocation will be used for the variable
}

OR 

if canLiveOnStack == true && doesEscapeFunction == false {
    useStackMemory()
} else {
    useHeapMemory()
}



