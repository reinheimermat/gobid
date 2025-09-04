package product

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/reinheimermat/gobid/internal/validator"
)

type CreateProductRequest struct {
	SellerID    uuid.UUID `json:"seller_id"`
	ProductName string    `json:"product_name"`
	Description string    `json:"description"`
	Baseprice   float64   `json:"baseprice"`
	AuctionEnd  time.Time `json:"auction_end"`
}

const minAuctionDuration = 2 * time.Hour // Minimum auction duration of 2 hours

func (r CreateProductRequest) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(validator.NotBlank(r.ProductName), "product_name", "must be provided")
	eval.CheckField(validator.NotBlank(r.Description), "product_name", "must be provided")
	eval.CheckField(
		validator.MinChars(r.Description, 10) &&
			validator.MaxChars(r.Description, 255), "description", "must be between 10 and 255 characters")
	eval.CheckField(r.Baseprice > 0, "baseprice", "must be a positive number")
	eval.CheckField(time.Until(r.AuctionEnd) >= minAuctionDuration, "auction_end", "must be at least 2 hours from now")

	return eval
}
