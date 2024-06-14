package mock_db

import "gorm.io/gorm"

type DB interface {
	*gorm.DB
}
