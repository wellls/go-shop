package main

import "net/http"

func (app *application) VirtualTerminal(rw http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("Hit the handler")
}
