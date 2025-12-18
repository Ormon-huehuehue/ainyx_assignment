package models

type CreateUserRequest struct {
	Name string `json:"name" validate:"required,min=2"`
	Dob  string `json:"dob" validate:"required,datetime=2006-01-02"` // year-month-day
}

type UpdateUserRequest struct {
	Name string `json:"name" validate:"required,min=2"`
	Dob  string `json:"dob" validate:"required,datetime=2006-01-02"`
}

type UserResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	Dob  string `json:"dob"`
	Age  int    `json:"age"`
}

type ListUsersResponse struct {
	Users []UserResponse `json:"users"`
	Total int64          `json:"total"`
	Page  int            `json:"page"`
	Limit int            `json:"limit"`
}
