# golang
## build
```
go build helloworld.go
```
## os.Args
os.Args[0] - output name of performed program
## for
tradicional for
```
for initializarion; condition; aftereffect {
    // zero or more instructions
}
```
traditional loop while
```
for condition {
    // choose your destiny...
}
```
traditional endless loop
```
for {
    // choose your destiny...
}

## define variables
```
s := ""             // can define only in function
var s string        // by default for using
var s = ""          // for 2 or more variables
var s string = ""   // so unnecessarily
```
## nil
like None in Python

## Verbs
%d          // decimal integer       
%x,%o,%b    // integer 16,8,2
%f,%g,%e    // float: 3.141593, 3.141592653589793, 3.141593e+00
%t          // bool
%c          // rune(Unicode)
%s          // string
%q          // output like "abc" or 'c'
%v          // anyway value
%T          // anyway type
%%          // %