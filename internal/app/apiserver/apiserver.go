package apiserver

type APIServer struct {
}

func New() *APIServer {
	return &APIServer{}
}

func (api *APIServer) Start() error {
	return nil
}
