Here's a list of the null values for the different types:

int => 0
float64 => 0.0
string => "" (i.e., an empty string)
bool => false


----
- go run . // in case multiple mains are declared will throw error
- go run <file_name>

- go mod init go-with-maximilian/best-practices
- defineing variable outside main can't use shorter symtax ":=" to init variable

- Splitting Code to Multiple Files under same package 
 - functions can be directly be used or imported under same package
- Folder name should be SAME as of PACKAGE name NOT FILE name
- variable/functions/constants/struts starts with PascalCasing (UpperCase) are availaible to other packages
- default EXPORTS in case starts with CapitalLetter any function/varibale/constant name
- How to Install Packages:
 - e.g go get github.com/Pallinder/go-randomdata
 - Updates go.mod file to add pkg
 - Adds (first time) / Updates with hash
- `go.mod -> Package + 3rd Party dependencies are listed here
- It is important in go lang to handle return value when calling from caller 


## Pointers
- Directly mutate copies
- Avoiding Unncessary Value Copy
 - By default `Go` creates a copy when passing values 
   to function
e.g https://gyazo.com/c30dd05fb452f2707734679de88a3567


func outputAge(age) -> byValue -> creating a copy of variable address
func outputAge(&age) -> ByRef  -> mutating same address


- All values in Go have a so-called "Null Value" - i.e., the value that's set as a default if no value is assigned to a variable.
 - e.g 
    - int variable is 0. 
    - float64, it would be 0.0.
    - string, it's "".
    - pointer `nil` // special built into Go
      - nil represents the absence of an address value - i.e., a pointer pointing at no address / no value in memory.
      - In go you can't perform operations on pointer until you dereference value
- In GoLang PackageName and func name can't be same


## Structs
- Group related data together using Structs
- Struct in go are called either "Struct Literal" or "Composite Literal"
- Pointer technically derefernce the structs internally via short way 
   e.g u.firstName otherwise it should be like (*u).firstName
- Reciever function is defined which automatically binds to struct level
- Go also has byValue/ByReference concept when passing struts to either 
   - Reciever function     (Link methods with structs)
   - Method with Arguments (Not Linked to specific structs)
- Go Provides to have Constructor function with `new` Convention only not by Go
- Structs allowed validation logic very easy to do with in Constructor function

// Capital P in Person make it exportable outside package also the fields inside package 
 too e.g "Name/Age" all are exportable by go due to PascalCasing .
  - This is only valid when structs lives in their own packages and needs to be exported to 
   other packages

type Person struct {
	Name string
	Age int
	Email string
	CreatedAt time.Time
}

- Structs has no classes and no Inheritance but it has concept called Struct Embedding aka 
 using of 1 struct to another similar like we try to use existing class fields/methods in OOP 
 based language we can do something like in structs too see eg. struct_embedding.go

- Embedding Structs

 type User struct {
	// Explicitly set to lowercase to make it unexported
	name  string
	age   int
	email string
}

type Admin struct {
	User
	role string
}

- With Annoymous Structs embedding e.g User instead of user User you can directly use reciever functions of User struct 
from admin instance rather then creating wrapper reciever functions for Admin Struct e.g see `user_struct.go`


- Struct tags are essentially metadata that you can add in structs fields e.g

type Note struct {
   Title: string          `json:"title"` // backticks and "property"
   Content string         `json:"content"`
   CreatedAt: time.Time   `json:"created_at"`
}

- You can build Custom Types and Reciever functions can be attached to custom types e.g custom_types.go

---



## Interfaces

- You can define interfaces having methods that can be common can be defined in interfaces
- Methods define in interface should be implemented individual in files otherw go will complain
 e.g  X does not implement Y Interface

 See e.g interface.go 
  - SaveFile // method and custom business logic should be impl into respective files
  - In Go Interface not needs to be explicitly connected with other types
  - Go Validates whatever argument is passed if it does satisfy contract IMPL in concreate file it will 
  will call respective method from respective concreate file to what ever it is passed.
   - Think like instancs of Class it calls it respective method based on callee is called from instance
  - Wrapper Func e.g SaveData accepts Interface argument 
  - So Whenever this Wrapper Func is called with Passed args Go assume "Accpeted" argument 
   implements that interface contract in concrete class



Cursor Explanation:


1: Embedding interfaces/structs: When you include an interface or struct name (without a field name) in your struct definition, you're embedding it. This is what you did with Saver and Printer in your struct.


```go
type EmbeddedSaverPrinter struct {
    Name string
    Saver    // This is embedding the Saver interface
    Printer  // This is embedding the Printer interface
}
```

Wrong:

What you were trying to do with CustomDisplayer() inside the struct definition isn't valid Go syntax. 
 You can't declare method signatures inside a struct - those belong in interfaces.

```go
type EmbeddedSaverPrinter struct {
    Name string
    Saver    // This is embedding the Saver interface
    Printer  // This is embedding the Printer interface
    CustomDisplayer()
}
```

2: Defining methods: Methods are always defined outside the struct using the receiver syntax.

```go
func (e *EmbeddedSaverPrinter) CustomDisplayer() {
    fmt.Println("Custom Displayer...")
}
```

3:Inside your struct, this isn't valid Go syntax. You can embed types (interfaces or structs), but you can't embed method signatures or implementations.
When you embed an interface like Saver, you're saying "this struct should have all the methods from Saver available" - but you still need to implement those methods outside the struct definition. Embedding doesn't automatically implement the methods.
So yes, your struct can implement multiple interfaces and have additional methods, but all method implementations must be defined outside the struct definition using the receiver syntax.


4: Correct way to add declaration in interface and embed interface in struct NOT in struct directly

if you want to declare that a struct should have a certain method, you would:
  - Create an interface with that method signature
  - Either embed that interface in your struct or just implement the method
  - Implement the method outside the struct definition
The Go compiler doesn't need or allow method declarations inside structs because it uses "duck typing" - a struct implements an interface automatically if it has all the required methods, without needing to explicitly declare that relationship.

Conclusion
- if your struct has methods matching an interface's method signatures, it automatically implements that interface.
- The compiler will check if your struct implements all required methods at compile time, but there's no way to declare method signatures inside the struct definition itself.

## Install Packages
- go get github.com/Pallinder/go-randomdata

Scan vs ScanLn 
 - Scan doesn't closes stdin when pressing enter
 - ScanLn does close stdin when pressing enter 

 ---

 - use of .(type) outside type switch is only for switch statement to get type see e.g special_interface.go file
 - Alternate syntax see e.g special_interface.go file

 - slices [2:4] StartingIndex is included but values are included till EndingIndex-1
 - you can't put index with higher bound in array e.g array[7] and you have only 6 len arrray
   e.g panic: runtime error: slice bounds out of range [:10] with capacity 4
 - you can't put -ve index in slice e.g array[-2] which works in JS but not in Go
 - slices is by ref and can update value in original array
  - you create 1 array in memeory and slices is just small window or reference to part of array 
    and is pointing to same memory reference as original array.
  - In slices cap is count towards end of the array based on slices is cretaed from
  - In slices cap is omitted on left side


  ## Maps

  - Maps are always dynamic no Fixed size can be created
  - Maps are mutable in go
  - Deletion/Additon in Maps are done in original struct
  - Access values in map using key instead of Indexes
  - Map vs Struct
    - you can't delete a key/value pair from a struct
    - struct is good when you clear set of data 
    - map is good when properties can be added/removed dynamically

```go
type Website struct {
	Google string
	Facebook string
	Twitter string
	Reddit string
	Youtube string
	Instagram string
	Linkedin string
}
// vs 

var websites := map[string][string]

/*
Reason-1:
How it is easy to save ANY site value in map dynamically using maps where 
in struct every different site has to define different website key/value 
and define each value into it.

Reason-2:

In Struct you can't have e.g

type Website struct {
	AWS WEBSITE URL string
}

In Map You can have

e.g 
websites["AWS WEBSITE URL"] =  "https://aws.amazon.com/"
*/

```

- In go when you do `[]string{}` you create a slice a dynamic slice
- In go when you do `[3]string{}`you create an an array of fixed size and you can't append dynamically you would need to create a slice to append it
- In go you can have special keyword called `make` which can preserved some slots and appending will done AFTER WORDS of preserved slots.
  e.g  
    - make([]string, 3) // Makeing dynamic slice with 3 preserved slot
    - make(map[string]string, 10) // We can do the same with maps which is more memory efficent instead and instead of assigning it will 
     have preserved memory (no relocate of memory) and in case maps becomes bigger it will try to use e.g 10 length before then 
     after it in case 11th it will then reassign a new Map again IMO with 11 reserved which will be same as before without make(map[string]string, len)
     and is not memory efficient


## Functions
- In functions we can pass functions as parametrs value for other funcitons
- functions are just values in go
- value you can execute when treated as function
- value you can use as simple native int/string simple value
- value you can pass around to functions



Gotchas

- Under same package go doesn't allow to have 
 e.g functions name or type to be same in 1 package
