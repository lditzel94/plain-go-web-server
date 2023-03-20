package main

func main() {
	server := NewServer(":4400")
	server.Handle("GET", "/", RootHandler)
	server.Handle("POST", "/create", PostRequest)
	server.Handle("POST", "/user", UserPostRequest)
	server.Handle("POST", "/api", server.AddMiddleware(HomeHandler, CheckAuth(), Logging()))
	server.Listen()
}
