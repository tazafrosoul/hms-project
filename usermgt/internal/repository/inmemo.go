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
	memoryStorage [][]string
}

func NewMemoRepo() *MemoRepo {
	title := []string{"ID", "FullName", "Username", "Email", "Avatar", "HashedPassword", "Role", "CreatedAt", "CreatedBy", "ModifiedAt", "ModifiedBy"}
	return &MemoRepo{
		memoryStorage: [][]string{title},
	}
}

// Create function stores provided user in memory and returns id or error
func (mr *MemoRepo) Create(newuser s.User) (string, error) {
	log.Println("in the repository, inmemo.go")

	newuser.ID = uuid.NewString()
	newuser.CreatedAt = time.Now()

	log.Printf("assigned id: %v\n and time: %v", newuser.ID, newuser.CreatedAt)

	data := []string{
		newuser.ID,
		newuser.FullName,
		newuser.Username,
		newuser.Email,
		newuser.Avatar,
		newuser.HashedPassword,
		newuser.Role,
		newuser.CreatedAt.String(),
		newuser.CreatedBy,
		newuser.ModifiedAt.String(),
		newuser.ModifiedBy,
	}
	mr.memoryStorage = append(mr.memoryStorage, data)
	lastIndex := len(mr.memoryStorage) - 1
	storedUser := mr.memoryStorage[lastIndex]
	if storedUser[0] != newuser.ID {
		return "", errors.New("failed to add new user")
	}

	log.Printf("Stored a User in memory database.\n currently have %v users\n User added: %v", lastIndex, storedUser)
	return storedUser[0], nil
}
