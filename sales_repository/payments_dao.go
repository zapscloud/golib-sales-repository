package sales_repository

import (
	"github.com/zapscloud/golib-dbutils/db_common"
	"github.com/zapscloud/golib-sales-repository/sales_repository/mongodb_repository"
	"github.com/zapscloud/golib-utils/utils"
)

// PaymentDao - Card DAO Repository
type PaymentDao interface {
	// InitializeDao
	InitializeDao(client utils.Map, businessId string)
	//List - List all Collections
	List(filter string, sort string, skip int64, limit int64) (utils.Map, error)
	// Get - Get by code
	Get(paymentId string) (utils.Map, error)
	// Find - Find by filter
	Find(filter string) (utils.Map, error)
	// Create - Create Collection
	Create(indata utils.Map) (utils.Map, error)
	// Update - Update Collection
	Update(paymentId string, indata utils.Map) (utils.Map, error)
	// Delete - Delete Collection
	Delete(paymentId string) (int64, error)
}

// NewPaymentDao - Contruct Business Payment Dao
func NewPaymentDao(client utils.Map, business_id string) PaymentDao {
	var daoPayment PaymentDao = nil

	// Get DatabaseType and no need to validate error
	// since the dbType was assigned with correct value after dbService was created
	dbType, _ := db_common.GetDatabaseType(client)

	switch dbType {
	case db_common.DATABASE_TYPE_MONGODB:
		daoPayment = &mongodb_repository.PaymentMongoDBDao{}
	case db_common.DATABASE_TYPE_ZAPSDB:
		// *Not Implemented yet*
	case db_common.DATABASE_TYPE_MYSQLDB:
		// *Not Implemented yet*
	}

	if daoPayment != nil {
		// Initialize the Dao
		daoPayment.InitializeDao(client, business_id)
	}

	return daoPayment
}
