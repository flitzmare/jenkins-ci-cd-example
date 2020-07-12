package jenkins_ci_cd_example

import (
	"fmt"
	"github.com/labstack/echo"
)

func main(){
	fmt.Println(Print())
	fmt.Println(Print2())

	e := echo.New()
	e.GET("/", func(context echo.Context) error {
		return context.String(200, Print())
	})
	e.Logger.Fatal(e.Start(":1234"))
}

func Print() string {
	return "Hello World"
}

func Print2() string {
	return "Nice bro"
}
