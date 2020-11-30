package main

func square() func() int {
    var x int
    return func() int {
        x++
        return x * x
    }
}

func main() {
    f := square()
    f()
    f()
}
