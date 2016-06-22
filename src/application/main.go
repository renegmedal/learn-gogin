package main

func main() {
	r := registerRoutes()

	// or use this
	// r.StaticFS("/public", http.Dir("./public"))

	r.Run(":3000")

}
