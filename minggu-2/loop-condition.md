# Looping dan Conditional

## Looping

1. Three component loop

```go
sum := 0
for i := 1; i < 5; i++ {
    sum += i
}
fmt.Println(sum) // 10 (1+2+3+4)
```

2. While loop

```go
n := 1
for n < 5 {
    n *= 2
}
fmt.Println(n) // 8 (1*2*2*2)
```

3. Infinite loop (while true)

```go
for {
    fmt.Println("Hello")
}
```

4. For-each range loop

```go
strings := []string{"hello", "world"}
for i, s := range strings {
    fmt.Println(i, s)
}
```

5. Exit a loop

```go
sum := 0
for i := 1; i < 5; i++ {
    if i%2 != 0 { // skip odd numbers
        continue
    }
    sum += i
}
fmt.Println(sum) // 6 (2+4)
```
