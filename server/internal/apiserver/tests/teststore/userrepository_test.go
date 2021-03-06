package teststore_test

import (
	"testing"

	ModelUser "github.com/Gugush284/Go-server.git/internal/apiserver/model/user"
	"github.com/Gugush284/Go-server.git/internal/apiserver/store/teststore"
	ModelUserTest "github.com/Gugush284/Go-server.git/internal/apiserver/tests/model/users"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	u, err := s.User().Create(ModelUserTest.TestUser(t))
	/*if u != nil {
		fmt.Println(u.ID, u.Login, u.DecryptedPassword, u.Password)
	}*/

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByLogin(t *testing.T) {
	s := teststore.New()
	u := ModelUserTest.TestUser(t)
	/*if u != nil {
		fmt.Println("1", u.ID, u.Login, u.DecryptedPassword, u.Password)
	}*/

	user, err := s.User().FindByLogin(u.Login)
	/*if user != nil {
		fmt.Println("2", user.ID, user.Login, user.DecryptedPassword, user.Password)
	}*/
	assert.Error(t, err)
	assert.Nil(t, user)

	user, err = s.User().Create(&ModelUser.User{
		Login:             u.Login,
		DecryptedPassword: u.DecryptedPassword,
	})
	/*if user != nil {
		fmt.Println("3", user.ID, user.Login, user.DecryptedPassword, user.Password)
	}*/
	assert.NoError(t, err)
	assert.NotNil(t, user)

	user, err = s.User().FindByLogin(u.Login)
	/*if user != nil {
		fmt.Println("4", user.ID, user.Login, user.DecryptedPassword, user.Password)
	}*/
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserRepository_Find(t *testing.T) {
	s := teststore.New()
	u := ModelUserTest.TestUser(t)
	/*if u != nil {
		fmt.Println("1)", u.ID, u.Login, u.DecryptedPassword, u.Password)
	}*/

	u, err := s.User().Create(&ModelUser.User{
		Login:             u.Login,
		DecryptedPassword: u.DecryptedPassword,
	})
	/*if u != nil {
		fmt.Println("2)", u.ID, u.Login, u.DecryptedPassword, u.Password, err)
	}*/
	assert.NoError(t, err)
	assert.NotNil(t, u)

	user, err := s.User().Find(u.ID)
	/*if user != nil {
		fmt.Println("3)", user.ID, user.Login, user.DecryptedPassword, user.Password, err)
	}*/
	assert.NoError(t, err)
	assert.NotNil(t, user)
}
