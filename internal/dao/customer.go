// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"lark-hotel-server/internal/dao/internal"
)

// internalCustomerDao is internal type for wrapping internal DAO implements.
type internalCustomerDao = *internal.CustomerDao

// customerDao is the data access object for table customer.
// You can define custom methods on it to extend its functionality as you wish.
type customerDao struct {
	internalCustomerDao
}

var (
	// Customer is globally public accessible object for table customer operations.
	Customer = customerDao{
		internal.NewCustomerDao(),
	}
)

// Fill with you ideas below.