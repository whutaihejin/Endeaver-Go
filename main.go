package main

func closure() func() int {
    var x int
    return func() int {
        x++
        return x * x
    }
}

func main() {
    f := closure()
    f()
    f()
}
