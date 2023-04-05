// "\main.go"
package main
import (
	"github.com/2003NURS/Golang-CourseProject.git/app/space"
)

func main(){
	println("Hello")
	println("http://localhost:8000/welcome")

	space.HandlerRequest()
}