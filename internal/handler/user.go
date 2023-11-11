package handler

import (
  "encoding/json"
  "log/slog"
  "net/http"
)

type User struct {
  ID    int    `json:"id"`
  Name  string `json:"name"`
  Email string `json:"email"`
}

// Get fake user
//	@Summary		Get Fake user
//	@Description	Search library attributes
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	User
//	@Failure		500
//	@Router			/user [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
  user := User{
    ID:    1,
    Name:  "John Doe",
    Email: "jonh.doe@email.com",
  }

  userMarshal, err := json.Marshal(user)
  if err != nil {
    slog.Error("Error marshalling user", err)
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Write([]byte(userMarshal))
}
