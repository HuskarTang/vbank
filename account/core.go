package account

import "time"

type VbkAccount struct {
	// Id 账户ID, 雪花ID
	Id uint64

	//RemarkName 账户备注名称
	RemarkName string

	//AccessKey 查询密码
	AccessKey string

	//TransactionKey 交易密码
	TransactionKey string

	// Balance 可用余额，coins
	Balance uint64

	//FrozenBalance 冻结余额，coins
	FrozenBalance uint64

	//Status 账户状态
	Status uint16

	// OwnerId 账户所属用户ID
	OwnerId uint64

	// CreateAt 创建事件
	CreateAt time.Time
}
