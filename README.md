# Challenge - Backend Developer
Write a program that prints all the numbers from 1 to 100. However, for
multiples of 3, instead of the number, print "Linio". For multiples of 5 print
"IT". For numbers that are multiples of both 3 and 5, print "Linianos".
But here's the catch: you can use only one `if`. No multiple branches, ternary
operators or `else`.

# Requirements
* 1 if
* You can't use `else`, `else if` or ternary
* Unit tests
* Feel free to apply your SOLID knowledge
* The language used for this challenge must be GoLang

### Build From Source

1. [Install Go](https://golang.org/doc/install)
2. Clone the repository:

   ```shell
   git clone https://github.com/garciaolais/challenge-backend.git
   ```

3. Run `make` from the source directory

   ```shell
   cd challenge-backend
   make build
   ./challenge
   ```
### Run Tests

```shell
   cd challenge-backend
   make test
    go test -v test/*.go -short
    === RUN   TestDivideByThree        
    --- PASS: TestDivideByThree (0.00s)
    === RUN   TestDivideByFive
    --- PASS: TestDivideByFive (0.00s)        
    === RUN   TestDivideByFiveOrThree        
    --- PASS: TestDivideByFiveOrThree (0.00s)
    === RUN   TestRun
    --- PASS: TestRun (0.01s)
    === RUN   TestGetOnlyInts
    --- PASS: TestGetOnlyInts (0.01s)
    PASS
    ok      command-line-arguments  0.011s 
```

### Build Docker Image

```shell
   cd challenge-backend
   make docker-build
   docker run --rm challenge:1.2
   1
   2
   Linio
   4
   IT
   Linio
   .
   .
   .
   97
   98
   Linio
   IT
```

### Getting Started

See usage with:
```shell   
    challenge -help
```

Run from 1 to 100:
```shell   
    challenge
```

Run custom list:
```shell   
    challenge -first 1 -last 10
    1
    2
    Linio
    4
    IT
    Linio
    7
    8
    Linio
    IT
```

## DEMO

To see this library in action check the implementation as an API

[Challenge-Full](https://github.com/garciaolais/challenge-full)

https://user-images.githubusercontent.com/14133711/141163375-0b7ceb31-1c95-49fc-9a64-65aa148791b0.mp4