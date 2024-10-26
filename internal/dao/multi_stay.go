// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"lark-hotel-server/internal/dao/internal"
)

// internalMultiStayDao is internal type for wrapping internal DAO implements.
type internalMultiStayDao = *internal.MultiStayDao

// multiStayDao is the data access object for table multi_stay.
// You can define custom methods on it to extend its functionality as you wish.
type multiStayDao struct {
	internalMultiStayDao
}

var (
	// MultiStay is globally public accessible object for table multi_stay operations.
	MultiStay = multiStayDao{
		internal.NewMultiStayDao(),
	}
)

// Fill with you ideas below.
