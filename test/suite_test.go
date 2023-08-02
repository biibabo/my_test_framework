package test

import (
	"testing"
)

func TestUserManagement(t *testing.T) {
	t.Run("TestCreateUser", testCreateUser)
	t.Run("TestGetUser", testGetUser)
	t.Run("TestUpdateUser", testUpdateUser)
	t.Run("TestDeleteUser", testDeleteUser)
}
