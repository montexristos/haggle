package parsers

import (
	"fmt"
	"github.com/gocolly/colly"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"haggle/models"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
	"unicode"
)

type Parser interface {
	Initialize()
	//GetSiteID() int
	GetDB() *gorm.DB
	Scrape() (bool, error)
	ScrapeHome() (bool, error)
	ScrapeLive() (bool, error)
	ScrapeToday() (bool, error)
	ScrapeTournament(tournamentUrl string) (bool, error)
	GetEventID(event map[string]interface{}) string
	GetEventName(event map[string]interface{}) string
	GetEventCanonicalName(event map[string]interface{}) string
	GetEventDate(event map[string]interface{}) string
	GetEventIsAntepost(event map[string]interface{}) bool
	GetEventMarkets(event map[string]interface{}) []interface{}
	ParseMarketType(market map[string]interface{}) string
	MatchMarketType(market map[string]interface{}, marketType string) (models.MarketType, error)
	ParseMarketId(market map[string]interface{}) string
	ParseMarketName(market map[string]interface{}) string
	ParseSelectionName(selectionData map[string]interface{}) string
	ParseSelectionPrice(selectionData map[string]interface{}) float64
	ParseSelectionLine(selectionData map[string]interface{}, marketData map[string]interface{}) float64
	GetMarketSelections(marketData map[string]interface{}) []interface{}
	SetDB(db *gorm.DB)
	SetConfig(config *models.SiteConfig)
	GetConfig() *models.SiteConfig
	FetchEvent(e *models.Event) error
	GetEventUrl(event map[string]interface{}) string
}

func GetSiteID(p Parser) int {
	return p.GetConfig().SiteID
}

func ParseEvent(p Parser, event map[string]interface{}) (*models.Event, error) {
	if p.GetEventIsAntepost(event) {
		return &models.Event{}, fmt.Errorf("antepost")
	}
	eventID := p.GetEventID(event)
	if eventID == "" {
		return nil, fmt.Errorf("wrong event id")
	}
	eventName := p.GetEventName(event)
	if strings.Index(eventName, "SRL") > 0 {
		return nil, fmt.Errorf("skip parsing srl")
	}
	if strings.Index(eventName, "Srl") > 0 {
		return nil, fmt.Errorf("skip parsing srl")
	}
	if strings.Index(eventName, "Srl") > 0 {
		return nil, fmt.Errorf("skip parsing srl")
	}
	if strings.Index(eventName, "esport") > 0 {
		return nil, fmt.Errorf("skip parsing esport")
	}
	if strings.Index(eventName, "Esport") > 0 {
		return nil, fmt.Errorf("skip parsing esport")
	}
	eventCanonicalName := TransformName(eventName)
	eventMarkets := p.GetEventMarkets(event)
	date := p.GetEventDate(event)
	db := p.GetDB()
	siteID := GetSiteID(p)
	e := models.GetCreateEvent(p.GetDB(), eventID, siteID, eventName)
	e.Date = date
	e.CanonicalName = eventCanonicalName
	e.Url = p.GetEventUrl(event)
	// try to get all event details
	GetEventDetails(p, &e)

	if len(e.Markets) == 0 {
		markets := make([]models.Market, 0)
		for _, market := range eventMarkets {
			m, err := ParseMarket(p, market.(map[string]interface{}), e)
			if err == nil {
				markets = append(markets, m)
			}
		}
		e.Markets = markets
	}
	db.Debug().Session(&gorm.Session{FullSaveAssociations: true}).Save(&e)
	return &e, nil
}

func GetEventDetails(p Parser, e *models.Event) {
	err := p.FetchEvent(e)
	if err != nil {
		fmt.Errorf("error fetching event from parser %s", p.GetConfig().Id)
	}
}

func UpdateMarket(p Parser, market models.Market, markets []interface{}) {
	for _, m := range markets {
		if p.ParseMarketType(m.(map[string]interface{})) == market.Type {
			selections := ParseMarketSelections(p, m.(map[string]interface{}))
			UpdateSelections(p, market, selections)
		}
	}
}
func UpdateSelections(p Parser, market models.Market, selections []interface{}) {
	for _, sel := range selections {
		fmt.Println(sel)
	}
}

func ParseMarket(p Parser, market map[string]interface{}, event models.Event) (models.Market, error) {
	marketType := p.ParseMarketType(market)
	marketTypeId, matchError := p.MatchMarketType(market, marketType)
	if matchError != nil {
		return models.Market{}, matchError
	}
	marketName := p.ParseMarketName(market)
	//marketId := p.ParseMarketId(market)
	m := models.Market{
		Name:       marketName,
		Type:       marketType,
		MarketType: marketTypeId.Name,
	}
	selections := ParseMarketSelections(p, market)
	for _, selection := range selections {
		if selection != nil {
			sel := ParseSelection(p, market, selection.(map[string]interface{}))
			if sel.Line > 0.0 {
				m.Line = sel.Line
			}
			m.Selections = append(m.Selections, sel)
		}
	}
	p.GetDB().Debug().Session(&gorm.Session{FullSaveAssociations: true}).Updates(&m)
	return m, nil
}

func ParseMarketSelections(p Parser, market map[string]interface{}) []interface{} {
	selections := p.GetMarketSelections(market)
	return selections
}

func ParseSelection(p Parser, market map[string]interface{}, selection map[string]interface{}) models.Selection {
	sel := models.Selection{
		Name:  p.ParseSelectionName(selection),
		Price: p.ParseSelectionPrice(selection),
		Line:  p.ParseSelectionLine(selection, market),
	}
	return sel
}
func GetCollector() *colly.Collector {
	c := colly.NewCollector(
		//colly.CacheDir("./_instagram_cache/"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)

	// Limit the number of threads started by colly to one
	// slow but make sure we don't get banned.
	// can be limited to domain
	_ = c.Limit(&colly.LimitRule{
		Parallelism: 1,
		Delay:       3 * time.Second,
	})
	// Before making a request put the URL with
	// the key of "url" into the context of the request
	c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("path", r.URL.String())
		log.Println("Visiting", r.URL)
	})
	c.OnError(func(r *colly.Response, e error) {
		log.Println("error:", e, r.Request.URL, string(r.Body))
	})
	return c
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func TransformName(name string) string {
	name = strings.ReplaceAll(name, "/", "-")
	name = strings.ReplaceAll(name, "1. ", "")
	teams := strings.Split(name, "-")
	teamNames := GetNames()
	matchedNames := make([]string, 0)
	unmatchedNames := make(map[string]string)
	for _, team := range teams {
		team = strings.TrimSpace(team)
		teamName, err := MatchTeam(teamNames, team)
		if err != nil {
			matchedNames = append(matchedNames, team)
			teamNames[team] = team
			unmatched, _ := ioutil.ReadFile("config/names_unmatched.yaml")
			err := yaml.Unmarshal(unmatched, unmatchedNames)
			if err != nil {
				log.Println("error reading unmatched")
			}
			//write team names back to file
			unmatchedNames[team] = team
			var out []byte
			out, _ = yaml.Marshal(unmatchedNames)
			_ = ioutil.WriteFile("config/names_unmatched.yaml", out, os.FileMode(755))
		} else {
			matchedNames = append(matchedNames, teamName)
		}
	}
	return strings.Join(matchedNames, " - ")
}

func MatchTeam(teamNames map[string]interface{}, team string) (string, error) {
	if name, found := teamNames[team]; found {
		return name.(string), nil
	}
	return "", fmt.Errorf("not matched")
}

func GetNames() map[string]interface{} {
	var teamNames map[string]interface{}
	yamlFile, err := ioutil.ReadFile("config/names.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &teamNames)
	if err != nil {
		log.Printf("yamlFile.parse err   #%v ", err)
	}
	return teamNames
}

type rule struct {
	expr        *regexp.Regexp
	replacement string
}

func applyRules(s string, rules []rule) string {
	res := []byte(s)

	for _, rule := range rules {
		res = rule.expr.ReplaceAll(res, []byte(rule.replacement))
	}

	return string(res)
}

var greekToGreeklishRules = []rule{
	{regexp.MustCompile("ΓΧ"), "GX"},
	{regexp.MustCompile("γχ"), "gx"},
	{regexp.MustCompile("ΤΘ"), "T8"},
	{regexp.MustCompile("τθ"), "t8"},
	{regexp.MustCompile("(θη|Θη),"), "8h"},
	{regexp.MustCompile("ΘΗ"), "8H"},
	{regexp.MustCompile("αυ"), "au"},
	{regexp.MustCompile("Αυ"), "Au"},
	{regexp.MustCompile("ΑΥ"), "AY"},
	{regexp.MustCompile("ευ"), "eu"},
	{regexp.MustCompile("εύ"), "eu"},
	{regexp.MustCompile("εϋ"), "ey"},
	{regexp.MustCompile("εΰ"), "ey"},
	{regexp.MustCompile("Ευ"), "Eu"},
	{regexp.MustCompile("Εύ"), "Eu"},
	{regexp.MustCompile("Εϋ"), "Ey"},
	{regexp.MustCompile("Εΰ"), "Ey"},
	{regexp.MustCompile("ΕΥ"), "EY"},
	{regexp.MustCompile("ου"), "ou"},
	{regexp.MustCompile("ού"), "ou"},
	{regexp.MustCompile("οϋ"), "oy"},
	{regexp.MustCompile("οΰ"), "oy"},
	{regexp.MustCompile("Ου"), "Ou"},
	{regexp.MustCompile("Ού"), "Ou"},
	{regexp.MustCompile("Οϋ"), "Oy"},
	{regexp.MustCompile("Οΰ"), "Oy"},
	{regexp.MustCompile("ΟΥ"), "OY"},
	{regexp.MustCompile("Α"), "A"},
	{regexp.MustCompile("α"), "a"},
	{regexp.MustCompile("ά"), "a"},
	{regexp.MustCompile("Ά"), "A"},
	{regexp.MustCompile("Β"), "B"},
	{regexp.MustCompile("β"), "b"},
	{regexp.MustCompile("Γ"), "G"},
	{regexp.MustCompile("γ"), "g"},
	{regexp.MustCompile("Δ"), "D"},
	{regexp.MustCompile("δ"), "d"},
	{regexp.MustCompile("Ε"), "E"},
	{regexp.MustCompile("ε"), "e"},
	{regexp.MustCompile("έ"), "e"},
	{regexp.MustCompile("Έ"), "E"},
	{regexp.MustCompile("Ζ"), "Z"},
	{regexp.MustCompile("ζ"), "z"},
	{regexp.MustCompile("Η"), "H"},
	{regexp.MustCompile("η"), "h"},
	{regexp.MustCompile("ή"), "h"},
	{regexp.MustCompile("Ή"), "H"},
	{regexp.MustCompile("Θ"), "TH"},
	{regexp.MustCompile("θ"), "th"},
	{regexp.MustCompile("Ι"), "I"},
	{regexp.MustCompile("Ϊ"), "I"},
	{regexp.MustCompile("ι"), "i"},
	{regexp.MustCompile("ί"), "i"},
	{regexp.MustCompile("ΐ"), "i"},
	{regexp.MustCompile("ϊ"), "i"},
	{regexp.MustCompile("Ί"), "I"},
	{regexp.MustCompile("Κ"), "K"},
	{regexp.MustCompile("κ"), "k"},
	{regexp.MustCompile("Λ"), "L"},
	{regexp.MustCompile("λ"), "l"},
	{regexp.MustCompile("Μ"), "M"},
	{regexp.MustCompile("μ"), "m"},
	{regexp.MustCompile("Ν"), "N"},
	{regexp.MustCompile("ν"), "n"},
	{regexp.MustCompile("Ξ"), "KS"},
	{regexp.MustCompile("ξ"), "ks"},
	{regexp.MustCompile("Ο"), "O"},
	{regexp.MustCompile("ο"), "o"},
	{regexp.MustCompile("Ό"), "O"},
	{regexp.MustCompile("ό"), "o"},
	{regexp.MustCompile("Π"), "P"},
	{regexp.MustCompile("π"), "p"},
	{regexp.MustCompile("Ρ"), "R"},
	{regexp.MustCompile("ρ"), "r"},
	{regexp.MustCompile("Σ"), "S"},
	{regexp.MustCompile("σ"), "s"},
	{regexp.MustCompile("Τ"), "T"},
	{regexp.MustCompile("τ"), "t"},
	{regexp.MustCompile("Υ"), "Y"},
	{regexp.MustCompile("Ύ"), "Y"},
	{regexp.MustCompile("Ϋ"), "Y"},
	{regexp.MustCompile("ΰ"), "y"},
	{regexp.MustCompile("ύ"), "y"},
	{regexp.MustCompile("ϋ"), "y"},
	{regexp.MustCompile("υ"), "y"},
	{regexp.MustCompile("Φ"), "F"},
	{regexp.MustCompile("φ"), "f"},
	{regexp.MustCompile("Χ"), "X"},
	{regexp.MustCompile("χ"), "x"},
	{regexp.MustCompile("Ψ"), "Ps"},
	{regexp.MustCompile("ψ"), "ps"},
	{regexp.MustCompile("Ω"), "w"},
	{regexp.MustCompile("ω"), "w"},
	{regexp.MustCompile("Ώ"), "w"},
	{regexp.MustCompile("ώ"), "w"},
	{regexp.MustCompile("ς"), "s"},
	{regexp.MustCompile(";"), "?"},
}

// Greeklish returns s transliterated into greeklish.
func Greeklish(s string) string {
	return applyRules(s, greekToGreeklishRules)
}
