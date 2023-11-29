package main
import (
	"fmt"
	"strconv"
	"github.com/memphisdev/memphis.go"
)
func CheckSeverity(message []byte, headers map[string]string, inputs map[string]string) ([]byte, map[string]string, error) {
	numInt , _ := strconv.Atoi(inputs["name"])
	fmt.Println("input", 5/numInt)
        return message,  inputs, fmt.Errorf("sveta error")
}
func main() {
	fmt.Println("hello world")
	memphis.CreateFunction(CheckSeverity)
}
