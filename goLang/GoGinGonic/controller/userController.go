package controller

import (
	"GoGinGonic/database"
	"GoGinGonic/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	// Get email and password from the  body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to load body",
		})

		return
	}

	//Hash the Password

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to hash",
		})
		return
	}

	// Create User
	user := models.User{Email: body.Email, Password: string(hash)}

	result := database.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create user db",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
	// Get a body of req containing emalid and pass
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to laod Body",
		})
		return
	}

	// verify the hash
	var user models.User
	database.DB.First(&user, "email=?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to find user",
		})
		return
	}

	hashed := user.Password

	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "incorrect Password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create token",
		})
		return

	}

	// set Cookies
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization",tokenString,3600*24*30,"","",false,true)


	c.JSON(http.StatusOK, gin.H{})
} 

func Validate(c * gin.Context) {

	user,_ := c.Get("user") 

	c.JSON(http.StatusOK,gin.H{
		"message":user,
	})
}