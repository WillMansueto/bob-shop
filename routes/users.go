package routes

import (
	"fmt"
	"net/http"
	"bob-shop/utils"
	"bob-shop/models"
	"bob-shop/sessions"
)

func registerGetHandler(w http.ResponseWriter, r *http.Request) {
	message, alert := sessions.Flash(r, w)
	utils.ExecuteTemplate(w, "register.html", struct{
		Alert utils.Alert
	}{
		Alert: utils.NewAlert(message, alert),
	})
}

func registerPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var user models.User
	user.FirstName = r.PostForm.Get("firstname")
	user.LastName = r.PostForm.Get("lastname")
	user.Email = r.PostForm.Get("email")
	user.Password = r.PostForm.Get("password")
	_, err := models.NewUser(user)
	checkErrRegister(err, w, r)
}

func checkErrRegister(err error, w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
	message := "Cadastrado com sucesso!"
	if err != nil {
		switch(err) {
		case models.ErrDuplicatedKeyEmail,
			 models.ErrRequiredFirstName,
			 models.ErrRequiredLastName,
			 models.ErrRequiredEmail,
			 models.ErrInvalidEmail,
			 models.ErrRequiredPassword,
			 models.ErrMaxLimit:
			 message = fmt.Sprintf("%s", err)
			 break;
		default:
			utils.InternalServerError(w)
			return
		}
		session.Values["MESSAGE"] = message
		session.Values["ALERT"] = "danger"
		session.Save(r, w)
		http.Redirect(w, r, "/register", 302)
		return
	}
	session.Values["MESSAGE"] = message
	session.Values["ALERT"] = "success"
	session.Save(r, w)
	http.Redirect(w, r, "/login", 302)
}