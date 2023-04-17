package repos

import (
	"context"
	"time"

	"github.com/ganiszulfa/concise/backend/internal/models"
	"github.com/ganiszulfa/concise/backend/pkg/generate"
	"github.com/ganiszulfa/concise/backend/pkg/trace"
	"gorm.io/gorm"
)

var DayExpired = 14
var SessionLen = 60

type SessionRepoInterface interface {
	GetById(ctx context.Context, id string) (session models.Session, err error)
	Create(ctx context.Context) (sessionId string, err error)
}

type SessionRepo struct {
	db *gorm.DB
}

func NewSessionRepo(db *gorm.DB) SessionRepoInterface {
	return &SessionRepo{db: db}
}

func (r SessionRepo) GetById(ctx context.Context, id string) (session models.Session, err error) {
	trace.Func()

	err = r.db.WithContext(ctx).First(&session, "id = ?", id).Error

	return
}

func (r SessionRepo) Create(ctx context.Context) (sessionId string, err error) {
	trace.Func()

	sessionId = generate.RandAlphabetsLowerCase(SessionLen)

	session := &models.Session{
		Id:        sessionId,
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().AddDate(0, 0, DayExpired),
	}

	err = r.db.WithContext(ctx).Create(session).Error
	return
}
