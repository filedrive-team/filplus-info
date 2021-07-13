package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type NotaryAllowance struct {
	Model
	TxnId     int64           `json:"txn_id"`
	Address   string          `json:"address" gorm:"size:255;index"`          // notary address
	Allowance decimal.Decimal `json:"allowance" gorm:"type:decimal(30,0)"`    // quota
	BlockTime int64           `json:"block_time"`                             // time
	Epoch     uint64          `json:"epoch" gorm:"index"`                     // height
	SignedCid string          `json:"signed_cid" gorm:"size:255;uniqueIndex"` // info signature cid
}

func (NotaryAllowance) TableName() string {
	return "notary_allowance"
}

func init() {
	autoMigrateModels = append(autoMigrateModels, &NotaryAllowance{})
}

func UpsertNotaryAllowance(n *NotaryAllowance) error {
	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "signed_cid"}},
		DoUpdates: clause.AssignmentColumns([]string{"txn_id", "address", "allowance", "block_time", "epoch"}),
	}).Create(n).Error
}

func GetLastEpochFromNotaryAllowance() (epoch uint64, err error) {
	err = db.Model(NotaryAllowance{}).Select("epoch").Order("epoch desc").Limit(1).First(&epoch).Error
	if err == gorm.ErrRecordNotFound {
		return 0, nil
	}
	return
}
