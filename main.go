package main

import "gin/routers"

func main() {
	var PORT = ":2121"

	routers.StartServer().Run(PORT)
}
