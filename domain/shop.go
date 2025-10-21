package domain

type Shop struct {
	Id             int    `json:"id" db:"id"`
	Uuid           string `json:"unique_id" db:"unique_id" validate:"required,unique"`
	Name           string `json:"name" db:"name" validate:"required"`
	Location       string `json:"location" db:"location"`
	ContactNumber  string `json:"contact" db:"contact"`
	PaymentDetails string `json:"payment_details" db:"payment_details"`
	CreatedBy      string `json:"created_by" db:"created_by"`
	ShopOwner      string `json:"shop_owner" db:"shop_owner"`
	CoverImage     string `json:"cover_image" db:"cover_image"`
	CreatedAt      string `json:"created_at" db:"created_at"`
	UpdatedAt      string `json:"updated_at" db:"updated_at"`
}
