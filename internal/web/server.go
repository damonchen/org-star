package web

import (
	"fmt"
	"github.com/damonchen/org-star/internal/config"
	"github.com/go-chi/chi"
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

func (svr *Server) Handle(w http.ResponseWriter, req *http.Request) {
	pvd := chi.URLParam(req, "provider")
	log.Debugf("svr handle, provider %s", pvd)
	if pvd == "" {
		log.Error("provider is empty")
		w.WriteHeader(500)
		return
	}

	return
}

func (svr *Server) Auth(next http.Handler) http.Handler {
	auth := svr.Cfg.Auth
	var isNoneAuth = auth.Type == "none" || auth.Type == ""
	var isBasicAuth = auth.Type == "basic"
	log.Debugf("auth type: %s\n", auth.Type)
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// 判定下是否已经认证通过了
		if isNoneAuth {
			next.ServeHTTP(w, req)
			return
		}

		if isBasicAuth {
			name, password, ok := req.BasicAuth()
			if !ok {
				log.Error("not supply correct basic auth")
				w.WriteHeader(500)
				return
			}

			if name == auth.BasicAuth.Name && password == auth.BasicAuth.Password {
				next.ServeHTTP(w, req)
			} else {
				log.Debugf("basic name %s and password %s not correct", name, password)
				w.WriteHeader(401)
			}
		}

		// err := proxyAuth(auth.ProxyAuth, w, req)
		// if err != nil {
		//	 log.Errorf("proxy auth failed: %s", err)
		//	 w.WriteHeader(401)
		// }

		next.ServeHTTP(w, req)
		return
	})
}