package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/op/go-logging"
	"github.com/zignig/viewer/util"
)

var u *universe

var log = logging.MustGetLogger("universe")

func main() {
	fmt.Println("Running Hub Server")
	conf := util.GetConfig("universe.toml")
	u = AndLetThereBeLight(conf)
	fmt.Println(u)
	// spin up the universe
	u.run()
	// set up the templates
	r := gin.Default()
	u.LoadTemplates()
	r.SetHTMLTemplate(u.templ)

	r.GET("/static/*filepath", u.staticFiles)
	r.GET("/", u.index)
	r.GET("/ws", func(c *gin.Context) {
		u.wshandler(c.Writer, c.Request)
	})
	r.GET("/asset/*path", u.asset)
	r.Run(":8090")
}

func (u *universe) index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func (u *universe) asset(c *gin.Context) {
	// send to asset manager
	path := c.Params.ByName("path")
	data, err := u.cache.Cat(path)
	if err != nil {
		c.String(500, err.Error())
	}
	c.Data(200, "", data)
}

func (u *universe) staticFiles(c *gin.Context) {
	static, err := rice.FindBox("static")
	if err != nil {
		fmt.Println("Static Error")
	}
	original := c.Request.URL.Path
	c.Request.URL.Path = c.Params.ByName("filepath")
	http.FileServer(static.HTTPBox()).ServeHTTP(c.Writer, c.Request)
	c.Request.URL.Path = original
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (u *universe) wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}
	c := NewConnection(u.h, conn)
	u.h.register <- c
	//world.register <- c
	// todo  , move this to write pump and push a new player
	go c.writePump()
	c.readPump()
}

func (u *universe) LoadTemplates() {
	templateBox, err := rice.FindBox("templates")
	if err != nil {
		log.Critical("template fail ", err)
	}
	// collect all the templates
	fileList := []string{}
	visit := func(path string, f os.FileInfo, inerr error) (err error) {
		fmt.Println(path)
		if f.IsDir() == false {
			fileList = append(fileList, path)
		}
		return nil
	}
	walkerr := templateBox.Walk("", visit)
	if walkerr != nil {
		log.Critical("walk %s", err)
	}
	templates := template.New("")
	fmt.Println(fileList)
	for _, x := range fileList {
		templateString, err := templateBox.String(x)
		if err != nil {
			log.Fatal(err)
		}
		_, err = templates.New(x).Parse(templateString)
		if err != nil {
			log.Fatal(err)
		}
	}
	u.templ = templates
}
