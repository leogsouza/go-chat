package main

import (
	"github.com/leogsouza/go-chat/defaultport"
	"log"
	"net/http"
)

func main() {

	port := utils.DefaultPort()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
            <html>
                <head>
                    <title>Chat</title>
                </head>
                <body>
                    Let's chat!
                </body>
            </html>
        `))
	})

	// start the web server
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
