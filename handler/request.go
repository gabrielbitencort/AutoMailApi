package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// RegisterUser

type RegisterUserRequest struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

func (r *RegisterUserRequest) Validate() error {
	if r.Name == "" && r.Email == "" && r.PasswordHash == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}
	if r.PasswordHash == "" {
		return errParamIsRequired("password_hash", "string")
	}

	return nil
}

// UpdateUser

type UpdateUserRequest struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

func (r *UpdateUserRequest) Validate() error {
	// If any fields is provided, validation is truthy
	if r.Name != "" || r.Email != "" || r.PasswordHash != "" {
		return nil
	}

	// If none of the fields were provided, return falsy
	return fmt.Errorf("at least one valid field must be provided")
}
