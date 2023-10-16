package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"blue.study.golang/internal/models"
	"blue.study.golang/internal/validator"

	"github.com/julienschmidt/httprouter"
)


type userSignUpForm struct {
	Name				string	`form:"name"` 
	Email				string	`form:"email"`
	Password			string	`form:"password"`
	validator.Validator 		`form:"-"`
}

type userLoginForm struct {
	Email 				string `form:"email"`
	Password 			string `form:"password"`
	validator.Validator 	   `form:"-"`
}

type snippetCreateForm struct {
	Title 				 string 		`form:"title"`
	Content 			 string 		`form:"content"`
	Expires 			 int			`form:"expires"`
	validator.Validator	 				`form:"-"`
}


func (app *application) home(w http.ResponseWriter, r *http.Request){

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)

	data.Snippets = snippets

	app.render(w, "home.tmpl.html", http.StatusOK, data)
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request){

	params:= httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord){
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Snippet = snippet

	app.render(w, "view.tmpl.html", http.StatusOK, data)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request){
	data := app.newTemplateData(r)

	data.Form = snippetCreateForm{ 
		Expires: 365,
	}

	app.render(w, "create.tmpl.html", http.StatusOK, data)
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request){
	
	
	var snippetForm snippetCreateForm
	
	err := app.decodePostForm(r, &snippetForm)
	if err != nil {
		app.clientError(w, http.StatusBadRequest, nil)
		return
	}
	
	// checking fields and validating form
	snippetForm.CheckFiled(validator.IsNotBlank(snippetForm.Title), "title", "title can't be blank")
	snippetForm.CheckFiled(validator.IsMaxChars(snippetForm.Title, 100), "title", "title can't be more than 100 charaters")
	snippetForm.CheckFiled(validator.IsNotBlank(snippetForm.Content), "content", "content can't be blank")
	snippetForm.CheckFiled(validator.IsPermittedValue(snippetForm.Expires, 1, 7, 365), "expires", "This field must be equal to 1, 7 or 365")

	if !snippetForm.Validator.Valid() {
		data := app.newTemplateData(r)
		data.Form = snippetForm
		app.render(w, "create.tmpl.html", http.StatusUnprocessableEntity, data)
		return
	}

	id, err := app.snippets.Insert(snippetForm.Title, snippetForm.Content, snippetForm.Expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.sessionManager.Put(r.Context(), "flash", "Snippet Sucessfully Created.")

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}

func (app *application) userSignUp(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = &userSignUpForm{}
	app.render(w, "signup.tmpl.html", http.StatusOK, data)
}

func (app *application) userSignUpPost(w http.ResponseWriter, r *http.Request) {
	var form userSignUpForm
	
	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest, nil)
		return
	}

	// checking fields and validating form
	form.CheckFiled(validator.IsNotBlank(form.Name), "name", "Name can't be blank")
	form.CheckFiled(validator.IsNotBlank(form.Email), "email", "Email can't be blank")
	form.CheckFiled(validator.Matches(form.Email, validator.EmailRX), "email", "Not a valid email address")
	form.CheckFiled(validator.IsNotBlank(form.Password), "password", "Password can't be blank")
	form.CheckFiled(validator.IsMinChars(form.Password, 8), "password", "Password must be of atleast 8 charaters")
	
	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, "signup.tmpl.html", http.StatusUnprocessableEntity, data)
		return
	}

	err = app.users.Insert(form.Name, form.Email, form.Password)
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			data := app.newTemplateData(r)
			form.AddFieldError("email", "This email is already in use")
			data.Form = form
			app.render(w, "signup.tmpl.html", http.StatusUnprocessableEntity, data)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.sessionManager.Put(r.Context(), "flash", "User created sucessfully, Log in.")

	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func (app *application) userLogin(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = &userLoginForm{}
	app.render(w, "login.tmpl.html", http.StatusOK, data)
}	

func (app *application) userLoginPost(w http.ResponseWriter, r *http.Request) {
	var form userLoginForm
	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusUnprocessableEntity, err)
	}

	// checking fields and validating form
	form.CheckFiled(validator.IsNotBlank(form.Email), "email", "Email can't be blank")
	form.CheckFiled(validator.Matches(form.Email, validator.EmailRX), "email", "Email is not valid")
	form.CheckFiled(validator.IsNotBlank(form.Password), "password", "Password should not be blank")

	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = &form
		app.render(w, "login.tmpl.html", http.StatusUnprocessableEntity, data)
		return
	}

	// authenticate
	id, err := app.users.Authenticate(form.Email, form.Password)
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredential){
			form.AddNonFiledError("Email or Password is incorrect")
			data := app.newTemplateData(r)
			data.Form = form
			app.render(w, "login.tmpl.html", http.StatusUnprocessableEntity, data)
		} else {
			app.serverError(w, err)
		}
		return
	}

	err = app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.sessionManager.Put(r.Context(), "authenticatedUserId", id)

	http.Redirect(w, r, "/snippet/create", http.StatusSeeOther)
}

func (app *application) userLogoutPost(w http.ResponseWriter, r *http.Request) {

	err := app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.sessionManager.Remove(r.Context(), "authenticatedUserId")
	app.sessionManager.Put(r.Context(), "flash", "You have been logged out")

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

