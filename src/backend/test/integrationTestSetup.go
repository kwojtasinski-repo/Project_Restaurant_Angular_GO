package test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
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
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/suite"
)

const TestConfigFile = "config.test.yml"

var sessionCookies map[string]*http.Cookie = make(map[string]*http.Cookie)

type ErrorResponse struct {
	Errors string `json:"errors"`
}

type IntegrationTestSuite struct {
	suite.Suite
	config     config.Config
	database   *sql.DB
	router     *gin.Engine
	users      map[string]dto.AddUserDto
	categories []dto.CategoryDto
}

// this function executes before the test suite begins execution
func (suite *IntegrationTestSuite) SetupSuite() {
	log.Println("Running integration tests")
	log.Println("---- Setup before all tests ----")
	log.Println("Loading  config file ", TestConfigFile)
	configFile := config.LoadConfig(filepath.Join(config.GetRootPath(), TestConfigFile))
	suite.config = configFile
	app.InitApp(configFile)
	log.Println("Open connection...")
	database, err := sql.Open("mysql", suite.config.DatabaseMigration.Username+":"+suite.config.DatabaseMigration.Password+"@tcp(localhost:3306)/"+suite.config.Database.Name+"?parseTime=true")
	if err != nil {
		log.Fatal("Cannot open database ", configFile.Database.Name)
	}
	suite.database = database
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
	suite.router = api.SetupApi(suite.config)
	log.Println("Running migrations...")
	migrations.UpMigrations(suite.config, "")
	suite.users = make(map[string]dto.AddUserDto)
	suite.addUsers()
	suite.addCategories()
}

// this function executes after each test case
func (suite *IntegrationTestSuite) TearDownTest() {
	log.Println("---- Setup After Each Test ----")
	migrations.DownMigrations(suite.config, "")
}

func (suite *IntegrationTestSuite) CreateAuthorizedRequest(method, url string, body io.Reader) *http.Request {
	req := suite.CreateRequest(method, url, body)
	suite.AuthorizeRequest(req)
	return req
}

func (suite *IntegrationTestSuite) CreateAuthorizedRequestForUser(method, url string, body io.Reader, user dto.AddUserDto) *http.Request {
	req := suite.CreateRequest(method, url, body)
	suite.AuthorizeRequestForUser(req, user)
	return req
}

func (suite *IntegrationTestSuite) AuthorizeRequest(request *http.Request) {
	user := suite.users["admin"]
	sessionCookieFromMap := sessionCookies[user.Email]
	if sessionCookieFromMap != nil {
		request.AddCookie(sessionCookieFromMap)
		return
	}

	req := suite.CreateRequest(http.MethodPost, "/api/sign-in", createPayload(user))
	rec := suite.SendRequest(req)
	sessionCookie := suite.FindSessionCookie(rec.Result().Cookies())
	request.AddCookie(sessionCookie)
	sessionCookies[user.Email] = sessionCookie
}

func (suite *IntegrationTestSuite) AuthorizeRequestForUser(request *http.Request, user dto.AddUserDto) {
	sessionCookieFromMap := sessionCookies[user.Email]
	if sessionCookieFromMap != nil {
		request.AddCookie(sessionCookieFromMap)
		return
	}

	req := suite.CreateRequest(http.MethodPost, "/api/sign-in", createPayload(user))
	rec := suite.SendRequest(req)
	sessionCookie := suite.FindSessionCookie(rec.Result().Cookies())
	request.AddCookie(sessionCookie)
	sessionCookies[user.Email] = sessionCookie
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

func (suite *IntegrationTestSuite) addUsers() {
	passwordHasher := services.CreatePassworHasherService()
	standardUser := dto.AddUserDto{
		Email:    "test@test.com",
		Password: "test123",
	}
	adminUser := dto.AddUserDto{
		Email:    "admin@admin-test.com",
		Password: "test123",
	}
	testUser := dto.AddUserDto{
		Email:    "test-user@test-abc123.com",
		Password: "test123",
	}
	revokeSessionUser := dto.AddUserDto{
		Email:    "test-revoke-session-user@test-abc123.com",
		Password: "test123",
	}
	suite.users["user"] = standardUser
	suite.users["admin"] = adminUser
	suite.users["test"] = testUser
	suite.users["revokeSessionUser"] = revokeSessionUser
	userRepository := repositories.CreateUserRepository(suite.database)
	commonUserPassword, err := passwordHasher.HashPassword(standardUser.Password)
	if err != nil {
		log.Fatal("Error while creating user", err.Error())
	}
	standardUserEmail, _ := valueobjects.NewEmailAddress(standardUser.Email)
	adminUserEmail, _ := valueobjects.NewEmailAddress(adminUser.Email)
	testUserEmail, _ := valueobjects.NewEmailAddress(testUser.Email)
	revokeSessionUserEmail, _ := valueobjects.NewEmailAddress(revokeSessionUser.Email)

	userRepository.Add(&entities.User{
		Email:    *standardUserEmail,
		Password: commonUserPassword,
		Role:     "user",
		Deleted:  false,
	})
	userRepository.Add(&entities.User{
		Email:    *adminUserEmail,
		Password: commonUserPassword,
		Role:     "admin",
		Deleted:  false,
	})
	userRepository.Add(&entities.User{
		Email:    *testUserEmail,
		Password: commonUserPassword,
		Role:     "test",
		Deleted:  false,
	})
	userRepository.Add(&entities.User{
		Email:    *revokeSessionUserEmail,
		Password: commonUserPassword,
		Role:     "admin",
		Deleted:  false,
	})
}

func (suite *IntegrationTestSuite) addCategories() {
	suite.categories = make([]dto.CategoryDto, 0)
	suite.categories = append(suite.categories, suite.AddCategory())
	suite.categories = append(suite.categories, suite.AddCategory())
	suite.categories = append(suite.categories, suite.AddCategory())
	suite.categories = append(suite.categories, suite.AddCategory())
	suite.categories = append(suite.categories, suite.AddCategory())
}

func (suite *IntegrationTestSuite) AddCategory() dto.CategoryDto {
	value := rand.Intn(100000) + 1
	category := &dto.CategoryDto{
		Name: fmt.Sprint("category#", value),
	}
	categoryService, err := api.CreateCategoryService()
	suite.Require().Nil(err)
	category, errAdd := categoryService.Add(category)
	suite.Require().Nil(errAdd)
	return *category
}

func (suite *IntegrationTestSuite) AddProduct() dto.ProductDetailsDto {
	category := suite.AddCategory()
	value := rand.Intn(100000) + 1
	product := &dto.AddProductDto{
		Name:        "Product#1",
		Description: "Description#123456789",
		CategoryId:  category.Id,
		Price:       decimal.New(int64(value), 1),
	}
	productService, err := api.CreateProductService()
	suite.Require().Nil(err)
	productAdded, errAdd := productService.Add(product)
	suite.Require().Nil(errAdd)
	return *productAdded
}

func (suite *IntegrationTestSuite) AddProductToCart() {
	product := suite.AddProduct()
	addCart := dto.AddCart{
		ProductId: product.Id,
	}
	req := suite.CreateAuthorizedRequest(http.MethodPost, "/api/carts", createPayload(addCart))
	rec := suite.SendRequest(req)
	suite.Require().Equal(http.StatusCreated, rec.Result().StatusCode)
}

func (suite *IntegrationTestSuite) AddProductToCartForUser(user dto.AddUserDto) {
	product := suite.AddProduct()
	addCart := dto.AddCart{
		ProductId: product.Id,
	}
	req := suite.CreateAuthorizedRequestForUser(http.MethodPost, "/api/carts", createPayload(addCart), user)
	rec := suite.SendRequest(req)
	suite.Require().Equal(http.StatusCreated, rec.Result().StatusCode)
}

func (suite *IntegrationTestSuite) AddProductWithIdToCart(productId dto.IdObject) {
	addCart := dto.AddCart{
		ProductId: productId,
	}
	req := suite.CreateAuthorizedRequest(http.MethodPost, "/api/carts", createPayload(addCart))
	rec := suite.SendRequest(req)
	suite.Require().Equal(http.StatusCreated, rec.Result().StatusCode)
}

func (suite *IntegrationTestSuite) AddProductWithIdToCartForUser(productId dto.IdObject, user dto.AddUserDto) {
	addCart := dto.AddCart{
		ProductId: productId,
	}
	req := suite.CreateAuthorizedRequestForUser(http.MethodPost, "/api/carts", createPayload(addCart), user)
	rec := suite.SendRequest(req)
	suite.Require().Equal(http.StatusCreated, rec.Result().StatusCode)
}

func (suite *IntegrationTestSuite) getErrorResponse(rec *httptest.ResponseRecorder) ErrorResponse {
	var errorResponse ErrorResponse
	suite.Require().Nil(json.Unmarshal(rec.Body.Bytes(), &errorResponse))
	return errorResponse
}
