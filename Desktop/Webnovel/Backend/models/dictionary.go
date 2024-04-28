package models

type DictEntry struct {
	Word     string    `json:"word"`
	Pinyin   string    `json:"pinyin"`
	Phonetic string    `json:"phonetic"`
	Compound string    `json:"compound"`
	Kind     string    `json:"kind"`
	Measure  Measure   `json:"measure"`
	Snym     Snym      `json:"snym"`
	Content  []Content `json:"content"`
	Topics   []Topic   `json:"topic"`
	CNVI     string    `json:"cn_vi"`
	Examples []Example `json:"examples"`
}

type Example struct {
	Id string `json:"id"`
	E  string `json:"e"`
	M  string `json:"m"`
	P  string `json:"p"`
}

type Measure struct {
	Measure string `json:"measure"`
	Pinyin  string `json:"pinyin"`
	Example string `json:"example"`
}

type Snym struct {
	Anto []string `json:"anto"`
	Syno []string `json:"syno"`
}

type Content struct {
	Kind  string `json:"kind"`
	Means []Mean `json:"means"`
}

type Mean struct {
	Mean     string `json:"mean"`
	Explain  string `json:"explain"`
	Examples []int  `json:"examples"`
}

type Topic struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
