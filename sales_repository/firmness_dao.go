package sales_repository

import (
	"github.com/zapscloud/golib-dbutils/db_common"
	"github.com/zapscloud/golib-sales-repository/sales_repository/mongodb_repository"
	"github.com/zapscloud/golib-utils/utils"
)

// FirmnessDao - Card DAO Repository
type FirmnessDao interface {
	// InitializeDao
	InitializeDao(client utils.Map, businessId string)
	//List - List all Collections
	List(filter string, sort string, skip int64, limit int64) (utils.Map, error)
	// Get - Get by code
	Get(firmness_Id string) (utils.Map, error)
	// Find - Find by filter
	Find(filter string) (utils.Map, error)
	// Create - Create Collection
	Create(indata utils.Map) (utils.Map, error)
	// Update - Update Collection
	Update(firmness_Id string, indata utils.Map) (utils.Map, error)
	// Delete - Delete Collection
	Delete(firmness_Id string) (int64, error)
}

// NewFirmnessDao - Contruct Business Firmness Dao
func NewFirmnessDao(client utils.Map, business_id string) FirmnessDao {
	var daoFirmness FirmnessDao = nil

	// Get DatabaseType and no need to validate error
	// since the dbType was assigned with correct value after dbService was created
	dbType, _ := db_common.GetDatabaseType(client)

	switch dbType {
	case db_common.DATABASE_TYPE_MONGODB:
		daoFirmness = &mongodb_repository.FirmnessMongoDBDao{}
	case db_common.DATABASE_TYPE_ZAPSDB:
		// *Not Implemented yet*
	case db_common.DATABASE_TYPE_MYSQLDB:
		// *Not Implemented yet*
	}

	if daoFirmness != nil {
		// Initialize the Dao
		daoFirmness.InitializeDao(client, business_id)
	}

	return daoFirmness
}
