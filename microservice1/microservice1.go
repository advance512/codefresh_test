package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"net/http"

	"github.com/labstack/echo/middleware"
	"github.com/parnurzeal/gorequest"

	"time"
	"fmt"
	"io/ioutil"
)

func uploadRAML(c echo.Context) error {

        fmt.Println("Handler for /v1/uploadRAML called.")

	fileHeader, err := c.FormFile("file")
	if err != nil {
		return err
	}
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	fileContents, err := ioutil.ReadAll(file)
	fileContentsString := string(fileContents[:])

	request := gorequest.New()

	_, body, errs := request.Post("http://ms2:1324/v1/verifyRAML").
		Send(fileContentsString).
		End()

	if errs != nil {
		return errs[0]
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<html><body>Result for file submitted on %s:<br><br><pre>%s</pre></body></html>", time.Now().Format(time.RFC850), body))
}

func showWelcomePage(c echo.Context) error {
        fmt.Println("Handler for /v1/index called.")

	return c.HTML(
		http.StatusOK,
		`<html>
			<body>
				Welcome to this ugly form. Please upload a RAML file.<br><br>
				<form action="/v1/uploadRAML" method="post" enctype="multipart/form-data">
		    			<div>
		        			<label for="file">File:</label>
		    				<input type="file" name="file"><br><br>
	    					<input type="submit" value="Submit">
					</div>
				</form>
			</body>
		</html>`)
}

func main() {
	e := echo.New()

	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
                fmt.Println("Handler for / called.")
		return c.Redirect(301, "/v1/index")
	})

	e.GET("/v1/index", showWelcomePage)

	e.POST("/v1/uploadRAML", uploadRAML)

	fmt.Println("Microservice 1 is now listening on port 1323...")
	e.Run(standard.New(":1323"))
}

