package data

import (
	datauser "github.com/buqiuwenda/gin-template/internal/data/user"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(
	NewData,
	ProvideDB,
	datauser.NewUserRepository,
)

func ProvideDB(d *Data) *gorm.DB {
	return d.DB()
}
