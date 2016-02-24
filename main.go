package main

import (
	"github.com/kraken97/configEditor/Reader"
	"github.com/kraken97/configEditor/Utils"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"strings"
)

func main() {
	utils.Copy(`/home/kraken/.config/i3/bacp`, `/home/kraken/.config/i3/config`)
	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/", indexHandler)

	m.Post("/save", saveHandler)
	m.Run()
}

func saveHandler(r *http.Request, rnd render.Render) {
	str := getStr(r)
	file := readU.GetFile()
	sStr := "#gostart"

	start := strings.Index(file, sStr) + len(sStr)

	end := strings.Index(file, "#goend")
	WriteString := file[:start] + str + file[end:]
	readU.Write(WriteString)
	rnd.Redirect("/")
}
func getStr(r *http.Request) string {
	str := "\n"
	for i := 0; i < len(Data.Info); i++ {
		str += Data.Info[i].Name + "#" + r.FormValue(Data.Info[i].Name) + "\n"
	}
	return str
}

var Data struct {
	Info []readU.Exp
}

func indexHandler(rnd render.Render) {
	Data.Info = readU.GetStruct()

	rnd.HTML(200, "index", Data)

}
