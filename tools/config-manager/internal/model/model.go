package model

// Site 对应 public/config/site.json
type Site struct {
	Brand            Brand             `json:"brand"`
	NavLinks         []NavLink         `json:"navLinks"`
	FooterLinks      []FooterGroup     `json:"footerLinks"`
	FriendLinks      []FriendLink      `json:"friendLinks,omitempty"`
	ICP              string            `json:"icp,omitempty"`
	HeroSlides       []HeroSlide       `json:"heroSlides"`
	Stats            []Stat            `json:"stats"`
	AboutTabs        []AboutTab        `json:"aboutTabs"`
	ResourceGroups   []ResourceGroup   `json:"resourceGroups"`
	ResourceArticles []ResourceArticle `json:"resourceArticles"`
	JoinInfo         JoinInfo          `json:"joinInfo"`
}

type Brand struct {
	Name       string `json:"name"`
	EnName     string `json:"enName"`
	FullEnName string `json:"fullEnName"`
	Slogan     string `json:"slogan"`
	School     string `json:"school"`
	Logo       string `json:"logo"`
}

type NavLink struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Kind  string `json:"kind"`
	To    string `json:"to,omitempty"`
}

type FooterGroup struct {
	Title string       `json:"title"`
	Items []FooterItem `json:"items"`
}

type FooterItem struct {
	Label string `json:"label"`
	To    string `json:"to"`
}

type FriendLink struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

type HeroSlide struct {
	Image    string `json:"image"`
	Tag      string `json:"tag"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Desc     string `json:"desc"`
}

type Stat struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Suffix string `json:"suffix"`
	Label  string `json:"label"`
}

type AboutTab struct {
	Key        string        `json:"key"`
	Label      string        `json:"label"`
	Images     []string      `json:"images,omitempty"`
	Paragraphs []string      `json:"paragraphs,omitempty"`
	Groups     []MemberGroup `json:"groups,omitempty"`
	Items      []string      `json:"items,omitempty"`
}

type MemberGroup struct {
	Role  string   `json:"role"`
	Names []string `json:"names"`
}

type ResourceGroup struct {
	Title  string   `json:"title"`
	Icon   string   `json:"icon"`
	Accent string   `json:"accent"`
	Links  []string `json:"links"`
}

type ResourceArticle struct {
	ID       string   `json:"id"`
	Category string   `json:"category"`
	Title    string   `json:"title"`
	Author   string   `json:"author"`
	Summary  string   `json:"summary"`
	Cover    string   `json:"cover"`
	Tags     []string `json:"tags"`
	Content  []string `json:"content"`
}

type JoinInfo struct {
	Eyebrow   string `json:"eyebrow"`
	Title     string `json:"title"`
	Highlight string `json:"highlight"`
	Subtitle  string `json:"subtitle"`
	Notice    Notice `json:"notice"`
	Steps     []Step `json:"steps"`
	Note      string `json:"note"`
	CTA       CTA    `json:"cta"`
}

type Notice struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type Step struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type CTA struct {
	Label string `json:"label"`
	To    string `json:"to"`
}

// Services 对应 public/config/services.json
type Services struct {
	Section    ServicesSection `json:"section"`
	Categories []string        `json:"categories"`
	Services   []Service       `json:"services"`
}

type ServicesSection struct {
	Eyebrow    string `json:"eyebrow"`
	Title      string `json:"title"`
	Highlight  string `json:"highlight"`
	Subtitle   string `json:"subtitle"`
	AllLabel   string `json:"allLabel"`
	Hint       string `json:"hint"`
	HintAction string `json:"hintAction"`
}

type Service struct {
	ID       string         `json:"id"`
	Icon     string         `json:"icon"`
	Title    string         `json:"title"`
	Category string         `json:"category"`
	Desc     string         `json:"desc"`
	Tags     []string       `json:"tags"`
	Status   string         `json:"status"`
	Hot      bool           `json:"hot,omitempty"`
	Access   *ServiceAccess `json:"access,omitempty"`
}

// ServiceAccess 描述点击「申请该服务」后的访问方式：
//   - modal 弹窗提示：弹出提示框显示 Note
//   - link  链接跳转：在新标签页打开 URL
type ServiceAccess struct {
	Type  string `json:"type,omitempty"`
	URL   string `json:"url,omitempty"`
	Label string `json:"label,omitempty"`
	Title string `json:"title,omitempty"`
	Note  string `json:"note,omitempty"`
}

// Activities 对应 public/config/activities.json
type Activities struct {
	API        string            `json:"api"`
	Section    ActivitiesSection `json:"section"`
	Activities []Activity        `json:"activities"`
}

type ActivitiesSection struct {
	Eyebrow       string `json:"eyebrow"`
	Title         string `json:"title"`
	Highlight     string `json:"highlight"`
	Subtitle      string `json:"subtitle"`
	UpcomingLabel string `json:"upcomingLabel"`
	PastLabel     string `json:"pastLabel"`
	EmptyText     string `json:"emptyText"`
}

type Activity struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Date     string `json:"date"`
	Location string `json:"location"`
	Status   string `json:"status"`
	Type     string `json:"type"`
	Summary  string `json:"summary"`
}

// Directions 对应 public/config/directions.json
type Directions struct {
	Section    DirectionsSection `json:"section"`
	Directions []Direction       `json:"directions"`
	CTF        CTF               `json:"ctf"`
}

type DirectionsSection struct {
	Eyebrow   string `json:"eyebrow"`
	Title     string `json:"title"`
	Highlight string `json:"highlight"`
	Subtitle  string `json:"subtitle"`
}

type Direction struct {
	ID         string   `json:"id"`
	Icon       string   `json:"icon"`
	Name       string   `json:"name"`
	Tagline    string   `json:"tagline"`
	Desc       string   `json:"desc"`
	Highlights []string `json:"highlights"`
	Stack      []string `json:"stack"`
}

type CTF struct {
	Enabled      bool          `json:"enabled"`
	Eyebrow      string        `json:"eyebrow"`
	Title        string        `json:"title"`
	Subtitle     string        `json:"subtitle"`
	BestRank     BestRank      `json:"bestRank"`
	Team         Team          `json:"team"`
	Members      []Member      `json:"members"`
	Competitions []Competition `json:"competitions"`
}

type BestRank struct {
	Label string `json:"label"`
	Value string `json:"value"`
	Unit  string `json:"unit"`
	Title string `json:"title"`
	Note  string `json:"note"`
}

type Team struct {
	Name     string      `json:"name"`
	FullName string      `json:"fullName"`
	Slogan   string      `json:"slogan"`
	URL      string      `json:"url"`
	Fields   []TeamField `json:"fields"`
}

type TeamField struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type Member struct {
	Name      string `json:"name"`
	Handle    string `json:"handle"`
	Role      string `json:"role"`
	Direction string `json:"direction"`
}

type Competition struct {
	Date   string `json:"date"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Result string `json:"result"`
	Note   string `json:"note"`
}
