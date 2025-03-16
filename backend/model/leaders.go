package model

type Leader struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

const (
	TheodoreRoosevelt = iota
	AbrahamLincoln
	Saladin
	Montezuma
	PedroII
	QinShiHuang
	WuZetian
	FrederickBarbarossa
	LudwigII
	Victoria
	ElizabethI
	Gandhi
	Chandragupta
	Gilgamesh
	CatherineMedici
	EleanorAquitaine
	PeterI
	KublaiKhan
	GenghisKhan
	Tribhuwana
	MenelikII
	Tokugawa
	RamessesII
	Cleopatra
	PhilipII
	Suleiman
	Ambiorix
	Shaka
)

var Leaders = [...]Leader{
	{TheodoreRoosevelt, "Теодор Рузвельт", ""},
	{AbrahamLincoln, "Авраам Линкольн", ""},
	{Saladin, "Саладин", ""},
	{Montezuma, "Монтесума", ""},
	{PedroII, "Педру II", ""},
	{QinShiHuang, "Цинь Шихуанди", ""},
	{WuZetian, "Ву Цзэтянь", ""},
	{FrederickBarbarossa, "Фридрих Барбаросса", ""},
	{LudwigII, "Людвиг II", ""},
	{Victoria, "Виктория", ""},
	{ElizabethI, "Елизавета I", ""},
	{Gandhi, "Ганди", ""},
	{Chandragupta, "Чандрагупта", ""},
	{Gilgamesh, "Гильгамеш", ""},
	{CatherineMedici, "Катерина Медичи", ""},
	{EleanorAquitaine, "Элеонора Аквитанская", ""},
	{PeterI, "Петр I", ""},
	{KublaiKhan, "Хубилай", ""},
	{GenghisKhan, "Чингисхан", ""},
	{Tribhuwana, "Трибхувана ", ""},
	{MenelikII, "Менелик II", ""},
	{Tokugawa, "Сёгун Токугава", ""},
	{RamessesII, "Рамсес II", ""},
	{Cleopatra, "Клеопатра", ""},
	{PhilipII, "Филипп II", ""},
	{Suleiman, "Сулейман", ""},
	{Ambiorix, "Амбиорикс", ""},
	{Shaka, "Чака", ""},
}
