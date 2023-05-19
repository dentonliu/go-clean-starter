package util

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/dentonliu/go-clean-starter/internal/config"
)

type TestCase struct {
	Method           string
	Url              string
	Body             string
	ExpectedStatus   int
	ExpectedResponse string
}

func RunApiTests(t *testing.T, r *gin.Engine, testCases []TestCase) {
	for _, tc := range testCases {
		w := httptest.NewRecorder()

		req, _ := http.NewRequest(tc.Method, tc.Url, bytes.NewBufferString(tc.Body))
		req.Header.Add("Content-Type", "application/json")

		r.ServeHTTP(w, req)
		assert.Equal(t, tc.ExpectedStatus, w.Code)

		if tc.ExpectedResponse != "" {
			assert.Contains(t, w.Body.String(), tc.ExpectedResponse)
		}
	}
}

func RunWebTests(t *testing.T, r *gin.Engine, testCases []TestCase) {
	w := httptest.NewRecorder()

	for _, tc := range testCases {
		req, _ := http.NewRequest(tc.Method, tc.Url, bytes.NewBufferString(tc.Body))
		req.Header.Add("Content-Type", "appliation/www-form-urlencode")

		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)

		if tc.ExpectedResponse != "" {
			assert.Contains(t, w.Body.String(), tc.ExpectedResponse)
		}
	}
}

func MockApiRouter() *gin.Engine {
	return gin.Default()
}

func MockWebRouter() *gin.Engine {
	r := gin.Default()

	r.HTMLRender = CreateRenderer("../../web/template")

	return r
}

func MockJWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := User{
			ID: "1",
			IP: "127.0.0.1",
		}

		ctx.Set(CKUser, &user)
	}
}

func MockDB() *gorm.DB {
	c, err := config.Load("../../configs")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(mysql.Open(c.DSNTest), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
