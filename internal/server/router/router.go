package router

import (
	"dnd-game/internal/server"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

type Router struct {
	*http.ServeMux
	*server.Server
	prefix   string
	logger   *log.Logger
	upgrader websocket.Upgrader
}

func (router *Router) Log(msg string, args ...interface{}) {
	router.logger.Printf(msg, args...)
}

func NewRouter(server *server.Server, prefix string, logger *log.Logger) *Router {
	return &Router{
		prefix:   prefix,
		Server:   server,
		ServeMux: http.NewServeMux(),
		logger:   logger,
	}
}

func (router *Router) Route(pattern string, handler func(w http.ResponseWriter, r *http.Request) error) {
	router.HandleFunc(router.addPrefixToPattern(pattern), func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		router.HandleHTTPError(err, w)
	})
}

func (router *Router) Websocket(pattern string, handler func(conn *websocket.Conn, w http.ResponseWriter, r *http.Request) error) {
	router.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		conn, err := router.upgrader.Upgrade(w, r, nil)
		if err != nil {
			router.logger.Println("Failed to upgrade connection")
			router.HandleHTTPError(err, w)
			return
		}
		err = handler(conn, w, r)
		router.HandleHTTPError(err, w)
		defer conn.Close()
	})
}

func (router *Router) addPrefixToPattern(pattern string) (path string) {
	split := strings.Split(pattern, " ")
	if len(split) == 2 {
		method := split[0] + " "
		pattern = split[1]
		path = method + router.prefix + pattern
	} else if len(split) == 1 {
		path = router.prefix + pattern
	} else {
		panic("invalid pattern " + pattern)
	}
	return
}
