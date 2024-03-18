package external

import (
	"films/internal/usecase"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Server struct {
	log       *logrus.Logger
	mux       *http.ServeMux
	ActorsUsecase usecase.ActorsUsecase
	FilmsUsecase  usecase.FilmsUsecase
}

func NewServer(log *logrus.Logger, ActorsUsecase usecase.ActorsUsecase, FilmsUsecase usecase.FilmsUsecase) *Server {

	return &Server{
		log:                   log,
		ActorsUsecase: ActorsUsecase,
		FilmsUsecase: FilmsUsecase,
		mux: http.NewServeMux(),

	}
}

func (s *Server)Run(){
}
