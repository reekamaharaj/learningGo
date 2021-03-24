# learningGo

Notes while learning Go basics with "Getting Started with Go" course on Coursera.

## Slice
A slice is a data type. They are flexible and can change size. Slices are a "window" into an array. Every slice will have an underlying array.

Pointer: start of slice

Length: number of elements in the slice

Capacity: maximum number of elements. Defined by the pointer (start of the slice) and the end of the array.

## Hash Tables
Data structure. Fast access to large bodies of data. Key value pairs. With unique keys. A hash table has a lot of slots and a hash function will generate the number of the slot to insert the value.

Allows access to things in constant time, don't need indices to access them. You can use keys, as long as they are unique and constant. Instead of using an array index, you can just use the unique key. Access in an array might be array[1] and in the hashtable instead of 1, the index, you would put the key hashtable[uniquekey]

Advantage: Faster lookup than lists, constant time vs linear-time, don't have to go through the whole list. Arbitrary keys, not integars like slices or arrays, keys can have a meaning and not just be a number.

Disadvantage: May have collisions, two keys hash to the same slot. Collisions can slow things down, collisions are rare and are usually (?) handled by Go

### Hash Function
Used to compute the slot for a key. This happens behind the scenes, inside Go. It is does not need to be called.

Function that takes as its argument the key and returns the slot where you want to put the value.

## Maps
Implementation of a hash table.
Use make() to create a map.
string is the key type, int is the value type

Create an empty map
`var idMap map [string] int`
`idMap = make(map[string]int)`

map literal, initialize with values
`idMap := map[string]int { "joe": 123}`

Access in a map. Similar to an array but in the [] will be the key.
`idMap["joe"]`

Add or change a key/value
`idMap["jane"] = 456`

Delete a key/value
`delete(idMap, "joe")`

Two-value assignment tests for existence of the key
`id, p := idMap["joe"]`
p will be a boolean. It will be true if the key is present in the map.
`len()` will tell you how many KV pairs are in the map.

Iterate through a Map
Use a for loop with a range. Two-value assignment with range
`for key, val := range idMap { fmt.Println(key, val) }`

## Structs (Structure)
Aggregate data type. Groups together objects of arbitrary data types into one object. Useful for organization.

Example. Person Struct
- Name, Address, phone
- Option 1: 3 seprate variables, programmer remembers how they are related
- Option 2: Make a single struc, which contains all variables. Bringing
together variables that are related.

`type struc Person {
    name string
    addr string
    phone string
}

var p1 Person
    `

Each property is a field. p1 contains values for all fields.
Name, addr, phone are fields.
Access fields in a structure. use dot notation
`p1.name = joe`
`x = p1.addr`

Initialize Structs
`new()`
initialized fields with zero
`p1 := new(Person)`
struct literal
`p1 := Person (name: "joe", addr: "address", phone: "123")`

## RFC (Request for Comments)
Definition of a protocols and formats. Standard definition for data.
- HTML (Hypertext Markup Language), RFC 1866
- URI (Uniform Resource Identifier), RFC 3986
- HTTP (Hypertext Transfer Protocol), RFC 2616
- JSON (Javascript Object Notation), RFC 7159

## Protocol Packages
- To encode and decode protocol format
- net: TCP/IP and socket programming
    TCP/IP: Define the internet. FTP, Secure Shell, others. Basic TCP/IP, UDP stack
- JSON: Attribute-value pairs
    - struct or map
    - Go struct
        `p1 := Person(name: "joe", addr: "a st.", phone: "123")`
    - JSON
        - `{"name": "joe", "addr": "a st.", "phone": "123"}`
        - All unicode
        - Human-readable
        - Fairly compact representation
        - Types can be combined recursively (Array of strcts, struct in struct)

## File Access
- Linear access, not random access
    - Mechanical delay
- Basic operations
    - Open: get handle for access
    - Read: read bytes into []byte
    - Write: write []byte into file
    - Close: release handle
    - Seek: move read/write head

## Interesting thing learnt
spaces are a problem for the strings package, needed bufio package to handle spaces

switch cases will automatically break after a successful case

Underscore _ : Blank identifier. Anonymous placeholder. May be used like any other identifier in a declaration, but it does not introduce a binding. Can ignore values, import for side effects, or silence compiler

- From JS to Go article: "only got functions and structs that map to functions"

## 25 Keywords in Go

### Declaration
- const
- var
- func
- type
- import
- package

### Composite types
- chan: define a channerl
- interface: specify method set, list of methods for a type
- map: unordered kv pair collection
- struct: collection of fields

### Control flow
- break
- case
- continue
- default
- else
- fallthrough: use in case of switch statement, the next case is entered
- for
- goto: jump to labeled statement without condition
- if
- range: iterate over items. (like a map or an array)
- return
- select: lets a goroutine to wait during the simultaneous communication operations
- switch

### Function modifier
- defer: defer the execution of a function until the surrounding function executes
- go: triggers goroutine, managed by goland-runtime

## Additional notes for written programs
For findIan
- [ ] What about symbols and numbers? Should test for that...
- [ ] Should make the code shorter, doesn't need to be as long as it is right now
- [ ] Error handling have not considered that yet, should do that
- [ ] Make a function to handle user input
- [ ] Slice code should check inputs, and just... needs more work

## Links to articles, videos, tutorials
- [JS to GO. Medium article](https://steevehook.medium.com/my-journey-from-javascript-to-go-9fb1e5d49fc2)
- [JS to Go YT video](https://youtu.be/r_nW-fMTXFw)
- [Coursera Getting Started with Go](https://www.coursera.org/learn/golang-getting-started/home/welcome)
- [25 keywords in Go](https://medium.com/wesionary-team/know-about-25-keywords-in-go-eca109855d4d)
- [Console Input bufio scanner & Type conversion](https://www.youtube.com/watch?v=1-bM3lSBDaA)