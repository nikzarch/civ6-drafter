package model

type Nation struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Leaders  []int  `json:"-"`
	ImageURL string `json:"image_url"`
}

var Nations = [...]Nation{
	{1, "Америка", []int{TheodoreRoosevelt, AbrahamLincoln}, ""},
	{2, "Аравия", []int{Saladin}, ""},
	{3, "Ацтеки", []int{Montezuma}, ""},
	{4, "Бразилия", []int{PedroII}, ""},
	{5, "Китай", []int{QinShiHuang, WuZetian, KublaiKhan}, ""},
	{6, "Германия", []int{FrederickBarbarossa, LudwigII}, ""},
	{7, "Англия", []int{Victoria, ElizabethI, EleanorAquitaine}, ""},
	{8, "Индия", []int{Gandhi, Chandragupta}, ""},
	{9, "Шумеры", []int{Gilgamesh}, ""},
	{10, "Франция", []int{CatherineMedici, EleanorAquitaine}, ""},
	{11, "Россия", []int{PeterI}, ""},
	{12, "Монголия", []int{KublaiKhan, GenghisKhan}, ""},
	{13, "Индонезия", []int{Tribhuwana}, ""},
	{14, "Эфиопия", []int{MenelikII}, ""},
	{15, "Япония", []int{Tokugawa}, ""},
	{16, "Египет", []int{RamessesII, Cleopatra}, ""},
	{17, "Испания", []int{PhilipII}, ""},
	{18, "Османская Империя", []int{Suleiman}, ""},
	{19, "Галлы", []int{Ambiorix}, ""},
	{20, "Зулусы", []int{Shaka}, ""},
}

func (n *Nation) GetLeaders() []Leader {
	var leaders []Leader
	for _, i := range n.Leaders {
		leaders = append(leaders, Leaders[i])
	}
	return leaders
}
