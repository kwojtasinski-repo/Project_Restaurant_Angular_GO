package test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/config"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/api"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/app"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/dto"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities"
	valueobjects "github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/entities/value-objects"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/repositories"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/services"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/internal/settings"
	"github.com/kwojtasinski-repo/Project_Restaurant_Angular_GO/migrations"
	"github.com/stretchr/testify/suite"
)

const TestConfigFile = "config.test.yml"

var sessionCookie *http.Cookie

type IntegrationTestSuite struct {
	suite.Suite
	config   config.Config
	database *sql.DB
	router   *gin.Engine
	users    map[string]dto.AddUserDto
}

// this function executes before the test suite begins execution
func (suite *IntegrationTestSuite) SetupSuite() {
	log.Println("Running integration tests")
	log.Println("---- Setup before all tests ----")
	log.Println("Loading  config file ", TestConfigFile)
	configFile := config.LoadConfig(filepath.Join(config.GetRootPath(), TestConfigFile))
	suite.config = configFile
	app.InitApp(configFile)
	log.Println("Running migrations...")
	migrations.UpMigrations(configFile, "")
	log.Println("Open connection...")
	database, err := sql.Open("mysql", suite.config.DatabaseMigration.Username+":"+suite.config.DatabaseMigration.Password+"@tcp(localhost:3306)/"+suite.config.Database.Name+"?parseTime=true")
	if err != nil {
		log.Fatal("Cannot open database ", configFile.Database.Name)
	}
	suite.database = database
	suite.users = make(map[string]dto.AddUserDto)
	suite.createUsers()
	suite.router = api.SetupApi(suite.config)
}

// this function executes after all tests executed
func (suite *IntegrationTestSuite) TearDownSuite() {
	log.Println("---- Clean up after all tests ----")
	log.Println("Dropping user ", suite.config.Database.Username)
	if _, err := suite.database.Exec("DROP USER " + suite.config.Database.Username); err != nil {
		log.Fatal("ERROR: ", err)
	}

	log.Println("Dropping database ", suite.config.Database.Name)
	if _, err := suite.database.Exec("DROP DATABASE " + suite.config.Database.Name); err != nil {
		log.Fatal("ERROR: ", err)
	}
	if err := suite.database.Close(); err != nil {
		log.Fatal("ERROR: ", err)
	}
}

func (suite *IntegrationTestSuite) SetupTest() {
	log.Println("---- Setup Before Each Test ----")
	// add test data?
}

// this function executes after each test case
func (suite *IntegrationTestSuite) TearDownTest() {
	log.Println("---- Setup After Each Test ----")
}

func (suite *IntegrationTestSuite) CreateAuthorizedRequest(method, url string, body io.Reader) *http.Request {
	req := suite.CreateRequest(method, url, body)
	suite.AuthorizeRequest(req)
	return req
}

func (suite *IntegrationTestSuite) AuthorizeRequest(request *http.Request) {
	if sessionCookie != nil {
		request.AddCookie(sessionCookie)
		return
	}

	user := suite.users["admin"]
	req := suite.CreateRequest("POST", "/api/sign-in", createPayload(user))
	rec := suite.SendRequest(req)
	sessionCookie := suite.FindSessionCookie(rec.Result().Cookies())
	request.AddCookie(sessionCookie)
}

func (suite *IntegrationTestSuite) FindSessionCookie(cookies []*http.Cookie) *http.Cookie {
	suite.Require().NotEmpty(cookies)
	for _, cookie := range cookies {
		if cookie.Name == settings.CookieSessionName {
			return cookie
		}
	}

	return nil
}

func (suite *IntegrationTestSuite) CreateRequest(method, url string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, body)
	suite.Require().NoError(err)
	return req
}

func (suite *IntegrationTestSuite) SendRequest(request *http.Request) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	suite.router.ServeHTTP(rec, request)
	return rec
}

func createPayload(value interface{}) *bytes.Reader {
	data, err := json.Marshal(value)
	if err != nil {
		log.Fatal(err)
	}
	return bytes.NewReader(data)
}

func (suite *IntegrationTestSuite) createUsers() {
	passwordHasher := services.CreatePassworHasherService()
	standardUser := dto.AddUserDto{
		Email:    "test@test.com",
		Password: "test123",
	}
	adminUser := dto.AddUserDto{
		Email:    "admin@admin-test.com",
		Password: "test123",
	}
	suite.users["user"] = standardUser
	suite.users["admin"] = adminUser
	userRepository := repositories.CreateUserRepository(suite.database)
	standarUserPassword, err := passwordHasher.HashPassword(standardUser.Password)
	if err != nil {
		log.Fatal("Error while creating user", err.Error())
	}
	adminUserPassword, err := passwordHasher.HashPassword(standardUser.Password)
	if err != nil {
		log.Fatal("Error while creating user", err.Error())
	}
	standardUserEmail, _ := valueobjects.NewEmailAddress(standardUser.Email)
	adminUserEmail, _ := valueobjects.NewEmailAddress(adminUser.Email)

	userRepository.Add(&entities.User{
		Email:    *standardUserEmail,
		Password: standarUserPassword,
		Role:     "user",
		Deleted:  false,
	})
	userRepository.Add(&entities.User{
		Email:    *adminUserEmail,
		Password: adminUserPassword,
		Role:     "admin",
		Deleted:  false,
	})
}
