package sales_repository

import (
	"github.com/zapscloud/golib-dbutils/db_common"
	"github.com/zapscloud/golib-sales-repository/sales_repository/mongodb_repository"
	"github.com/zapscloud/golib-utils/utils"
)

// OTPDao - Contact DAO Repository
type OTPDao interface {
	// InitializeDao
	InitializeDao(client utils.Map, businessId string)

	// List
	List(filter string, sort string, skip int64, limit int64) (utils.Map, error)

	// Get - Get Contact Details
	Get(state_Id string) (utils.Map, error)

	// Find - Find by code
	Find(filter string) (utils.Map, error)

	// Create - Create Contact
	Create(indata utils.Map) (utils.Map, error)

	// Update - Update Collection
	Update(state_Id string, indata utils.Map) (utils.Map, error)

	// Delete - Delete Collection
	Delete(state_Id string) (int64, error)

	// DeleteAll - DeleteAll Collection
	DeleteAll() (int64, error)

	// Verify OTP
	Verify(key, otp string) (utils.Map, error)
}

// NewOTPDao - Contruct Holiday Dao
func NewOTPDao(client utils.Map, business_id string) OTPDao {
	var daoOTPDao OTPDao = nil

	// Get DatabaseType and no need to validate error
	// since the dbType was assigned with correct value after dbService was created
	dbType, _ := db_common.GetDatabaseType(client)

	switch dbType {
	case db_common.DATABASE_TYPE_MONGODB:
		daoOTPDao = &mongodb_repository.OTPMongoDBDao{}
	case db_common.DATABASE_TYPE_ZAPSDB:
		// *Not Implemented yet*
	case db_common.DATABASE_TYPE_MYSQLDB:
		// *Not Implemented yet*
	}

	if daoOTPDao != nil {
		// Initialize the Dao
		daoOTPDao.InitializeDao(client, business_id)
	}

	return daoOTPDao
}
