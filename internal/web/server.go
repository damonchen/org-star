package web

import (
	"fmt"
	"github.com/damonchen/org-star/internal/config"
	"github.com/op/go-logging"
	"net/http"
)

//// Example format string. Everything except the message has a custom color
//// which is dependent on the log level. Many fields have a custom output
//// formatting too, eg. the time returns the hour down to the milli second.
//var format = logging.MustStringFormatter(
//	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
//)

var log = logging.MustGetLogger("web")

type Server struct {
	Cfg *config.Configuration
}

func (svr *Server) Run() error {
	//log.Infof("support web providers %s", provider.GetSupportProviders())

	r := NewRouter(svr)

	var port string
	if svr.Cfg.Host == "0.0.0.0" {
		port = fmt.Sprintf(":%s", svr.Cfg.Port)
	} else {
		port = fmt.Sprintf("%s:%s", svr.Cfg.Host, svr.Cfg.Port)
	}

	log.Infof("start listen to %s", port)
	return http.ListenAndServe(port, r)
}

func (svr *Server) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// TODO: auth

		next.ServeHTTP(w, req)
		return
	})
}
