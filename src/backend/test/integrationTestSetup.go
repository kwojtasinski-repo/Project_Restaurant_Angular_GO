package test

import (
	"database/sql"
	"log"
	"path/filepath"

	"github.com/kamasjdev/Project_Restaurant_Angular_GO/config"
	"github.com/kamasjdev/Project_Restaurant_Angular_GO/migrations"
	"github.com/stretchr/testify/suite"
)

const TestConfigFile = "config.test.yml"

type IntegrationTestSuite struct {
	suite.Suite
	config   config.Config
	database sql.DB
}

// this function executes before the test suite begins execution
func (suite *IntegrationTestSuite) SetupSuite() {
	log.Println("Running integration tests")
	log.Println("---- Setup before all tests ----")
	log.Println("Loading  config file ", TestConfigFile)
	configFile := config.LoadConfig(filepath.Join(config.GetRootPath(), TestConfigFile))
	suite.config = configFile
	log.Println("Running migrations...")
	migrations.RunMigrations(configFile, "")
	log.Println("Open connection...")
	database, err := sql.Open("mysql", configFile.DatabaseMigration.Username+":"+configFile.DatabaseMigration.Password+"@tcp(localhost:3306)/"+configFile.Database.Name+"?parseTime=true")
	if err != nil {
		log.Fatal("Cannot open database ", configFile.Database.Name)
	}
	suite.database = *database
}

// this function executes after all tests executed
func (suite *IntegrationTestSuite) TearDownSuite() {
	log.Println("---- Clean up after all tests ----")
	if _, err := suite.database.Exec("DROP DATABASE " + suite.config.Database.Name); err != nil {
		log.Fatal("ERROR: ", err)
	}
	if err := suite.database.Close(); err != nil {
		log.Fatal("ERROR: ", err)
	}
}

func (suite *IntegrationTestSuite) SetupTest() {
	log.Println("---- Setup Before Each Test ----")
}
