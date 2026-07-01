# Value vs Pointer Receivers in Go

In Go, when you attach a method to a struct, you choose between a **value receiver** and a **pointer receiver**. This choice determines whether your method can modify the original struct or just a copy of it.

## The Problem: Value Receiver

```go
type Counter struct {
    value int
}

func (c Counter) increment() {
    c.value++ // modifies a COPY — original unchanged!
}

func main() {
    counter := Counter{value: 0}
    counter.increment()
    fmt.Println(counter.value) // still 0
}
```

`increment()` receives a **copy** of `counter`. The `c.value++` happens on the copy, and the original is untouched.

## The Fix: Pointer Receiver

```go
func (c *Counter) increment() {
    c.value++ // modifies the ORIGINAL
}

func main() {
    counter := Counter{value: 0}
    counter.increment()
    fmt.Println(counter.value) // 1
}
```

Now `c` is a pointer to the original `Counter`. Changes persist.

## The Auto-Address Magic

Notice we still call `counter.increment()` — not `(&counter).increment()`. Go automatically takes the address for pointer receiver methods. This is **syntactic sugar** that only works for receiver functions.

## Regular Functions Don't Get This

```go
func increment(c *Counter) {
    c.value++
}

increment(&counter) // YOU must pass & explicitly
```

For standalone functions, there's no auto-addressing. Same goes for primitives:

```go
func double(x *int) {
    *x = *x * 2
}

num := 10
double(&num) // must use &
```

## When to Use Which?

| Use | When |
|-----|------|
| Value receiver `(c Counter)` | Method only reads data, struct is small |
| Pointer receiver `(c *Counter)` | Method mutates state, or struct is large |

## One Rule to Remember

> The auto-addressing magic is **only** for the receiver (the thing before the dot), never for regular function arguments.

```go
counter.increment()    // pointer receiver → auto &  ✓
increment(&counter)    // regular function → manual & ✓
increment(counter)     // regular function → missing & ✗
```

That's it. Value receivers get copies, pointer receivers get access to the original, and Go only auto-addresses for method receivers.
