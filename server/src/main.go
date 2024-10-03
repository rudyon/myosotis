package main

import (
    "log"

    "github.com/labstack/echo/v5"
    "github.com/pocketbase/pocketbase"
    "github.com/pocketbase/pocketbase/core"
)

func main() {
    app := pocketbase.New()

    // Custom API endpoint
    app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
        e.Router.GET("/hello", helloHandler)
        return nil
    })

    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}

// helloHandler responds with a simple message
func helloHandler(c echo.Context) error {
    return c.JSON(200, map[string]string{"message": "hello from go!"})
}

