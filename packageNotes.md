# Go Package Notes

## fmt: Format Package
- Implements formatting for input and output.
- Format 'verbs' derived from C, but simpler. Print with Printf
    - %q: single-quoted character literal
    - %d base 10
    - %T type value
    - %v value in a default format

## os: Operating System
- A platform-independent interface to operating system functionality
- more control over file access than ioutil
- os.Open(): Opens a file, returns a file descriptor (File)
    - file descriptor, used to access the file
- os.Close(): Closes a file
- os.Read(): Reads from a file into a []byte, fills the [] byte, contorl the amount read based on the []byte size
- os.Write(): writes a []byte into a file

- Open, Read, Close
    - `f, err := os.Open("dt.txt")`
    - `barr := make([]byte, 10)`
    - `nb, err := f.Read(barr)`
    - `f.Close()`
    - nb: number of bytes read
    - Head: where in the file something is, will increment with each Read. Close will reset the head
    - Read and fills barr
    - Read returns # of bytes read
    - May be less than []byte length

- Create, Write
    - `f, err :=`
    - `os.Create("outfile.txt")`
    - `barr := []byte{1, 2, 3}`
    - `nb, err := f.Write(barr)`
    - `nb, err := f.WriteString("Hi")`

    - WriteString() writes a string
    - Write() writes a []byte, Any unicode sequence

## bufio: Buffered Input/Output
- Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer object, creating another object (Reader or Writer) that also implements the interface but provides buffering and some help for textual I/O.
- `scanner := bufio.NewScanner(os.Stdin)`
    - Create a NewScanner object as "scanner"
    - Object has os.Stdin (Stdin: what is input on the command line)
- `scanner.Scan()`
    - Scan the user input
    - Scanned as a string
- `input := scanner.Text()`
    - Save the scanned input as "input"

## strings

## strconv: String Conversion
- Implements conversions to and from string represendation of basic data types
- `input, _ := strconv.ParseInt(scanner.Text(), 10, 64)`
Will convert the input string to an integer at base 10 and a size of 64.
_ will ignore an error if there is one and continue

## math
- `math.Trunc()`

## json
- JSON marshalling: Generating JSON representation from an object
` type struct Person{`
    `name string`
    `addr string`
    `phone string`
`}`
`p1 := Person(name: "joe", addr: "a st.", phone: "123")`
`barr, err := `
`json.Marshal(p1)`
- Marshal() returns JSON representation as []byte
- JSON Unmarshalling: Convert a JSON[]byte into a Go object
`var p2 Person`
`err := json.Unmarshal(barr, &p2)`
Person struct has to fit the JSON[]byte. Key in Person struct needs to match the attribute in the JSON representation. Error will be thrown if it doesn't fit

## ioutil package
- file access package
- "io/ioutil" has basic functions

    ReadFile
    `dat, e :=`
    `ioutil.ReadFile("test.txt")`
    - dat is []byte filled with contents of entire file
    - ReadFile will open and close the file, will work for small files

    WriteFile
    `dat = "Hello, world`
    ` err := ioutil.WriteFile`
    `("outfile.txt", dat, 0777)`

- writes []byte to file
- creates a file
- Unix-style permission bytes (0777, gives permission to read/write)
- WriteFile creates a file, dumps whatever it needs to into the file and then closes the file. Not used for appending a file.

## net/http
- web communication protocol
- net: TCP/IP and socket programming
    - TCP/IP: Define the internet. FTP, Secure Shell, others. Basic TCP/IP, UDP stack