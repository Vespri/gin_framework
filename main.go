package main

import "gin/routers"

func main() {
	var PORT = ":1234"

	routers.StartServer().Run(PORT)
}
