package repository

//this is the test database. this will only be used for a development only. will not be included in production
import (
	"errors"
	"log"
	"time"

	s "github.com/tazafrosoul/hms-project/common/structs"
	"github.com/tazafrosoul/hms-project/common/utility"

	"github.com/google/uuid"
)

type MemoRepo struct {
	memoryStorage [][]string
}

func NewMemoRepo() *MemoRepo {
	title := []string{"ID", "FullName", "Username", "Avatar", "HashedPassword", "Role", "CreatedAt", "CreatedBy", "ModifiedAt", "ModifiedBy"}

	return &MemoRepo{
		memoryStorage: [][]string{title},
	}
}

func (mr *MemoRepo) Init() error {
	//NOTE: This shit is UNSAFE.
	//TODO remember to remove it from production code
	hashedpw, err := utility.Hash("p@ssw0rd123_321")
	if err != nil {
		//TODO add log service
		log.Fatalf("could not hash the provided password: %v\n", err)
		return err
	}

	mr.Create(s.User{
		ID:             uuid.NewString(),
		FullName:       "Super User",
		Username:       "Superuser",
		Avatar:         "http://nil.nil",
		HashedPassword: hashedpw,
		Role:           "superuser",
		CreatedAt:      time.Now(),
		CreatedBy:      "",
		ModifiedAt:     time.Time{},
		ModifiedBy:     "",
	})
	return nil
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
