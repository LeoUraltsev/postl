package main

import "postl/internal/app"

const configPath = "./config/local.yaml"

func main() {
	app.Run(configPath)
}
