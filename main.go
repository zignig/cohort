package main

import (
	"fmt"
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/zignig/viewer/util"
)

var u *universe

func main() {
	fmt.Println("Running Hub Server")
	conf := util.GetConfig("universe.toml")
	u = AndLetThereBeLight(conf)
	fmt.Println(u)
	// spin up the universe
	u.run()

	r := gin.Default()
	r.LoadHTMLFiles("index.html")
	r.GET("/static/*filepath", u.staticFiles)
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	r.GET("/ws", func(c *gin.Context) {
		u.wshandler(c.Writer, c.Request)
	})
	r.GET("/asset/*path", u.asset)
	r.Run(":8090")
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
