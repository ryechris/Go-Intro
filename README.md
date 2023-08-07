# Go-Intro
Introduction to Go

This repository houses some code I played with when learning Go:

1. Square Root Calculator
2. Reader Interface
3. Addition Calculator
4. Binary Trees Comparison function
5. Web Crawler

You can `git pull` the code into your machine,
to see for yourself what they yield:

1. Open the Terminal or Command Line app.

2. Create an empty folder in the location of your choice.

3. Navigate to the empty folder

4. Run the following code in that folder:
```
git init
git pull https://github.com/ryechris/Go-Intro.git
```

5. Then, to run `01-sqrt.go`, for example, you can run it on your Terminal or Command Line with this command:
```
go run 01-sqrt.go
```

(*If you don't already have Go installed, you can visit their [installation](https://go.dev/doc/install) page.)

The code are primarily answers to some problems on the tour given by the Go authors at go.dev.

For comments or inquiries, please contact:
riyan [at] linux.com



## Contents

### 1. Square Root Calculator
Here, we create a program that calculates the square root of a number, to illustrate these concepts:
- Flow Control
- Error Handling


### 2.  Reader: a Go library Interface
In its io package, Go provides the io.Reader interface.

As an exercise, we create our own struct (named MyReader), and we turn it into a Reader. 
This one imitates io.Reader, but emits an infinite stream of the character 'A'.

In so doing, we cover the following concepts:
- Interface
- Reader


### 3.  Summing 0 to 9: Concurrency with Go
Concurrency is built into Go; it is one of the reasons many people choose Go.

Summing the numbers -- such as 0,1,2,3,4,5,6,7,8,9 -- can illustrate how concurrency might be useful.
We can instruct the computer to perform this addition in one sequence.
However, with concurrency, we can divide the work into multiple threads, all runniing at the same time.
So a thread does less work, and thus completes in less time. Hence higher efficiency.

Here we do just that: we divide the sum into two threads; and by so doing we demonstrate the following constructions:
- Goroutines
- Channels


### 4. Binary Trees
Several binary trees can represent the same sequence of values.

In other languages, functions to check whether two binary trees hold the same sequence can be complex.

But with Go's concurrency features, we can make a simple solution.
Go provides the [tree](https://cs.opensource.google/go/x/tour/+/v0.1.0:tree/tree.go) package, where it defines this struct:
```
type Tree struct { // binary tree with integer values
    Left  *Tree
    Value int
    Right *Tree
}
```


### 5. Web Crawler
Here, we use Go's concurrency features to write a crawler that fetches URLs in parallel, without fetching the same URL twice.
