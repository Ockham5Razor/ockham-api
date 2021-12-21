package model

import (
	"gorm.io/gorm"
	"time"
)

type ServicePlanSubscription struct {
	gorm.Model

	SubscriptionTitle       string `gorm:"type:VARCHAR(24)"` // 标题
	SubscriptionDescription string `gorm:"type:LONGTEXT"`    // 描述
	SubscriptionEnabled     bool   // 启用

	TermSpacingDays int16     // 用期周期
	TermsRemaining  int16     // 用期剩余数：-1 -- 不限次数；0 -- 已扣光次数，当前不可用。1 -- 最后一次用期的示数，并生成一个 billing，下个月核算时将被扣为 0。
	TermLastDate    time.Time // 上次用期时间

	EachTermInheritsSurplusTraffic bool  // 每次用期继承剩余流量
	EachTermIncreasingTrafficBytes int64 // 每次用期增加流量

	SubscriptionRenewalAddingTerms    int16   // 续订后增加的用期数
	SubscriptionRenewalTimesRemaining int16   // 续订剩余次数：-1 -- 不限次数的永久合约；0 -- 已扣光次数不再可续期，如需继续使用需订阅新的 plan。1 -- 最后一次续约，需提醒用户。
	SubscriptionRenewalFee            float32 // 续订费用：将用于生成一个 billing。

	UserID uint  `json:"-"` // 用户（外键，屏蔽 json）
	User   *User `json:"-"` // 用户（引用，屏蔽 json）
}
