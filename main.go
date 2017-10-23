package main

import (
	"flag"
	"github.com/leogsouza/go-chat/defaultport"
	"github.com/leogsouza/trace"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

// templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {

	port := utils.DefaultPort()
	var addr = flag.String("addr", ":"+port, "The addr of the application.")
	flag.Parse() // parse the flags

	// set up gomniauth
	gomniauth.SetSecurityKey("I was born to win!")
	gomniauth.WithProviders(
		facebook.New("key", "secret", "https://go-chat-leogsouza.c9users.io/auth/callback/facebook"),
		google.New("927100418429-9bssnoqr5f94483d7ei6mp6vhuiirrgv.apps.googleusercontent.com", "O5ndCaPaCwobNcPeE-HoR6O8", "https://go-chat-leogsouza.c9users.io/auth/callback/google"),
		github.New("key", "secret", "https://go-chat-leogsouza.c9users.io/auth/callback/github"),
	)

	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)

	// get the room going
	go r.run()

	// start the web server
	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
