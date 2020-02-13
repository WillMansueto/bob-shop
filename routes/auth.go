package routes

import(
	"net/http"
	"bob-shop/utils"
	"bob-shop/auth"
	"bob-shop/sessions"
	"bob-shop/models"
	"fmt"
)

func loginGetHandler(w http.ResponseWriter, r *http.Request) {
	message, alert := sessions.Flash(r, w)
	utils.ExecuteTemplate(w, "login.html", struct{
		Alert utils.Alert
	}{
		Alert: utils.NewAlert(message, alert),
	})
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	_, err := auth.Signin(email, password)
	checkErrAuthenticate(err, w, r)
}

func checkErrAuthenticate(err error, w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
	if err != nil {
		switch(err){
			case auth.ErrEmptyFields, 
				 auth.ErrEmailNotFound,
				 models.ErrInvalidEmail,
				 auth.ErrInvalidPassword:
				 session.Values["MESSAGE"] = fmt.Sprintf("%s", err)
				 session.Values["ALERT"] = "danger"
				 session.Save(r, w)
				 http.Redirect(w, r, "/login", 302)
				 return
			default:
				utils.InternalServerError(w)
				return
		}
	}
	w.Write([]byte("Logado com sucesso!"))
}