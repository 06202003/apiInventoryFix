package authcontroller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/06202003/apiInventory/config"

	"github.com/golang-jwt/jwt/v4"
	"github.com/06202003/apiInventory/helper"
	"gorm.io/gorm"

	"github.com/06202003/apiInventory/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {

	// mengambil inputan json
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// ambil data user berdasarkan Email
	var user models.User
	if err := models.DB.Where("Email = ?", userInput.Email).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "Email atau password salah"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"message": err.Error()}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}

	// cek apakah password valid
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		response := map[string]string{"message": "Email atau password salah"}
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	
    // proses pembuatan token jwt
    expTime := time.Now().Add(24 * time.Hour) // Set expiration to 24 hours
    claims := &config.JWTClaim{
        Email: user.Email,
        StandardClaims: jwt.StandardClaims{
            Issuer:    "go-jwt-mux",
            ExpiresAt: expTime.Unix(),
        },
    }

    // medeklarasikan algoritma yang akan digunakan untuk signing
    tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    // signed token
    token, err := tokenAlgo.SignedString(config.JWT_KEY)
    if err != nil {
        response := map[string]string{"message": err.Error()}
        helper.ResponseJSON(w, http.StatusInternalServerError, response)
        return
    }

    // set token ke cookie
    http.SetCookie(w, &http.Cookie{
        Name:     "token",
        Path:     "/",
        Value:    token,
        HttpOnly: true,
		MaxAge:   24 * 60 * 60,
    })

    response := map[string]string{"message": "login berhasil"}
    helper.ResponseJSON(w, http.StatusOK, response)
}

func Register(w http.ResponseWriter, r *http.Request) {

	// mengambil inputan json
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// hash pass menggunakan bcrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	// insert ke database
	if err := models.DB.Create(&userInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// hapus token yang ada di cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	response := map[string]string{"message": "logout berhasil"}
	helper.ResponseJSON(w, http.StatusOK, response)
}
