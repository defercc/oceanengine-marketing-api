package enum

type AdvertiserRole string

const (
	ROLE_ADVERTISER       AdvertiserRole = "ROLE_ADVERTISER"       // 普通广告主（直客）
	ROLE_CHILD_ADVERTISER AdvertiserRole = "ROLE_CHILD_ADVERTISER" // 普通广告主（代理商子客户）
	ROLE_CHILD_AGENT      AdvertiserRole = "ROLE_CHILD_AGENT"      // 二级代理商
	ROLE_AGENT            AdvertiserRole = "ROLE_AGENT"            // 一级代理商
)
