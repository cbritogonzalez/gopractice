# Basic notes
To use sub modules we will use a go.work file that will be created in the root of the project and then have go.mod files within each subfolder where we want to code
```
go work init
go work use <MODULE DIRECTORIES>
```
we add directories that have the go.mod files to add to the workspace

```
go mod init
```
to create the go.mod file within the directory


## Types
```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

## Short declarations
```
test := "Hi"
var test_bool := true
```
Using this we let go infer the type

int, uint, float64 are the most common, if needed we may look for another type annotation to optimize

# Extra
byte -> 8 bits
nibble -> 4 bits