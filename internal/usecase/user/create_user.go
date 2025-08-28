package user

import (
	"context"

	"github.com/reinheimermat/gobid/internal/validator"
)

type CreateUserReq struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
}

func (req CreateUserReq) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(validator.NotBlank(req.UserName), "user_name", "user_name is required")
	eval.CheckField(validator.NotBlank(req.Email), "email", "email is required")
	eval.CheckField(validator.NotBlank(req.Bio), "bio", "bio is required")
	eval.CheckField(validator.MinChars(req.Bio, 10), "bio", "bio must be at least 10 characters long")
	eval.CheckField(validator.MaxChars(req.Bio, 255), "bio", "bio must be at most 255 characters long")
	eval.CheckField(validator.MinChars(string(req.Password), 8), "password", "password must be at least 8 characters long")
	eval.CheckField(validator.Matches(req.Email, validator.EmailRX), "email", "email must be a valid email address")

	return eval
}
