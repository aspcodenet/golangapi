package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Employee struct {
	Id   int
	Age  int
	Namn string
	City string
}

func onHandleEmployee(c echo.Context) error {

	allaEmployees := []Employee{}
	allaEmployees = append(allaEmployees, Employee{Id: 1, Namn: "Stefan", Age: 25})
	allaEmployees = append(allaEmployees, Employee{Id: 2, Namn: "Oliver", Age: 5})
	allaEmployees = append(allaEmployees, Employee{Id: 3, Namn: "Josefine", Age: 12})

	//	s, _ := json.MarshalIndent(allaEmployees, "", " ")

	return c.JSON(http.StatusOK, allaEmployees)
}

func main() {

	e := echo.New()

	e.GET("/api/employee", onHandleEmployee)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/contact", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<h1>Please do NOT contact us</h1>")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
