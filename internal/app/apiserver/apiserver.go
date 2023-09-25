package apiserver

import (
	"github.com/YerkanatSultanov/http-rest-api/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (api *APIServer) Start() error {
	if err := api.configLogger(); err != nil {
		return err
	}

	api.configRouter()

	if err := api.configStore(); err != nil {
		return err
	}

	api.logger.Info("Starting api server")
	return http.ListenAndServe(api.config.BindAddr, api.router)
}

func (api *APIServer) configLogger() error {
	level, err := logrus.ParseLevel(api.config.LogLevel)

	if err != nil {
		return err
	}

	api.logger.SetLevel(level)

	return nil
}

func (api *APIServer) configRouter() {
	api.router.HandleFunc("/hello", api.handleHello())
}

func (api *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := io.WriteString(w, "Hello")
		if err != nil {
			return
		}
	}
}

func (api *APIServer) configStore() error {
	st := store.New(api.config.Store)

	if err := st.Open(); err != nil {
		return err
	}

	api.store = st

	return nil

}
