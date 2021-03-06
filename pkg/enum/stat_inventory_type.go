package enum

type StatInventoryType string

const (
	INVENTORY_FEED              StatInventoryType = "INVENTORY_FEED"              // 头条信息流（广告投放）
	INVENTORY_TEXT_LINK         StatInventoryType = "INVENTORY_TEXT_LINK"         // 头条文章详情页（已废弃）
	INVENTORY_VIDEO_FEED        StatInventoryType = "INVENTORY_VIDEO_FEED"        // 西瓜信息流（广告投放）
	INVENTORY_HOTSOON_FEED      StatInventoryType = "INVENTORY_HOTSOON_FEED"      // 火山信息流（广告投放）
	INVENTORY_AWEME_FEED        StatInventoryType = "INVENTORY_AWEME_FEED"        // 抖音信息流（广告投放）
	INVENTORY_UNION_SLOT        StatInventoryType = "INVENTORY_UNION_SLOT"        // 穿山甲（广告投放）
	UNION_BOUTIQUE_GAME         StatInventoryType = "UNION_BOUTIQUE_GAME"         // 穿山甲精选休闲游戏（广告投放）
	INVENTORY_UNION_SPLASH_SLOT StatInventoryType = "INVENTORY_UNION_SPLASH_SLOT" // 穿山甲开屏广告（广告投放）
	INVENTORY_AWEME_SEARCH      StatInventoryType = "INVENTORY_AWEME_SEARCH"      // 搜索广告——抖音位（广告投放）
	INVENTORY_SEARCH            StatInventoryType = "INVENTORY_SEARCH"            // 搜索广告——头条位（广告投放）
	INVENTORY_UNIVERSAL         StatInventoryType = "INVENTORY_UNIVERSAL"         // 通投智选（广告投放）
	INVENTORY_BEAUTY            StatInventoryType = "INVENTORY_BEAUTY"            // 轻颜相机
	INVENTORY_PIPIXIA           StatInventoryType = "INVENTORY_PIPIXIA"           // 皮皮虾
	INVENTORY_AUTOMOBILE        StatInventoryType = "INVENTORY_AUTOMOBILE"        // 懂车帝
	INVENTORY_STUDY             StatInventoryType = "INVENTORY_STUDY"             // 好好学习
	INVENTORY_FACE_U            StatInventoryType = "INVENTORY_FACE_U"            // faceu
)
