package shop

import (
	"fmt"

	"github.com/enghasib/laundry_service/domain"
	"github.com/enghasib/laundry_service/service/shop"
	"github.com/google/uuid"

	"github.com/jmoiron/sqlx"
)

type ShopRepo interface {
	shop.ShopRepo
}

type shopRepo struct {
	db *sqlx.DB
}

func NewShopRepo(db *sqlx.DB) ShopRepo {
	return &shopRepo{
		db: db,
	}
}

func (rep *shopRepo) Create(shop domain.Shop) (*domain.Shop, error) {

	var userUniqueId string

	if shop.ShopOwner == "" {
		userUniqueId = shop.CreatedBy
	} else {
		userQuery := `
		SELECT unique_id FROM users WHERE email=$1
	`
		userRow := rep.db.QueryRow(userQuery, shop.ShopOwner)
		userRow.Scan(&userUniqueId)

		if userRow.Err() != nil {
			fmt.Println("err", userRow.Err())
			return nil, userRow.Err()
		}
	}

	query := `
		INSERT INTO shops(
			unique_id,
			name,
			location,
			contact,
			payment_details,
			created_by,
			shop_owner
		) VALUES(
			$1, $2, $3, $4, $5, $6, $7
		) 
		RETURNING id
	`
	shop.Uuid = uuid.New().String()

	row := rep.db.QueryRow(
		query,
		shop.Uuid,
		shop.Name,
		shop.ContactNumber,
		shop.Location,
		shop.PaymentDetails,
		shop.CreatedBy,
		userUniqueId,
	)

	if row.Err() != nil {
		fmt.Println("err", row.Err())
		return nil, row.Err()
	}

	row.Scan(&shop.Id)

	return &shop, nil

}

func (rep *shopRepo) Get(shopId string) (*domain.Shop, error) {

	var shop domain.Shop

	query := `
		SELECT
			id,
			unique_id,
			name,
			location,
			payment_details,
			created_by,
			shop_owner,
			created_at,
			updated_at
		FROM shops
		WHERE unique_id = $1
	`
	err := rep.db.Get(&shop, query, shopId)

	if err != nil {
		fmt.Println("err:", err.Error())
		return nil, err
	}
	return &shop, nil
}
func (rep *shopRepo) List(limit, page int) ([]*domain.Shop, error) {

	if limit <= 0 || limit >= 100 {
		limit = 10
	}
	if page <= 0 {
		page = 1
	}

	query := `
		SELECT
			id,
			unique_id,
			name,
			location,
			payment_details,
			created_by,
			shop_owner,
			created_at,
			updated_at
		FROM shops
		ORDER BY id
		LIMIT $1 OFFSET $2
	`
	offset := (page - 1) * limit

	rows, err := rep.db.Queryx(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shops []*domain.Shop
	for rows.Next() {
		var shop domain.Shop
		err := rows.Scan(
			&shop.Id,
			&shop.Uuid,
			&shop.Name,
			&shop.Location,
			&shop.PaymentDetails,
			&shop.CreatedBy,
			&shop.ShopOwner,
			&shop.CreatedAt,
			&shop.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		shops = append(shops, &shop)
	}

	return shops, nil
}

func (rep *shopRepo) Update(shopId string, Shop domain.Shop) (*domain.Shop, error) {

	var shopOwnerUniqueId string

	if Shop.ShopOwner != "" {
		userQuery := `
		SELECT unique_id FROM users WHERE email=$1
	`
		userRow := rep.db.QueryRow(userQuery, Shop.ShopOwner)
		userRow.Scan(&shopOwnerUniqueId)

		if userRow.Err() != nil {
			fmt.Println("err", userRow.Err())
			return nil, userRow.Err()
		}
	}

	query := `
		UPDATE shops
		SET
			name = $1,
			location = $2,
			contact = $3,
			payment_details = $4
			shop_owner = $5
		WHERE unique_id = $6
		RETURNING
			id,
			unique_id,
			name,
			location,
			payment_details
	`
	row := rep.db.QueryRow(
		query,
		Shop.Name,
		Shop.Location,
		Shop.ContactNumber,
		Shop.PaymentDetails,
		shopOwnerUniqueId,
		shopId,
	)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var updatedShop domain.Shop
	err := row.Scan(
		&updatedShop.Id,
		&updatedShop.Uuid,
		&updatedShop.Name,
		&updatedShop.Location,
		&updatedShop.PaymentDetails,
	)
	if err != nil {
		return nil, err
	}

	return &updatedShop, nil
}

func (rep *shopRepo) Delete(shopId string) error {
	query := `
		DELETE FROM shops
		WHERE unique_id = $1
	`
	_, err := rep.db.Exec(query, shopId)
	if err != nil {
		return err
	}
	return nil
}
