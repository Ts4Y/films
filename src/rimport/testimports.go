package rimport

import (
	"log"
	"os"
	"vk-film-library/config"
	"vk-film-library/internal/repository"
	"vk-film-library/internal/transaction"

	"go.uber.org/mock/gomock"
)

type TestRepositoryImports struct {
	Config         config.Config
	SessionManager *transaction.MockSessionManager
	MockRepository MockRepository
	ctrl           *gomock.Controller
}

func NewTestRepositoryImports(
	ctrl *gomock.Controller,
) TestRepositoryImports {
	config, err := config.NewConfig(os.Getenv("CONF_PATH"))
	if err != nil {
		log.Fatalln(err)
	}

	return TestRepositoryImports{
		ctrl:           ctrl,
		Config:         config,
		SessionManager: transaction.NewMockSessionManager(ctrl),
		MockRepository: MockRepository{
			Actor: repository.NewMockActor(ctrl),
			Movie: repository.NewMockMovie(ctrl),
		},
	}
}

func (t *TestRepositoryImports) MockSession() *transaction.MockSession {
	ts := transaction.NewMockSession(t.ctrl)

	ts.EXPECT().Start().Return(nil).AnyTimes()
	ts.EXPECT().Rollback().Return(nil).AnyTimes()

	return ts
}

func (t *TestRepositoryImports) MockSessionWithCommit() *transaction.MockSession {
	ts := t.MockSession()

	ts.EXPECT().Commit().Return(nil).AnyTimes()

	return ts
}

func (t *TestRepositoryImports) RepositoryImports() RepositoryImports {
	return RepositoryImports{
		SessionManager: t.SessionManager,
		Config:         t.Config,
		Repository: Repository{
			Actor: t.MockRepository.Actor,
			Movie: t.MockRepository.Movie,
		},
	}
}
