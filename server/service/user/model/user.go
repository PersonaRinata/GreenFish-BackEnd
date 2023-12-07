package model

type User struct {
	ID              int64
	Username        string `gorm:"index:idx_username,unique"`
	Password        string
	Avatar          string
	BackGroundImage string
	Signature       string
	Department      string
}

type MedicalHistoryInfo struct {
	Symptom     string // 患者的主要症状或原因
	Description string // 病情描述
	History     string // 过去的疾病、手术、药物过敏等历史记录
	FamilyInfo  string // 患者家族中与当前疾病相关的疾病或遗传疾病的记录
}

type BodyInfo struct {
	BloodPressure string // 血压
	HeartRate     string // 心率
	Height        string // 身高
	Weight        string // 体重
	CreateTime    int64  // 创建时间
	UpdateTime    int64  // 更新时间
}

type IssueList struct {
	UserID             string             // 用户ID
	Username           string             // 用户名
	Gender             bool               // 性别
	Age                int32              // 年龄
	CreateTime         int64              // 创建时间
	UpdateTime         int64              // 更新时间
	Department         []string           // 科室列表
	MedicalHistoryInfo MedicalHistoryInfo // 病史信息
	BodyInfo           BodyInfo           // 身体指标
	Introduction       string             // 简介
	Medicine           []string           // 药物列表
}
