package main

import (
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	s := echo.New()

	s.Use(MW)

	s.GET("/status", Handler)

	err := s.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func MW(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		val := ctx.Request().Header.Get("User-Role")

		if val == "admin" {
			log.Println("red button user detected")
		}

		err := next(ctx)
		if err != nil {
			return err
		}
		return nil
	}
}
