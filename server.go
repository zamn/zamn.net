package main

import (
	"fmt"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func mainPage(r render.Render) {
	r.HTML(200, "hello", "world")
}

func main() {
	/*
		session, err := mgo.Dial("localhost")
		if err != nil {
			panic(err)
		}
		defer session.Close()

		session.SetMode(mgo.Monotonic, true)
		usersTable = session.DB("mockr").C("Users")
	*/

	fmt.Println("Running server..")
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Layout:     "main",
		Extensions: []string{".tmpl", ".html"},
		Charset:    "UTF-8",
		IndentJSON: true,
	}))
	m.Get("/", func(r render.Render) {
		r.HTML(200, "index", nil)
	})

	m.Run()
}
