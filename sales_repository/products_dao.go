package sales_repository

import (
	"github.com/zapscloud/golib-dbutils/db_common"
	"github.com/zapscloud/golib-sales-repository/sales_repository/mongodb_repository"
	"github.com/zapscloud/golib-utils/utils"
)

// ProductDao - Product DAO Repository
type ProductDao interface {
	// InitializeDao
	InitializeDao(client utils.Map, businessId string)
	//List - List all Collections
	List(filter string, sort string, skip int64, limit int64) (utils.Map, error)
	// Find - Find by code
	Get(productid string) (utils.Map, error)
	// Find - Find by filter
	Find(filter string) (utils.Map, error)
	// Create - Create Collection
	Create(indata utils.Map) (utils.Map, error)
	// Update - Update Collection
	Update(productid string, indata utils.Map) (utils.Map, error)
	// Delete - Delete Collection
	Delete(productid string) (int64, error)
}

// NewProductMongoDao - Contruct Business Product Dao
func NewProductDao(client utils.Map, business_id string) ProductDao {
	var daoProduct ProductDao = nil

	// Get DatabaseType and no need to validate error
	// since the dbType was assigned with correct value after dbService was created
	dbType, _ := db_common.GetDatabaseType(client)

	switch dbType {
	case db_common.DATABASE_TYPE_MONGODB:
		daoProduct = &mongodb_repository.ProductMongoDBDao{}
	case db_common.DATABASE_TYPE_ZAPSDB:
		// *Not Implemented yet*
	case db_common.DATABASE_TYPE_MYSQLDB:
		// *Not Implemented yet*
	}

	if daoProduct != nil {
		// Initialize the Dao
		daoProduct.InitializeDao(client, business_id)
	}

	return daoProduct
}
