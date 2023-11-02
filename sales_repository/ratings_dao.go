package sales_repository

import (
	"github.com/zapscloud/golib-dbutils/db_common"
	"github.com/zapscloud/golib-sales-repository/sales_repository/mongodb_repository"
	"github.com/zapscloud/golib-utils/utils"
)

// RatingsDao - Card DAO Repository
type RatingsDao interface {
	// InitializeDao
	InitializeDao(client utils.Map, businessId string)
	//List - List all Collections
	List(filter string, sort string, skip int64, limit int64) (utils.Map, error)
	// Get - Get by code
	Get(ratingId string) (utils.Map, error)
	// Find - Find by filter
	Find(filter string) (utils.Map, error)
	// Create - Create Collection
	Create(indata utils.Map) (utils.Map, error)
	// Update - Update Collection
	Update(ratingId string, indata utils.Map) (utils.Map, error)
	// Delete - Delete Collection
	Delete(ratingId string) (int64, error)
}

// NewRatingsDao - Contruct Business Ratings Dao
func NewRatingsDao(client utils.Map, business_id string) RatingsDao {
	var daoRatings RatingsDao = nil

	// Get DatabaseType and no need to validate error
	// since the dbType was assigned with correct value after dbService was created
	dbType, _ := db_common.GetDatabaseType(client)

	switch dbType {
	case db_common.DATABASE_TYPE_MONGODB:
		daoRatings = &mongodb_repository.RatingsMongoDBDao{}
	case db_common.DATABASE_TYPE_ZAPSDB:
		// *Not Implemented yet*
	case db_common.DATABASE_TYPE_MYSQLDB:
		// *Not Implemented yet*
	}

	if daoRatings != nil {
		// Initialize the Dao
		daoRatings.InitializeDao(client, business_id)
	}

	return daoRatings
}
