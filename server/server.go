package server

import "github.com/labstack/echo"

func NewServer() *echo.Echo {
	return echo.New()
}
