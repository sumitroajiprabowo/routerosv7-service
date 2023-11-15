// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 GitHub Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"github.com/sumitroajiprabowo/routerosv7-service/app"
	_ "github.com/sumitroajiprabowo/routerosv7-service/docs"
)

// @title RouterOS v7 Service API Documentation
// @description This is a sample server for RouterOS v7 Service API Documentation.
// @version 1.0.0
// @host localhost:3000
// @BasePath /api/v1
// @contact.name Megadata Pemalang
// @contact.url https://github.com/megadata-dev
// @contact.email danu@megadata.net.id
// @license.name MIT
// @license.url https://github.com/megadata-dev/routerosv7-service/blob/main/LICENSE
func main() {
	app.InitServerHTTP()
}
