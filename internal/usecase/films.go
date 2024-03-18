package usecase

import (
	"films/internal/entity/film"
	"films/internal/entity/global"
	"films/internal/repository"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type FilmsUsecase struct {
	log      *logrus.Logger
	filmRepo repository.Films
}

func NewFilms(log *logrus.Logger, filmRepo repository.Films) *FilmsUsecase {
	return &FilmsUsecase{
		log:      log,
		filmRepo: filmRepo,
	}
}

func (u *FilmsUsecase) CreateFilm(tx sqlx.Tx, film film.Film) error {
	if film.Name == "" || film.Description == "" || film.ReleaseDate.IsZero() || film.Rating < 0 {
		u.log.Errorln("Неверные параметры")
		return global.ErrIncorrectParams
	}

	err := u.filmRepo.CreateFilm(tx, film)
	if err != nil {
		u.log.Errorln("Не удалось создать актера, ошибка", err)
		return global.ErrInternalServerError
	}
	u.log.Infoln("Фильм успешно создан")
	return nil
}

func (u *FilmsUsecase) ChangeFilmInfo(tx sqlx.Tx, film film.Film) error {

	if film.ID <= 0 {
		u.log.Errorln("неправильный айди")
		return global.ErrIncorrectParams
	}

	err := u.filmRepo.ChangeFilmInfo(tx, film)
	if err != nil {
		u.log.Errorln("Не удалось изменить фильм ошибка", err)
		return global.ErrInternalServerError
	}

	u.log.Infoln("Фильм успешно изменен")
	return nil
}

func (u *FilmsUsecase) DeleteFilmById(tx sqlx.Tx, filmid int) error {

	if filmid <= 0 {
		u.log.Errorln("неправильный айди")
		return global.ErrIncorrectParams
	}

	err := u.filmRepo.DeleteFilmById(tx, filmid)
	if err != nil {
		u.log.Errorln("Не удалось удалить фильм,ошибка", err)
		return global.ErrInternalServerError
	}

	u.log.Infoln("Фильм успешно удален")
	return nil

}

func (u *FilmsUsecase) FindFilm(tx sqlx.Tx, title, actName string) (films []film.Film,err error) {
	lf := logrus.Fields{
		"film_title_fragment" :title,
		"actor_name_fragment" :actName,
	}


	if title != "" && actName ==""{
		films,err = u.filmRepo.GetFilmByTitle(tx,title)
		switch err{
		case nil:
		case global.ErrNoData:
			u.log.WithFields(lf).Debugln("нет фильмов с таким названием")
			return
		default:
			u.log.WithFields(lf).Errorln("не удалось найти фильмы, ошибка",err)
			err = global.ErrInternalServerError
			return
		}
	}else if title == "" && actName != ""{
		films,err = u.filmRepo.GetFilmByActor(tx,actName)
		switch err{
		case nil:
		case global.ErrNoData:
			u.log.WithFields(lf).Debugln("нет фильмов c этим актером")
			return
		default:
			u.log.WithFields(lf).Errorln("не удалось найти фильмы, ошибка",err)
			err = global.ErrInternalServerError
			return
		}
	}else{
		err = global.ErrInternalServerError
		return
	}
	return
}