# 2020 Advent of Code (Golang)

This repository implements solutions to the puzzles in the [2020 Advent of Code](https://adventofcode.com/2020) using Golang.


## Preface

This was a vehicle to learn Golang, so I presume not everything done here will be deemed idiomatic by Golang specialists.

Generally speaking, the solutions are organised predominantly for comprehension. They strive to arrive at an answer in a reasonable period of time, but they typically prioritise optimal understanding over optimal performance.

The examples are representative of my thinking and coding style.


## Getting Started

### Prerequisites

The project requires `golang 1.15.5`, but any reasonably current version of Golang will likely work.  I tend to code done the middle of any language specification.

If you use a Golang manager that responds to `.tool-versions`, you should be switched to `1.15.5` automatically. I recommend [ASDF](https://github.com/asdf-vm/asdf) for those on platforms that support it.

### Installation

You will need to install two Golang packages.

```
$ go get -u github.com/elliotchance/pie
$ go get -u github.com/franela/goblin
```

### File Structure

- [cmd](./cmd):   Runner for daily solutions
- [data](./data): Puzzle input organised by day
- [doc](./doc):   The calendar and puzzle descriptions
- [pkg](./pkg):   Daily solutions and other homegrown utilities


### Running Daily Solutions

Modify the runner in the `cmd` directory to import and execute the daily
package of your choice.

Then execute the following command in your terminal from the project root.

```
$ go run cmd/main.go
```

### Running Tests

The only tests are a set of checks to verify solved puzzles.

I often refactor my solutions for clarity (or as I learn new
techniques in subsequent puzzles), so it is helpful to have
these simple tests to give my refactors some confidence.

To execute the tests, simply execute the following command in
your terminal from the project root.

```
$ go test
```
