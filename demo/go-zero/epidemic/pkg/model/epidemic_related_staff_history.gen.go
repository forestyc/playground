// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameEpidemicRelatedStaffHistory = "epidemic_related_staff_history"

// EpidemicRelatedStaffHistory mapped from table <epidemic_related_staff_history>
type EpidemicRelatedStaffHistory struct {
	ID            int32     `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id"`
	Name          []byte    `gorm:"column:name;type:varbinary(500);index:name,priority:1" json:"name"`
	Company       string    `gorm:"column:company;type:varchar(50);not null" json:"company"`               // 公司：字典值
	Department    string    `gorm:"column:department;type:varchar(50);not null" json:"department"`         // 部门：字典值
	Location      string    `gorm:"column:location;type:varchar(50);not null" json:"location"`             // 工作地点：字典值
	IsolationType string    `gorm:"column:isolation_type;type:varchar(50);not null" json:"isolation_type"` // 隔离情况分类：字典值
	Details       string    `gorm:"column:details;type:text;not null" json:"details"`                      // 具体情况
	HomeBeginDate time.Time `gorm:"column:home_begin_date;type:date;not null" json:"home_begin_date"`      // 观察起始日期
	CompanyOaDate time.Time `gorm:"column:company_oa_date;type:date" json:"company_oa_date"`               // 公司OA提交
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime(3)" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:datetime(3)" json:"updated_at"`
	Mobile        string    `gorm:"column:mobile;type:varchar(64);index:mobile,priority:1" json:"mobile"`
	ReturnDate    time.Time `gorm:"column:return_date;type:date" json:"return_date"`                     // 返岗日期
	RelatedID     int32     `gorm:"column:related_id;type:int" json:"related_id"`                        // 关联id
	Organization  string    `gorm:"column:organization;type:varchar(50);not null" json:"organization"`   // 编制：字典值
	AppletReport  string    `gorm:"column:applet_report;type:varchar(50);not null" json:"applet_report"` // 是否上传小程序：字典值
}

// TableName EpidemicRelatedStaffHistory's table name
func (*EpidemicRelatedStaffHistory) TableName() string {
	return TableNameEpidemicRelatedStaffHistory
}