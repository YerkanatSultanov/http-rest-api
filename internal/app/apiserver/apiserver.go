package apiserver

type APIServer struct {
	config *Config
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
	}
}

func (api *APIServer) Start() error {
	return nil
}
