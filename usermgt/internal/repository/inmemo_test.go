package repository

import (
	"testing"

	s "github.com/tazafrosoul/hms-project/common/structs"
)

func TestCreate(t *testing.T) {
	user := s.User{
		FullName:       "user one",
		Username:       "user1",
		Avatar:         "https://user1.com/image/of/user/1",
		HashedPassword: "hashmeoriwillbehacked",
		Role:           "user",
	}

	repository := NewMemoRepo()
	if repository.memoryStorage[0][0] != "ID" {
		t.Errorf("FAILED, expected ID title")
	}

	result, err := repository.Create(user)

	if err != nil {
		t.Errorf("FAILED, error encountered %v", err)
	}

	if result == "" {
		t.Errorf("FAILED, expected some sort of UUID")
	} else {
		t.Logf("PASSED, returned %v", result)
	}
}
