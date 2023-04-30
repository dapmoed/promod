package app

type App struct {
	Servers []Server
	Logger  Logger
}

type Server interface {
	Serve(logger Logger) error
}

type Logger interface {
	Info(msg string, fields ...interface{})
	Error(msg string, fields ...interface{})
}

func New() *App {
	return &App{}
}

func (a *App) AddServer(s Server) {
	a.Servers = append(a.Servers, s)
}

func (a *App) Run() error {
	for _, s := range a.Servers {
		if err := s.Serve(a.Logger); err != nil {
			return err
		}
	}
	return nil
}
