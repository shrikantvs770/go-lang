# go-lang
https://www.notion.so/tactical-advantage/Go-33130115a66280dc94b7ccb214420b7c?source=copy_link

`go mod init example.com/go-basics`

```shell
@shrikantvs770 ➜ /workspaces/go-lang (main) $ go run ./01_setup_first_program/
Hello World!
@shrikantvs770 ➜ /workspaces/go-lang (main) $ go run ./01_setup_first_program/main.go 
Hello World!

```

Q. What is the entry point in Go program? how does both of the above works? significance of folder str.

Q. Is it necessary to have main package?

<br/>

Q. What's up with this?
//multiple-value os.Hostname() (value of type (name string, err error)) in single-value contextcompilerTooManyValues
// fmt.Println("os version", os.Hostname())

<br/>

Q. Declaring multiple variables at once, var vs shorthand way := in Go and any limitations
:= is only used for declaring a new variable, not assigning to an existing one.
= is only used for assigning to a EXISTING variable brother.
Hovering over shorthand ones in VS code gives you hint
<br/>

Q. Are string mutable or immutable in Go
Q. Println vs Printf
Q. Go if and else else should some where if } ends
Q. Does Go insert ; ?
Q. Arrays vs Slices in Go, arrays are fixed size, but slices are dynamic size

<details open>
  <summary>Use of capacity in slice in Go</summary>
  When you write:

```go
make([]int, 0, 5)
```

you’re creating a **slice**, not an array. It has:

* **length = 0**
* **capacity = 5**

### What does that mean?

A slice in Go has two key properties:

* **len** → how many elements are currently in the slice
* **cap** → how many elements it can hold *before needing to reallocate memory*

So here:

```go
s := make([]int, 0, 5)
```

* `len(s)` → `0` (no elements yet)
* `cap(s)` → `5` (space already allocated for 5 elements)

### Why is this useful?

It preallocates memory so that when you `append`, Go doesn’t have to keep reallocating.

Example:

```go
s := make([]int, 0, 5)

s = append(s, 10)
s = append(s, 20)
```

Now:

* `len(s)` → 2
* `cap(s)` → still 5

No reallocation happened yet.

### When does reallocation happen?

If you exceed capacity:

```go
s = append(s, 30, 40, 50, 60)
```

Now you’ve gone beyond capacity (5), so Go:

* allocates a new underlying array (doubles it brother)
* copies old elements

### Key intuition

Think of it like:

> “Give me an empty list, but reserve space for 5 items upfront.”

### Compare with other forms

```go
make([]int, 5)      // len=5, cap=5 (already filled with zero values)
make([]int, 0, 5)   // len=0, cap=5 (empty but ready to grow)
```

---

</details>

<br>
Q. Use of _ in Go everywhere


<details open>
  <summary>Multiple return values in Go</summary>
  Absolutely! 😎 Go **loves returning multiple values**, and it’s one of its most idiomatic features. I’ll show the **cleanest and best approaches** for beginners.

---

## 1️⃣ Basic multiple return

You can return multiple values directly from a function:

```go id="6vm7tx"
package main

import "fmt"

func divide(a, b int) (int, int) {
    quotient := a / b
    remainder := a % b
    return quotient, remainder
}

func main() {
    q, r := divide(17, 5)
    fmt.Printf("Quotient: %d, Remainder: %d\n", q, r)
}
```

✅ Output:

```
Quotient: 3, Remainder: 2
```

**Explanation:**

* `(int, int)` → function returns **two integers**
* `q, r := divide(...)` → receive multiple values

---

## 2️⃣ Ignoring a return value

If you only care about one value, use `_` (blank identifier):

```go
q, _ := divide(17, 5) // ignore remainder
fmt.Println("Quotient:", q)
```

---

## 3️⃣ Named return values (optional but readable)

You can **name the return values** in the function signature:

```go
func divide2(a, b int) (quotient int, remainder int) {
    quotient = a / b
    remainder = a % b
    return // can return without specifying values explicitly
}
```

* Makes code more **self-documenting**
* Useful for longer functions

---

## 4️⃣ Returning error along with a value (idiomatic Go)

Go often uses **multiple returns for errors**:

```go
func safeDivide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    result, err := safeDivide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result)
}
```

* Very idiomatic in Go
* `value, error` pattern is everywhere in the standard library

---

### ✅ Best Practices

1. Use multiple return values to **return value + error** or **related results**.
2. Use `_` if you only need one value.
3. Name returns if it improves readability for longer functions.
4. Keep it simple — don’t return **too many values**; if >3, consider a struct.

---
</details>