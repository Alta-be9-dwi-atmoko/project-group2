package business

// import (
// 	"fmt"
// 	"project/group2/features/users"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// //mock data success
// type mockUserData struct{}

// func (mock mockUserData) SelectData(param string) (data []users.Core, err error) {
// 	return []users.Core{
// 		{ID: 1, Name: "alta", Email: "alta@mail.com", Password: "qwerty"},
// 	}, nil
// }

// func (mock mockUserData) InsertData(data users.Core) (row int, err error) {
// 	return 1, nil
// }

// func (mock mockUserData) LoginUserDB(authData users.AuthRequestData) (token, name string, err error) {
// 	return
// }

// func (mock mockUserData) SelectDataById(id int) (data Core, err error) {

// }

// func (mock mockUserData) UpdateDataDB(data map[string]interface{}, idUser int) (row int, err error) {

// }

// func (mock mockUserData) DeleteDataByIdDB(idUser int) (row int, err error) {

// }

// //mock data faild
// type mockUserDataFailed struct{}

// func (mock mockUserDataFailed) SelectData(param string) (data []users.Core, err error) {
// 	return nil, fmt.Errorf("Failed to select data")
// }

// func (mock mockUserDataFailed) InsertData(data users.Core) (row int, err error) {
// 	return 0, fmt.Errorf("failed to insert data")
// }

// func (mock mockUserDataFailed) LoginUserDB(authData AuthRequestData) (token, name string, err error) {

// }

// func (mock mockUserDataFailed) SelectDataById(id int) (data Core, err error) {

// }

// func (mock mockUserDataFailed) UpdateDataDB(data map[string]interface{}, idUser int) (row int, err error) {

// }

// func (mock mockUserDataFailed) DeleteDataByIdDB(idUser int) (row int, err error) {

// }

// func TestGetAllData(t *testing.T) {
// 	t.Run("Test Get All data success", func(t *testing.T) {
// 		userBusiness := NewUserBusiness(mockUserData{})
// 		result, err := userBusiness.GetAllData(0, 0)
// 		assert.Nil(t, err)
// 		assert.Equal(t, "alta", result[0].Name)
// 	})

// 	t.Run("Test Get All data failed", func(t *testing.T) {
// 		userBusiness := NewUserBusiness(mockUserDataFailed{})
// 		result, err := userBusiness.GetAllData(0, 0)
// 		assert.NotNil(t, err)
// 		assert.Nil(t, result)
// 	})
// }

// func TestCreateData(t *testing.T) {
// 	t.Run("Test Insert data success", func(t *testing.T) {
// 		userBusiness := NewUserBusiness(mockUserData{})
// 		newUser := users.Core{
// 			Name:     "alta",
// 			Email:    "alta@gmail.com",
// 			Password: "qwerty",
// 		}
// 		result, err := userBusiness.CreateData(newUser)
// 		assert.Nil(t, err)
// 		assert.Equal(t, 1, result)
// 	})

// 	t.Run("Test Insert data failed", func(t *testing.T) {
// 		userBusiness := NewUserBusiness(mockUserDataFailed{})
// 		newUser := users.Core{
// 			Name:     "alta",
// 			Email:    "alta@gmail.com",
// 			Password: "qwerty",
// 		}
// 		result, err := userBusiness.CreateData(newUser)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, 0, result)
// 	})

// 	t.Run("Test Insert data failed when  email emty", func(t *testing.T) {
// 		userBusiness := NewUserBusiness(mockUserDataFailed{})
// 		newUser := users.Core{
// 			Name:     "alta",
// 			Password: "qwerty",
// 		}
// 		result, err := userBusiness.CreateData(newUser)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, -1, result)
// 	})

// }
