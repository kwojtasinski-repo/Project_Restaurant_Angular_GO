package test

import (
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/internal/repositories"
	"github.com/stretchr/testify/assert"
)

func (suite *IntegrationTestSuite) Test_ExistedEmailInDatabase_ExistsEmail_ShouldReturnTrue() {
	userRepository := repositories.CreateUserRepository(suite.database)

	exists, err := userRepository.ExistsByEmail("admin@admin.com")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), true, exists)
}

func (suite *IntegrationTestSuite) Test_NotExistedEmailInDatabase_ExistsEmail_ShouldReturnFalse() {
	userRepository := repositories.CreateUserRepository(suite.database)

	exists, err := userRepository.ExistsByEmail("asgsdgdsgsa@asfgsdgdsgsd.com")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), false, exists)
}
