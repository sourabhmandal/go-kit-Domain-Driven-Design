package sessionmanager

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

/*
 * Interface that defines the methods
 * which will communicate with database
 * i.e
 * methods that perform Database operations on sessions table
 */
type SessionRepository interface {
	createSession(ctx context.Context, session Session) error
	isActiveSession(ctx context.Context, userID string) (bool, error)
}

/*
 * class sessionRepository implements SessionRepository
 */
type sessionrepository struct {
	// local variables
	// db helps make database operations
	db *gorm.DB
}

func NewRepo(db *gorm.DB) sessionrepository {
	return sessionrepository{
		db: db,
	}
}

func (r sessionrepository) createSession(ctx context.Context, session Session) error {
	if session.UserID.String() == "" {
		return errors.New("User ID field in sessions cannot be empty")
	}
	fmt.Println("----SAVING SESSION TO DB----")
	// TODO: Write Database Statement
	return nil
}

func (r sessionrepository) isActiveSession(ctx context.Context, userID string) (bool, error) {
	if userID == "" {
		return false, errors.New("User ID field in sessions cannot be empty")
	}
	fmt.Printf("----QUERY SESSION %s FROM DB----", userID)
	// TODO: Write Database Statement
	return true, nil
}
