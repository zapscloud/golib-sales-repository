package sales_repository

import (
	"github.com/zapscloud/golib-dbutils/db_common"
	"github.com/zapscloud/golib-sales-repository/sales_repository/mongodb_repository"
	"github.com/zapscloud/golib-utils/utils"
)

// ServiceDao - Service DAO Repository
type ServiceDao interface {
	// InitializeDao
	InitializeDao(client utils.Map, businessId string)
	//List - List all Collections
	List(filter string, sort string, skip int64, limit int64) (utils.Map, error)
	// Find - Find by code
	Get(serviceid string) (utils.Map, error)
	// Find - Find by filter
	Find(filter string) (utils.Map, error)
	// Create - Create Collection
	Create(indata utils.Map) (utils.Map, error)
	// Update - Update Collection
	Update(serviceid string, indata utils.Map) (utils.Map, error)
	// Delete - Delete Collection
	Delete(serviceid string) (int64, error)
}

// NewServiceMongoDao - Contruct Business Service Dao
func NewServiceDao(client utils.Map, business_id string) ServiceDao {
	var daoService ServiceDao = nil

	// Get DatabaseType and no need to validate error
	// since the dbType was assigned with correct value after dbService was created
	dbType, _ := db_common.GetDatabaseType(client)

	switch dbType {
	case db_common.DATABASE_TYPE_MONGODB:
		daoService = &mongodb_repository.ServiceMongoDBDao{}
	case db_common.DATABASE_TYPE_ZAPSDB:
		// *Not Implemented yet*
	case db_common.DATABASE_TYPE_MYSQLDB:
		// *Not Implemented yet*
	}

	if daoService != nil {
		// Initialize the Dao
		daoService.InitializeDao(client, business_id)
	}

	return daoService
}
