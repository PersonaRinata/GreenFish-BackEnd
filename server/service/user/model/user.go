package model

type User struct {
	ID              int64
	Username        string `gorm:"index:idx_username,unique"`
	Password        string
	Avatar          string
	Nickname        string
	BackGroundImage string
	Signature       string
	Department      string
}

type DiseaseRelation struct {
	DiseaseIntroduction string
	FamilyDiseases      string
	HistoryDiseases     []HistoryDisease
}

type HistoryDisease struct {
	Symptom    string
	Medicines  []string
	Department string
	UpdateTime int64
}

type BodyInfo struct {
	BloodPressure string // 血压
	HeartRate     string // 心率
	Height        string // 身高
	Weight        string // 体重
	BloodSugar    string // 血糖
	UpdateTime    int64  // 更新时间
}

type IssueList struct {
	Username        string          // 用户名
	Gender          bool            // 性别
	Age             int32           // 年龄
	DiseaseRelation DiseaseRelation // 病史信息
	BodyInfo        BodyInfo        // 身体指标
}
