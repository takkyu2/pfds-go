# Purely Functional Data Structures in Go
An implementation of the immutable data structures in Okasaki's textbook, "Purely Functional Data Structures", in Go.

## Requirements
A compiler supporting Generics, e.g., Go 1.17 with `-gcflags=-G=3` flag, is necessary.

## Disclaimers
* I want to compare the time and memory efficiency of PFDS in Go to the ones of [PFDS in C++](https://github.com/takkyu2/purely-functional-data-structures), but, as you know, benchmarking is difficult, so I am not sure I can make it.
* Since Go 1.17 does not support exporting Generic functions, types, etc., I put everything in a single main package.
