package jenkins_ci_cd_example

import "fmt"

func main(){
	fmt.Println(Print())
	fmt.Println(Print2())
}

func Print() string {
	return "Hello World"
}

func Print2() string {
	return "Nice bro"
}
