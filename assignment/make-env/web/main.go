package main

import (
	"fmt"
	"net/http"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
)

var (
	manager LiveManager
)

type LiveInfo struct {
	app  string
	name string
}

func (i LiveInfo) key() string {
	return fmt.Sprintf("app=%s&name=%s", i.app, i.name)
}

func (i LiveInfo) URL() string {
	return fmt.Sprintf("http://localhost:8081/hls/%s.m3u8", i.name)
}

type LiveManager struct {
	publish chan *LiveInfo
	done    chan *LiveInfo
	lives   map[string]*LiveInfo
}

func (m *LiveManager) run() {
	for {
		select {
		case info := <-m.publish:
			manager.lives[info.key()] = info
		case info := <-m.done:
			if _, ok := manager.lives[info.key()]; ok {
				delete(manager.lives, info.key())
			}
		}
	}
}

var templateIndex = pongo2.Must(pongo2.FromFile("template/index.html"))

func main() {
	manager = LiveManager{
		publish: make(chan *LiveInfo, 64),
		done:    make(chan *LiveInfo, 64),
		lives:   map[string]*LiveInfo{},
	}
	go manager.run()
	e := echo.New()
	e.GET("/", handlerIndex)
	e.Static("/public", "public")
	e.GET("/chat/:room", handleChatroom)
	e.GET("/on_publish", handlerOnPublish)
	e.GET("/on_publish_done", handlerOnPublishDone)
	e.Logger.Fatal(e.Start(":8080"))

}

func handleChatroom(c echo.Context) error {
	room := echo.Context.Param(c, "room")
	chat := pongo2.Must(pongo2.FromFile("template/chat.html"))
	body, err := chat.Execute(
		pongo2.Context{"room": room},
	)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	return c.HTML(http.StatusOK, body)
}

func handlerIndex(c echo.Context) error {
	lives := make([]*LiveInfo, 0, len(manager.lives))
	for _, l := range manager.lives {
		lives = append(lives, l)
	}
	body, err := templateIndex.Execute(
		pongo2.Context{
			"lives": lives,
		},
	)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	return c.HTML(http.StatusOK, body)
}

func handlerOnPublish(c echo.Context) error {
	manager.publish <- GetLiveInfoFromContext(c)
	return c.String(http.StatusOK, "")
}

func handlerOnPublishDone(c echo.Context) error {
	manager.done <- GetLiveInfoFromContext(c)
	return c.String(http.StatusOK, "")
}

func GetLiveInfoFromContext(c echo.Context) *LiveInfo {
	app := c.QueryParam("app")
	name := c.QueryParam("name")
	return &LiveInfo{
		app:  app,
		name: name,
	}
}
