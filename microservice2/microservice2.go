package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"net/http"

	"github.com/labstack/echo/middleware"

	"github.com/kr/pretty"
	raml "gopkg.in/raml.v0"

	"fmt"
	"io/ioutil"
	"os"
)

func verifyRAML(c echo.Context) error {
        fmt.Println("Handler for /v1/verifyRAML called.")

	requestBody, err := ioutil.ReadAll(c.Request().Body())

	if err != nil {
		c.Error(err)
		return err
	}

	tmpfile, err := ioutil.TempFile("", "ramltmp")
	if err != nil {
		c.Error(err)
		return err
	}

	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(requestBody); err != nil {
		c.Error(err)
		return err
	}

	if err := tmpfile.Close(); err != nil {
		c.Error(err)
		return err
	}

	if apiDefinition, err := raml.ParseFile(tmpfile.Name()); err != nil {
		return c.String(http.StatusOK, err.Error()) // Application error, not a runtime error
	} else {
		result := fmt.Sprintf("Successfully parsed RAML file:\n\n%s", pretty.Sprintf("%s", apiDefinition))
		return c.String(http.StatusOK, result)
	}
}

func main() {
	e := echo.New()

	e.Use(middleware.Recover())

	e.POST("/v1/verifyRAML", verifyRAML)

	fmt.Println("Microservice 2 is now listening on port 1324...")
	e.Run(standard.New(":1324"))
}

