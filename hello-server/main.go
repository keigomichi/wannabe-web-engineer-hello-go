package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// type jsonData struct {
// 	Number int    `json:"number,omitempty"`
// 	String string `json:"string,omitempty"`
// 	Bool   bool   `json:"bool,omitempty"`
// }

type jsonData struct {
	Number int
	String string
	Bool   bool
}

type Student struct {
	Number int    `json:"student_number"`
	Name   string `json:"name"`
}
type Class struct {
	Number   int       `json:"class_number"`
	Students []Student `json:"students"`
}

var val int = 0

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World.\n")
	})
	e.GET("/keigomichi", func(c echo.Context) error {
		return c.String(http.StatusOK, "keigomichiです。273系が発表されましたね。\n")
	})
	e.GET("/json", jsonHandler)
	e.GET("/hello2/:name2", helloHandler2)
	e.POST("/post", postHandler)
	e.POST("/hello/:name", helloHandler)

	// 演習問題
	// GET /ping
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong\n")
	})

	// GET /incremental
	e.GET("incremental", func(c echo.Context) error {
		val++
		return c.String(http.StatusOK, strconv.Itoa(val))
	})

	// GET /fizzbuzz
	e.GET("fizzbuzz", fizzBuzzHandler)

	// POST /add
	e.POST("/add", addHandler)

	// GET /students/:class/:studentNumber
	studentsJson := [...]string{
		`{"class_number": 1, "students": [
		  {"student_number": 1, "name": "Humming"},
		  {"student_number": 2, "name": "masutech16"},
		  {"student_number": 3, "name": "ninja"}
		]}`,
		`{"class_number": 2, "students": [
		  {"student_number": 1, "name": "hukuda222"},
		  {"student_number": 2, "name": "takashi_trap"},
		  {"student_number": 3, "name": "nagatech"},
		  {"student_number": 4, "name": "whiteonion"}
		]}`,
		`{"class_number": 3, "students": [
		  {"student_number": 1, "name": "yamada"},
		  {"student_number": 2, "name": "tubotu"},
		  {"student_number": 3, "name": "tsukatomo"}
		]}`,
		`{"class_number": 4, "students": [
		  {"student_number": 1, "name": "g2"},
		  {"student_number": 2, "name": "hatasa-y"}
		]}`,
	}

	var students [len(studentsJson)]Class
	for i := 0; i < len(studentsJson); i++ {
		if err := json.Unmarshal([]byte(studentsJson[i]), &students[i]); err != nil {
			fmt.Println("error")
		}
	}

	e.GET("/students/:class/:studentNumber", func(c echo.Context) error {
		class := c.Param("class")
		studentNumber := c.Param("studentNumber")
		i, _ := strconv.Atoi(class)
		j, _ := strconv.Atoi(studentNumber)
		return c.JSON(http.StatusOK, Student{
			Number: i,
			Name:   j,
			// Name:   students[i].Students[j],
		})
	})

	e.Start(":8080")
}

func fizzBuzzHandler(c echo.Context) error {
	count := c.QueryParam("count")
	val, err := strconv.Atoi(count)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	var res string
	for i := 1; i < val+1; i++ {
		switch {
		case i%15 == 0:
			res += "FizzBuzz\n"
		case i%5 == 0:
			res += "Buzz\n"
		case i%3 == 0:
			res += "Fizz\n"
		default:
			res += strconv.Itoa(i) + "\n"
		}
	}
	return c.String(http.StatusOK, res)
}

func jsonHandler(c echo.Context) error {
	res := jsonData{
		Number: 10,
		String: "hoge",
		Bool:   false,
	}

	return c.JSON(http.StatusOK, &res)
}

func postHandler(c echo.Context) error {
	var data jsonData

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, data)
	}
	return c.JSON(http.StatusOK, data)
}

func addHandler(c echo.Context) error {
	var data struct {
		Right int `json:"right"`
		Left  int `json:"left"`
	}

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Error string `json:"error"`
		}{Error: "Bad Request"})
	}
	return c.JSON(http.StatusOK, struct {
		Answer int `json:"answer"`
	}{Answer: data.Right + data.Left})
}

func helloHandler(c echo.Context) error {
	name := c.Param("name")
	return c.String(http.StatusOK, "Hello, "+name+".\n")
}

func helloHandler2(c echo.Context) error {
	name := c.Param("name2")
	return c.String(http.StatusOK, "Hello2, "+name+".\n")
}

// func studentsHandler(c echo.Context) error {
// 	class := c.Param("class")
// 	studentNumber := c.Param("studentNumber")
// 	return c.String(http.StatusOK, struct{
// 		StudentNumber int
// 		Name string
// 	}{
// 		StudentNumber: studentNumber,
// 		Name: students
// 	})
// }
