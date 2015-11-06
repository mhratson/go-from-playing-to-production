package main
import "fmt"
func main() {
	fmt.Println(hello("go"))
}
func hello(one string) string {
	return fmt.Sprintf("hello %s world", one)
}
