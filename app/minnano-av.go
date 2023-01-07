package app

import (
	"bytes"
	"fmt"
	"stash-scrapers/common/utils"
	"stash-scrapers/services/log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	referer = "http://www.minnano-av.com"
	host    = "http://www.minnano-av.com"
)

// run
func MinnanoRun(item *Performers) {
	// list := getPerformerList()
	// for _, item := range list {
	body, jumpNum := Search(item.Name)
	MetaData(item, body, jumpNum)
	// }
}

// Search performer
func Search(actressName string) ([]byte, int) {
	url := fmt.Sprintf(`http://www.minnano-av.com/search_result.php?search_scope=actress&search_word=%s&search=+Go+`, actressName)
	code, body, err, jumpNum := utils.HTTPForMinnanoAV(url, referer, "")
	if err != nil || code != 200 {
		log.Println("搜索失败:", err, code)
		return nil, 0
	}
	// log.Println(code, string(body), jumpNum)
	referer = url
	return body, jumpNum
}

// prase MetaData from html
func MetaData(performer *Performers, body []byte, jumpNum int) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		log.Println(err)
		return
	}
	image := &PerformersImage{}
	// jump to detail page
	if jumpNum == 1 {
		detailPage(performer, doc, image)
	}
}

// pdetail page
func detailPage(performer *Performers, doc *goquery.Document, image *PerformersImage) {
	var title, content string
	var list []string
	aliasesMap := make(map[string]bool)
	// 元数据
	doc.Find(".act-profile").Find("table").Find("tr").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		// name = s.Find("h2").Text()
		title = s.Find("span").Text()
		content = s.Find("p").Text()
		if title == "別名" {
			i := strings.Index(content, "（")
			content = strings.ReplaceAll(content[:i], "　", "")
			aliasesMap[content] = true
		}
		if title == "生年月日" {
			i := strings.Index(content, "（")
			content = strings.ReplaceAll(content[:i], "年", "-")
			content = strings.ReplaceAll(content, "月", "-")
			content = strings.ReplaceAll(content, "日", "")
			performer.Birthdate.String = content
			performer.Birthdate.Valid = true
		}
		if title == "サイズ" {
			list = strings.Split(content, "/")
			var b, w, h, cup string
			for _, item := range list {
				item = strings.TrimSpace(item)
				if item[:1] == "T" { // 身高
					h, err := strconv.Atoi(item[1:])
					if err != nil {
						log.Println("change height fail", err, item[1:])
					} else {
						performer.Height.Int32 = int32(h)
						performer.Height.Valid = true
					}
				}
				if item[:1] == "B" { // 胸围
					log.Println(item[1:])
					i := strings.Index(item[1:], "(")
					b = item[1 : i+1]
					if len(item) > i+3 {
						cup = item[i+2 : i+3]
					}
				}
				if item[:1] == "W" { // 腰围
					w = item[1:]
				}
				if item[:1] == "H" { // 臀围
					h = item[1:]
				}
			}
			if len(b) != 0 && len(w) != 0 && len(h) != 0 && len(cup) != 0 {
				performer.Measurements.String = fmt.Sprintf("%s%s-%s-%s", b, cup, w, h)
				performer.Measurements.Valid = true
			}
		}
		if title == "AV出演期間" {
			if len(content) != 0 {
				performer.CareerLength.String = strings.ReplaceAll(content, "年", "")
				performer.CareerLength.Valid = true
			}
		}
		if title == "ブログ" {
			if len(content) != 0 {
				performer.Twitter.String = content
				performer.Twitter.Valid = true
			}
		}
		// log.Println(name, title, content)
	})
	imageURL, ok := doc.Find(".thumb").Find("img").Attr("src")
	if ok {
		image.PerformerID = performer.ID
		image.Image = GetImage(imageURL)
	}
	// 处理别名
	for k := range aliasesMap {
		if len(performer.Aliases.String) == 0 {
			performer.Aliases.String = k
			performer.Aliases.Valid = true
		} else {
			performer.Aliases.String += fmt.Sprintf(",%s", k)
		}
	}
	performer.Country.String = "JP"
	performer.Country.Valid = true
	performer.Ethnicity.String = "Asian"
	performer.Ethnicity.Valid = true
	performer.Gender.String = "FEMALE"
	performer.Gender.Valid = true
	savePerformer(performer)
	counts := checkPerformerImage(image)
	if counts > 0 {
		updatePerformerImage(image)
	} else {
		savePerformerImage(image)
	}
	log.Println(performer)
}

func GetImage(url string) string {
	if len(url) == 0 {
		return ""
	}
	code, body, err, _ := utils.HTTPForMinnanoAV(host+url, referer, "")
	if err != nil || code != 200 {
		log.Println("get image fail:", err, code)
		return ""
	}
	// log.Println(string(body))
	return string(body)
}
