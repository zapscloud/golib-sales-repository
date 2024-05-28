package sales_repository

import (
	"github.com/zapscloud/golib-dbutils/db_common"
	"github.com/zapscloud/golib-sales-repository/sales_repository/mongodb_repository"
	"github.com/zapscloud/golib-utils/utils"
)

// Product_unitDao - Product_unit DAO Repository
type Product_unitDao interface {
	// InitializeDao
	InitializeDao(client utils.Map, businessId string)

	// List
	List(filter string, sort string, skip int64, limit int64) (utils.Map, error)

	// Get - Get Product_unit Details
	Get(place_id string) (utils.Map, error)

	// Find - Find by code
	Find(filter string) (utils.Map, error)

	// Create - Create Product_unit
	Create(indata utils.Map) (utils.Map, error)

	// Update - Update Collection
	Update(place_id string, indata utils.Map) (utils.Map, error)

	// Delete - Delete Collection
	Delete(place_id string) (int64, error)
}

// NewProduct_unitDao - Contruct Site Dao
func NewProduct_unitDao(client utils.Map, businessid string) Product_unitDao {
	var daoProduct_unit Product_unitDao = nil

	// Get DatabaseType and no need to validate error
	// since the dbType was assigned with correct value after dbService was created
	dbType, _ := db_common.GetDatabaseType(client)

	switch dbType {
	case db_common.DATABASE_TYPE_MONGODB:
		daoProduct_unit = &mongodb_repository.Product_unitMongoDBDao{}
	case db_common.DATABASE_TYPE_ZAPSDB:
		// *Not Implemented yet*
	case db_common.DATABASE_TYPE_MYSQLDB:
		// *Not Implemented yet*
	}

	if daoProduct_unit != nil {
		// Initialize the Dao
		daoProduct_unit.InitializeDao(client, businessid)
	}

	return daoProduct_unit
}
