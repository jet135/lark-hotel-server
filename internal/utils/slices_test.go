package utils

import (
	"fmt"
	"lark-hotel-server/internal/model/entity"
	"testing"
)

func TestRemoveStructDuplicates(t *testing.T) {
	// 示例数据
	bills := []entity.Bill{
		{Name: "Alice"},
		{Name: "Bob"},
		{Name: "Alice"},
		{Name: "Charlie"},
		{Name: "Bob"},
	}

	// 使用 RemoveDuplicates 去重
	uniqueBillss := RemoveStructDuplicates(bills, func(p entity.Bill) string {
		return p.Name
	})

	fmt.Println("After removing duplicates:")

	for _, bill := range uniqueBillss {
		fmt.Println(bill.Name)
	}
}
