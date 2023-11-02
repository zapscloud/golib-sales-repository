package sales_repository

import (
	"github.com/zapscloud/golib-dbutils/db_common"
	"github.com/zapscloud/golib-sales-repository/sales_repository/mongodb_repository"
	"github.com/zapscloud/golib-utils/utils"
)

// PoliciesDao - Card DAO Repository
type PoliciesDao interface {
	// InitializeDao
	InitializeDao(client utils.Map, businessId string)
	//List - List all Collections
	List(filter string, sort string, skip int64, limit int64) (utils.Map, error)
	// Get - Get by code
	Get(policiesId string) (utils.Map, error)
	// Find - Find by filter
	Find(filter string) (utils.Map, error)
	// Create - Create Collection
	Create(indata utils.Map) (utils.Map, error)
	// Update - Update Collection
	Update(policiesId string, indata utils.Map) (utils.Map, error)
	// Delete - Delete Collection
	Delete(policiesId string) (int64, error)
}

// NewpoliciesDao - Contruct Business Policies Dao
func NewPoliciesDao(client utils.Map, business_id string) PoliciesDao {
	var daoPolicies PoliciesDao = nil

	// Get DatabaseType and no need to validate error
	// since the dbType was assigned with correct value after dbService was created
	dbType, _ := db_common.GetDatabaseType(client)

	switch dbType {
	case db_common.DATABASE_TYPE_MONGODB:
		daoPolicies = &mongodb_repository.PoliciesMongoDBDao{}
	case db_common.DATABASE_TYPE_ZAPSDB:
		// *Not Implemented yet*
	case db_common.DATABASE_TYPE_MYSQLDB:
		// *Not Implemented yet*
	}

	if daoPolicies != nil {
		// Initialize the Dao
		daoPolicies.InitializeDao(client, business_id)
	}

	return daoPolicies
}
