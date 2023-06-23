package main

import (
	tmp_session "tmp_server"
)

func main() {
	server := new(tmp_session.SessionServer)
	server.Run()
}
