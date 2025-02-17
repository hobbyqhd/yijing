package models

// Stem 天干
type Stem string

const (
	Jia  Stem = "甲"
	Yi   Stem = "乙"
	Bing Stem = "丙"
	Ding Stem = "丁"
	Wu4  Stem = "戊"
	Ji   Stem = "己"
	Geng Stem = "庚"
	Xin  Stem = "辛"
	Ren  Stem = "壬"
	Gui  Stem = "癸"
)

// Branch 地支
type Branch string

const (
	Zi   Branch = "子"
	Chou Branch = "丑"
	Yin  Branch = "寅"
	Mao  Branch = "卯"
	Chen Branch = "辰"
	Si   Branch = "巳"
	Wu   Branch = "午"
	Wei  Branch = "未"
	Shen Branch = "申"
	You  Branch = "酉"
	Xu   Branch = "戌"
	Hai  Branch = "亥"
)

// Element 五行
type Element string

const (
	Wood  Element = "木"
	Fire  Element = "火"
	Earth Element = "土"
	Metal Element = "金"
	Water Element = "水"
)

// BaziPillar 八字柱（年月日时每柱包含天干和地支）
type BaziPillar struct {
	Stem   Stem   `json:"stem"`   // 天干
	Branch Branch `json:"branch"` // 地支
}

// BaziChart 八字命盘
type BaziChart struct {
	Year  BaziPillar `json:"year"`  // 年柱
	Month BaziPillar `json:"month"` // 月柱
	Day   BaziPillar `json:"day"`   // 日柱
	Hour  BaziPillar `json:"hour"`  // 时柱
}

// BaziReading 八字占卜结果
type BaziReading struct {
	Chart     BaziChart `json:"chart"`     // 八字命盘
	DayMaster Stem      `json:"dayMaster"` // 日主（日柱天干）
	Luck      string    `json:"luck"`      // 运势分析
	Elements  []Element `json:"elements"`  // 五行分布
}
