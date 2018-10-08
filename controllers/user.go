package controllers

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jorgeAM/comments/commons"
	"github.com/jorgeAM/comments/config"
	"github.com/jorgeAM/comments/models"
)

//Login inicio de sesión
func Login(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	//mapear los datos de request al modelo user
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err)
		return
	}

	db := config.GetConnection()
	defer db.Close()

	c := sha256.Sum256([]byte(user.Password))
	pwd := base64.URLEncoding.EncodeToString(c[:32])

	db.Where("email = ? AND password = ?", user.Email, pwd).First(&user)
	if user.ID > 0 {
		user.Password = ""
		token := commons.GenerateJWT(user)
		j, err := json.Marshal(models.Token{Token: token})
		if err != nil {
			log.Fatalf("Error al convertir el token a JSON: %s", err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	} else {
		m := models.Message{
			Message: "Usuario o clave no válido",
			Code:    http.StatusUnauthorized,
		}

		commons.DisplayMessage(w, m)
	}
}

//UserCreate registro de usuario
func UserCreate(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	m := models.Message{}

	//mapear los datos de request al modelo user
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		m.Message = fmt.Sprintf("Error al leer el usuario a registrar: %s", err)
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}

	if user.Password != user.ConfirmPassword {
		m.Message = fmt.Sprintf("Las contraseñas no coinciden")
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}

	c := sha256.Sum256([]byte(user.Password))
	//forma más limpia (RECOMENDABLE)
	pwd := fmt.Sprintf("%x", c)
	user.Password = pwd

	picmd5 := md5.Sum([]byte(user.Email))
	picstr := fmt.Sprintf("%x", picmd5)
	user.Avatar = "https://gravatar.com/avatar/" + picstr + "?s=100"

	db := config.GetConnection()
	defer db.Close()

	err = db.Create(&user).Error
	if err != nil {
		m.Message = fmt.Sprintf("Error al crear el registro: %s", err)
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}

	m.Message = fmt.Sprintf("Usuario creado con exito")
	m.Code = http.StatusCreated
	commons.DisplayMessage(w, m)
}
