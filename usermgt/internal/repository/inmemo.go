package repository

//this is the test database
import (
	"errors"
	s "hms-project/common/structs"
	"log"
	"time"

	"github.com/google/uuid"
)

type MemoRepo struct {
	memoryStorage []s.User
}

func NewMemoRepo() *MemoRepo {
	return &MemoRepo{
		memoryStorage: []s.User{},
	}
}

// Create function stores provided user in memory and returns id or error
func (mr *MemoRepo) Create(newuser s.User) (string, error) {
	newuser.ID = uuid.NewString()
	newuser.CreatedAt = time.Now()

	mr.memoryStorage = append(mr.memoryStorage, newuser)
	lastIndex := len(mr.memoryStorage) - 1
	storedUser := mr.memoryStorage[lastIndex]
	if storedUser.ID != newuser.ID {
		return "", errors.New("failed to add new user")
	}

	log.Printf("Stored a User in memory database.\n currently have %v users\n User added: %v", lastIndex+1, storedUser)
	return storedUser.ID, nil
}
