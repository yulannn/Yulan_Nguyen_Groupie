package backend

type PlayerData struct {
	Tag          string  `json:"tag"`
	Name         string  `json:"name"`
	ExpLevel     int     `json:"expLevel"`
	Trophies     int     `json:"trophies"`
	BestTrophies int     `json:"bestTrophies"`
	Wins         int     `json:"wins"`
	Arena        Arena   `json:"arena"`
	Clan         Clan    `json:"clan"`
	CurrentDeck  []Card  `json:"currentDeck"`
	Badges       []Badge `json:"badges"`
	Card         []AllCard `json:"allCard"`
}

type CardData struct {
	Image string `json:"image"`
	Name  string `json:"name"`
}

type Favoris struct {
	ListeFavoris []string `json:"listeFavoris"`
}

type AllCard struct {
	Items []struct {
		Name              string `json:"name"`
		ID                int    `json:"id"`
		IconUrls          struct {
			Medium          string `json:"medium"`
			EvolutionMedium string `json:"evolutionMedium"`
		} `json:"iconUrls"`
	} `json:"items"`
}


type Badge struct {
	Name     string     `json:"name"`
	Level    int        `json:"level"`
	MaxLevel int        `json:"maxLevel"`
	IconUrls BadgesUrls `json:"iconUrls"`
}

type BadgesUrls struct {
	Large string `json:"large"`
}

type IconUrls struct {
	Medium string `json:"medium"`
}

type Card struct {
	Name     string   `json:"name"`
	IconUrls IconUrls `json:"iconUrls"`
}

type Arena struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Clan struct {
	Tag             string `json:"tag"`
	Name            string `json:"name"`
	Members         int    `json:"members"`
	ClanWarTrophies int    `json:"clanWarTrophies"`
}

type Clans struct {
	Clans []Clan `json:"items"`
}

type Coffre struct {
	Index int    `json:"index"`
	Name  string `json:"name"`
}

type ListeCoffres struct {
	Items []Coffre `json:"items"`
}

type ClanData struct {
	Tag               string       `json:"tag"`
	Name              string       `json:"name"`
	Type              string       `json:"type"`
	Description       string       `json:"description"`
	BadgeId           int          `json:"badgeId"`
	ClanScore         int          `json:"clanScore"`
	ClanWarTrophies   int          `json:"clanWarTrophies"`
	Location          Location     `json:"location"`
	RequiredTrophies  int          `json:"requiredTrophies"`
	DonationsPerWeek  int          `json:"donationsPerWeek"`
	ClanChestStatus   string       `json:"clanChestStatus"`
	ClanChestLevel    int          `json:"clanChestLevel"`
	ClanChestMaxLevel int          `json:"clanChestMaxLevel"`
	Members           int          `json:"members"`
	MemberList        []MemberInfo `json:"memberList"`
}

type Location struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	IsCountry bool   `json:"isCountry"`
}

type MemberInfo struct {
	Tag               string `json:"tag"`
	Name              string `json:"name"`
	Role              string `json:"role"`
	LastSeen          string `json:"lastSeen"`
	ExpLevel          int    `json:"expLevel"`
	Trophies          int    `json:"trophies"`
	Arena             Arena  `json:"arena"`
	ClanRank          int    `json:"clanRank"`
	PreviousClanRank  int    `json:"previousClanRank"`
	Donations         int    `json:"donations"`
	DonationsReceived int    `json:"donationsReceived"`
	ClanChestPoints   int    `json:"clanChestPoints"`
}

type BattleLog struct {
	Type     string   `json:"type"`
	Team     []Player `json:"team"`
	Opponent []Player `json:"opponent"`
}

type Player struct {
	Name         string `json:"name"`
	TrophyChange int    `json:"trophyChange"`
	Crowns       int    `json:"crowns"`
}


type TemplateData struct {
    Clans       []Clan
    Nom         string
    MinMembers  int
    MaxMembers  int
    MinTrophies int
	PagePrev int
	PageNext int
}

