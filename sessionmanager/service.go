package sessionmanager

import (
	"context"
	"log"

	"github.com/google/uuid"
)

/* This is a session interface
 *
 * Basically these are the endpoints we want to expose
 * to the transport layer (HTTP, gRPC).
 * APIs are allowed to only invoke these
 * functions of the entire service.
 */
type ISessionService interface {
	CreateSession(ctx context.Context, userID string) (string, error)
	IsActiveSession(ctx context.Context, userID string) (bool, error)
	GetSessionDetail(ctx context.Context, userID string) (*Session, error)
}

/*
 * sessionservice class implements SessionService
 */
type SessionService struct {
	repository SessionRepository
}

// exported
func NewSessionService(rep SessionRepository) ISessionService {
	return &SessionService{
		repository: rep,
	}
}

func (s SessionService) CreateSession(ctx context.Context, userID string) (string, error) {
	session := Session{
		UserID: uuid.New(),
	}

	if err := s.repository.createSession(ctx, session); err != nil {
		log.Println(err)
		return "Failed" , err
	}

	return "Success", nil
}


func (s SessionService) IsActiveSession(ctx context.Context, userID string) (bool, error) {
	isActive, err := s.repository.isActiveSession(ctx, userID)
	if err != nil {
		return isActive, err 
	}
	return isActive, nil
}

func (s SessionService)	GetSessionDetail(ctx context.Context, userID string) (*Session, error) {
	// TODO : Implement get session
	return &Session{}, nil	
}
