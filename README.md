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

