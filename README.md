# learningGo

Notes while learning Go basics with "Getting Started with Go" course on Coursera.

## Slice
A slice is a data type. They are flexible and can change size. Slices are a "window" into an array. Every slice will have an underlying array.

Pointer: start of slice

Length: number of elements in the slice

Capacity: maximum number of elements. Defined by the pointer (start of the slice) and the end of the array.

## Hash Tables

## Interesting thing learnt
spaces are a problem for the strings package, needed bufio package to handle spaces

switch cases will automatically break after a successful case

## Additional notes for written programs
For findIan
- [ ] What about symbols and numbers? Should test for that...
- [ ] Should make the code shorter, doesn't need to be as long as it is right now