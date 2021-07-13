package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ClientAllowance struct {
	Model
	Notary    string          `json:"notary" gorm:"size:255;index"`           // notary address
	Client    string          `json:"client" gorm:"size:255;index"`           // client address
	Allowance decimal.Decimal `json:"allowance" gorm:"type:decimal(30,0)"`    // quota
	BlockTime int64           `json:"block_time" gorm:"index"`                // time
	Epoch     uint64          `json:"epoch"`                                  // height
	SignedCid string          `json:"signed_cid" gorm:"size:255;uniqueIndex"` // info signature cid
}

func (ClientAllowance) TableName() string {
	return "client_allowance"
}

func init() {
	autoMigrateModels = append(autoMigrateModels, &ClientAllowance{})
}

func UpsertClientAllowance(c *ClientAllowance) error {
	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "signed_cid"}},
		DoUpdates: clause.AssignmentColumns([]string{"notary", "client", "allowance", "block_time", "epoch"}),
	}).Create(c).Error
}

func GetLastBlockTimeFromNotaryAllowance() (blockTime int64, err error) {
	err = db.Model(ClientAllowance{}).Select("block_time").Order("block_time desc").Limit(1).First(&blockTime).Error
	if err == gorm.ErrRecordNotFound {
		return 0, nil
	}
	return
}

type ClientAllowanceGranted struct {
	Date    int64           `json:"date"`
	Granted decimal.Decimal `json:"granted"`
}

func GetClientAllowanceGrantedDaily(limit int64) ([]*ClientAllowanceGranted, error) {
	sql := `
select 
*
from (
select
ca.block_time div 86400*86400 as date,
sum(ca.allowance) as granted
from client_allowance ca 
group by 1
order by 1 DESC 
limit ?
) t order by t.date
`
	var data []*ClientAllowanceGranted
	err := db.Raw(sql, limit).Scan(&data).Error
	return data, err
}
