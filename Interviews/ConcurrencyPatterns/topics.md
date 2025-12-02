# Go Concurrency Topics

A structured list of important and commonly asked concurrency concepts in Go.

---

## 1. Fundamentals
- **Alternate Printing**
- **Avoiding Data Races**
  - Understanding `Mutex` vs `RWMutex`

---

## 2. High-Probability Interview Topics
- **Worker Pool Pattern**
- **Rate Limiter** (Token Bucket / Leaky Bucket)
- **Timeout Handling** using `select` and `time.After`
- **Graceful Shutdown** with `context` and OS signals
- **Fan-in / Fan-out Pattern**

---

## 3. Essential Concurrency Patterns
- **Non-blocking Channel Operations** using `select`
- **Context Cancellation** (stopping goroutines cleanly)
- **Semaphore Pattern** (limiting concurrency via buffered channels)
- **Pipeline Pattern** (multi-stage processing with channels)

---
