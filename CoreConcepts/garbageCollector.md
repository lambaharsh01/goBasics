Go's Garbage Collector

> Go uses a concurrent, non-generational GC 
    >> Concurrent: works alongside the program does not stop it in any means
    >> Non-Generational: does not divide memory into "Young" or "Old" objects(JVM and .NET uses Generational GC)
    >> Tracing: It identifies live objects by following references starting from root variables (like stack pointers, global variables, etc.).
        >>> Tracing In Layman Terms: GC check of the root is still active or being used if it does it basically means the heap is still relevant and should stay alive.  


> GC Only Collects Heap-Allocated Memory
> What it does not collects?
    >> Stack Memory - As it automatically deallocated after use/end of the function
    >> OS Level Resources - Open Files or Sockets

> GC Uses Mark And Sweep Algorithm
    >> Mark:
        >>> Marks All Root References (i.e Global Variables, Local Variables, Goroutine Stacks)
        >>> Follow From root to every point where the root is referenced in memory  
        >>> Mark What is Still needed(reachable by the program) through the chaining of every node from the root
    >> Sweep:
        >>> Anything Unmarked is considered Garbage
        >>> GC adds Unmarked Heap Blocks to a free list or memory pool
            >>>> Layman: When new heap memory is needed, the GC reuses unmarked (unused) memory blocks by overwriting them with new data.

PERFORMANCE
> More Heap Allocations ----> More Memory to Track ----> GC has to scan more objects ----> CPU cycles consumed ----> High allocation rate ----> more frequent GC run ----> more CPU usage ----> less CPU for your actual logic


GOOD PRACTICES
> Use sync.Pool for temporary objects ----> Allows object reuse between GC cycles
> Profile with pprof or -gcflags="-m" ----> Find where allocations and escapes happen
        