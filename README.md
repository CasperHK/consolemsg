# ConsoleMsg
A console package for go lang

<br/>

## Screenshot

<br/>

## Basic Usage
First of all, we have to import the package.
```go
import (
    console "github.com/CasperHK/consolemsg"
)
```

```go
func main() {
    console.Err("1")
    console.Log("2")
    console.Sys("11")
    console.CusMsg('CusMsg', "3") / only accept characters in 5
}
```
The output in the terminal in like the following.
```bash
 [Error]: 1
   [Log]: 2
[System]: 11
[CusMsg]: 3
```

<br/>

## Features
* Better manage your console print out.
* Highligh the type of print out with different colors by https://github.com/fatih/color.

<br/>

## License

