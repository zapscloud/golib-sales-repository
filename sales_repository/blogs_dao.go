package sales_repository

import (
	"github.com/zapscloud/golib-dbutils/db_common"
	"github.com/zapscloud/golib-sales-repository/sales_repository/mongodb_repository"
	"github.com/zapscloud/golib-utils/utils"
)

// BlogDao - Card DAO Repository
type BlogDao interface {
	// InitializeDao
	InitializeDao(client utils.Map, businessId string)
	//List - List all Collections
	List(filter string, sort string, skip int64, limit int64) (utils.Map, error)
	// Get - Get by code
	Get(blogId string) (utils.Map, error)
	// Find - Find by filter
	Find(filter string) (utils.Map, error)
	// Create - Create Collection
	Create(indata utils.Map) (utils.Map, error)
	// Update - Update Collection
	Update(blogId string, indata utils.Map) (utils.Map, error)
	// Delete - Delete Collection
	Delete(blogId string) (int64, error)
}

// NewBlogDao - Contruct Business Blog Dao
func NewBlogDao(client utils.Map, business_id string) BlogDao {
	var daoBlog BlogDao = nil

	// Get DatabaseType and no need to validate error
	// since the dbType was assigned with correct value after dbService was created
	dbType, _ := db_common.GetDatabaseType(client)

	switch dbType {
	case db_common.DATABASE_TYPE_MONGODB:
		daoBlog = &mongodb_repository.BlogMongoDBDao{}
	case db_common.DATABASE_TYPE_ZAPSDB:
		// *Not Implemented yet*
	case db_common.DATABASE_TYPE_MYSQLDB:
		// *Not Implemented yet*
	}

	if daoBlog != nil {
		// Initialize the Dao
		daoBlog.InitializeDao(client, business_id)
	}

	return daoBlog
}
