package usecase

import (
	"films/internal/entity/actor"
	"films/internal/entity/global"
	"films/internal/repository"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ActorsUsecase struct {
	log       *logrus.Logger
	actorRepo repository.Actors
}

func NewActors(log *logrus.Logger, actorRepo repository.Actors) *ActorsUsecase {
	return &ActorsUsecase{
		log:       log,
		actorRepo: actorRepo,
	}
}

func (u *ActorsUsecase) CreateActor(tx sqlx.Tx, a actor.Actor) error {
	if a.Name == "" || a.Gender == "" || a.DateOfBirth.IsZero() {
		u.log.Errorln("неверные параметры")
		return global.ErrIncorrectParams
	}

	err := u.actorRepo.CreateActor(tx, a)
	if err != nil {
		u.log.Errorln("не удалось создать актера,ошибка:", err)
		return global.ErrInternalServerError
	}

	u.log.Infoln("актер успешно создан")
	return nil
}

func (u *ActorsUsecase) GetAllActors(tx sqlx.Tx) ([]actor.Actor, error) {

	actlist, err := u.actorRepo.GetAllActors(tx)
	if err != nil {
		u.log.Errorln("не удалось вывести актеров, ошибка", err)
		return nil, global.ErrInternalServerError
	}

	for index, act := range actlist {
		filmslist,err:=	u.actorRepo.GetFilmByActorsId(tx,act.ID)
		if err!= nil{
			u.log.Errorln("не удалось вывести список фильмов этого актера, ошибка", err)
			return nil,global.ErrInternalServerError
		}
		actlist[index].FilmsList = filmslist
		
	}
	return actlist, nil
}

func (u *ActorsUsecase) GetActorById(tx sqlx.Tx, actorid int)(actor.Actor, error){

	if actorid <=0 {
		u.log.Errorln("неправильный айди")
		return actor.Actor{}, global.ErrIncorrectParams
	}

	act, err := u.actorRepo.GetActorById(tx,actorid)
	if err != nil{
		u.log.Errorln("Не удалось вывести актера, ошибка", err)
		return actor.Actor{}, global.ErrInternalServerError
	}


	return act,nil}

func (u *ActorsUsecase) ChangeActorInfo(tx sqlx.Tx, actor actor.Actor) error{

	if actor.ID <=0 {
		u.log.Errorln("неправильный айди")
		return global.ErrIncorrectParams
	}

	err := u.actorRepo.ChangeActorInfo(tx,actor)
	if err != nil{
		u.log.Errorln("Не удалось изменить актера ошибка",err)
		return global.ErrInternalServerError
	}

	u.log.Infoln("Актер успешно изменен")
	return nil

}

func (u *ActorsUsecase)DeleteActorById(tx sqlx.Tx, actorid int)error{

	if actorid <=0 {
		u.log.Errorln("неправильный айди")
		return global.ErrIncorrectParams
	}

	err:= u.actorRepo.DeleteActorById(tx,actorid)
	if err!= nil {
		u.log.Errorln("Не удалось удалить актера,ошибка",err)
		return global.ErrInternalServerError
	}

	u.log.Infoln("Актер успешно удален")
	return nil

}

