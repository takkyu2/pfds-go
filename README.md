# Purely Functional Data Structures in Go
An implementation of the immutable data structures in Okasaki's textbook, "Purely Functional Data Structures", in Go.

Still work in progress.

## Requirements
A Go compiler supporting Generics, e.g., Go 1.18, accessible by gotip as of now, is necessary.

## Disclaimers
* Since Go 1.17 with `-gcflags=-G=3` flag does not support exporting Generic functions, types, etc., the code cannot be compiled by Go 1.17. I could put everything in a single main package, but that would hurt the readability.
* Polymorphic recursion seems impossible in Go generics (compiler hangs up) :(

## Others
* I guess Go's garbage collection and easy concurrency may be a good fit for PFDS :)
* I want to compare the time and memory efficiency of PFDS in Go to the ones of [PFDS in C++](https://github.com/takkyu2/purely-functional-data-structures), but, as you know, benchmarking is difficult, so I am not sure I can make it.
