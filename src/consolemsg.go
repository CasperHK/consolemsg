package consolemsg

import (
    "fmt"
    "strings"
)



func Err(msg string) {
	fmt.Println(" [Error]: ", replaceNewLine(msg))
}

func Log(msg string) {
	fmt.Println("   [Log]: ", replaceNewLine(msg))
}


func Sys(msg string) {
	fmt.Println("[System]: ", replaceNewLine(msg))
}

func CusMsg(keyword, msg string) {

}

var clear map[string]func() //create a map for storing clear funcs

func init() {
    clear = make(map[string]func()) //Initialize it
    clear["linux"] = func() { 
        cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
        cmd := exec.Command("cls") //Windows example it is untested, but I think its working 
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}

func Clr() {
    value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    } else { //unsupported platform
        panic("Your platform is unsupported! I can't clear terminal screen :(")
    }
}

func replaceNewLine(msg string) string {
	newMsg := strings.Replace(msg, "\n", "\n          ")
	return newMsg
}