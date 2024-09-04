package handlers

import (
	"net/http"

	"website/internal/models"
)

func (h *Handler) UserSignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "HTTP method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := generateID()
	name := r.FormValue("name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	password := r.FormValue("password")

	newUser := &models.User{
		ID:       userID,
		Name:     name,
		Email:    email,
		Phone:    phone,
		Password: password,
		Pets:     make([]models.Pet, 0),
	}

	err := h.database.AddUser(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	session, err := store.Get(r, "e-leash-session")
	if err != nil {
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}

	session.Values["userID"] = userID
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, "Failed to save session", http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Redirect", "/home")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "HTTP method not allowed", http.StatusMethodNotAllowed)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := h.database.GetUserByEmailAndPassword(email, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	session, err := store.Get(r, "e-leash-session")
	if err != nil {
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}

	session.Values["userID"] = user.ID
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, "Failed to save session", http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Redirect", "/home")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UserLogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "HTTP method not allowed", http.StatusMethodNotAllowed)
		return
	}

	session, err := store.Get(r, "e-leash-session")
	if err != nil {
		http.Error(w, "Failed to get session", http.StatusInternalServerError)
		return
	}

	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, "Failed to delete session", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/login", http.StatusFound)
}
