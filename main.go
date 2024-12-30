package main

func main() {
	server := NewAPIServer(":9090")
	server.Run()
}
