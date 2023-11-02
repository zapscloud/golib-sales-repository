package customer_repository

import (
	"github.com/zapscloud/golib-dbutils/db_common"
	"github.com/zapscloud/golib-sales-repository/sales_repository/mongodb_repository/customer_mongodb_repository"
	"github.com/zapscloud/golib-utils/utils"
)

// CustomerReviewDao - Card DAO Repository
type CustomerReviewDao interface {
	// InitializeDao
	InitializeDao(client utils.Map, businessId string, customerId string)
	//List - List all Collections
	List(filter string, sort string, skip int64, limit int64) (utils.Map, error)
	// Get - Get by code
	Get(reviewId string) (utils.Map, error)
	// Find - Find by filter
	Find(filter string) (utils.Map, error)
	// Create - Create Collection
	Create(indata utils.Map) (utils.Map, error)
	// Update - Update Collection
	Update(reviewId string, indata utils.Map) (utils.Map, error)
	// Delete - Delete Collection
	Delete(reviewId string) (int64, error)
}

// NewCustomerReviewDao - Contruct Business CustomerReview Dao
func NewCustomerReviewDao(client utils.Map, business_id string, customerId string) CustomerReviewDao {
	var daoCustomerReview CustomerReviewDao = nil

	// Get DatabaseType and no need to validate error
	// since the dbType was assigned with correct value after dbService was created
	dbType, _ := db_common.GetDatabaseType(client)

	switch dbType {
	case db_common.DATABASE_TYPE_MONGODB:
		daoCustomerReview = &customer_mongodb_repository.CustomerReviewMongoDBDao{}
	case db_common.DATABASE_TYPE_ZAPSDB:
		// *Not Implemented yet*
	case db_common.DATABASE_TYPE_MYSQLDB:
		// *Not Implemented yet*
	}

	if daoCustomerReview != nil {
		// Initialize the Dao
		daoCustomerReview.InitializeDao(client, business_id, customerId)
	}

	return daoCustomerReview
}
