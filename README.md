# Go programming language
This is a doc aimed for quick basics of Go

## How to run a .go file
Type  ```go run <filename.go>``` in your command line. (Of course after downloading and installing the software from [their website](https://www.golang.org)

1. ## Hello world
```go

package main
import "fmt"

func main(){
  fmt.Println("Hello World!")
}
```

2. ## Variables, Comments, Functions
```go

// First way
var a float32 = 2.2
var b float32 = 4.4

// Second way
var a, b float32 = 2.2, 4.4

// Third way
a, b := 2.2, 4.4

/* * Functions * */
func <name>(<args with type>) <return type>{ <body> }

//One way
func add(x float32, y float32) float32{
  return x + y
}

//Another way
func add(x, y string) (string, string){
  return x,y //Returns both the strings
}
fmt.Println(add("Hello","There")) // Prints **Hello There**
```
* *Notes*
    * Static typing, declared types don't change after compilation
    * External function have to be called till the subfolder/submodule in which the exist
    * Println space separated the arguments passed i.e the two strings returned from function get space separated
    * Run ```godoc <full-package-folder-name> <function-name>``` to get more details on a function

3. ## Pointers
```go

a := 5
x := &a // x now has the memory address of a
*x = 2 // a now has the value 2. Because change was made at the address of a
```
* *Notes*
    * Pointers work the same way they do in C/C++

4. ## Loops
```go
// Go has only one kind of loop i.e. *for* loop

// Method 1
for var i := 1; i < 10; i++{
  fmt.Println(i)
}

// Method 2 (Equivalent of a while loop)
var i := 1
for ; i < 10 ;{
  fmt.println(i)
  i += 1
}
```

5. ## Datatypes
```go

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point
     // In other words: a character

float32 float64

complex64 complex128

// Eg to print a type and value of a variable :
fmt.Printf("Type is %T and value is %v\n", x, x)
```

6. ## Concurrency
```go

func exec(s: string){
  for var i := 1; i <= 3; i++{
    fmt.Println(s)
    time.Sleep(time.Millisecond * 100)
  }
}

// Method 1 (without routines)
exec("Yo")
exec("Ma Man")
/* Output --
  Yo
  Yo
  Yo
  Ma Man
  Ma Man
  Ma Man
*/

// Method 2 (with routines)
go exec("Yo")
exec("Ma Man")
/* Output --
  Ma Man
  Yo
  Ma Man
  Yo
  Ma Man
  Yo
*/
```
* *Notes*
    * Function calls with goroutines makes them (the functions) execute in a separate thread than the whole program. This gives us the benefit of concurrency in Go

7. ## Maps
```go

occurences := make(map[string]int)

// Setting values
occurences["A"] = 12 // A occurs 12 times
occurences["ABC"] = 2

// Deleting values
delete(grades,"A")

```
* *Notes*
    * Maps are reference types (not like int or float or string). This means that ```var a = map[int]int``` behaves as an empty map but working on it will result in panic
    * Therefore we use ```make()``` to allocate and initialise an empty hash map that we can now work on.

8. ## Arrays
```go

var s [5]string
var friends := [3]string{"Popular", "Bookworm", "Funny one", "Fitness freak"}

fmt.Println(friends)

for index, item := range friends{
  fmt.Printf("Friend %d : %s", index, item)
}
```
* *Notes*
    * Arrays in go have fixed length; Once declared, the length cannot change

9. ## Slices
```go

sl := make([]int, num) // num is the number of elements in the list

primes := [6]int{2, 3, 5, 7, 11, 13}

var s []int = primes[1:4]
s[0] = 99

// primes is now {2, 99, 5, 7, 11, 13}

var q = []int{1,3,5,7,9,11,13}
fmt.Println(q)  // [1 3 5 7 9 11 13]

q = q[:0]
fmt.Println(q)  // []
```
* *Notes*
    * Slices are dynamically sized views into the elements of an array
    * every slice is declared with two indices like ```[start : end]```
    * Slices do not actually contain data. They are references to the elements of an array
    * If the elements of a slice are changed, the corresponding elements of the array are changed too
    * You can get the length of slice with ```len()``` and capacity of slice (the length of underlying array) with ```cap()```
    * Try [this](https://tour.golang.org/moretypes/11) page for a topic necessary to understand
