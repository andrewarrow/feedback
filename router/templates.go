package router

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"github.com/andrewarrow/feedback/util"
	"github.com/xeonx/timeago"
)

var EmbeddedTemplates embed.FS

const DATE_LAYOUT = "Monday, January 2, 2006 15:04"
const HUMAN = "Monday, January 2, 2006 3:04 PM"
const HUMAN_SMALL = "01/02/2006"
const (
	KB = 1024
	MB = KB * 1024
	GB = MB * 1024
)

func TemplateFunctions() template.FuncMap {
	cfg := timeago.English
	cfg.Max = 9223372036854775807
	fm := template.FuncMap{
		"mod":    func(i, j int) bool { return i%j == 0 },
		"tokens": func(s string, i int) string { return strings.Split(s, ".")[i] },
		"add":    func(i, j int) int { return i + j },
		"mul":    func(i, j int64) int64 { return i * j },
		"k":      func(i int64) string { return fmt.Sprintf("%0.2f", float64(i)/1000.0) },
		"humanSize": func(bytes int64) string {

			size := float64(bytes)

			result := ""
			if size < KB {
				result = fmt.Sprintf("%d bytes", bytes)
			} else if size < MB {
				result = fmt.Sprintf("%.2f KB", size/KB)
			} else if size < GB {
				result = fmt.Sprintf("%.2f MB", size/MB)
			} else {
				result = fmt.Sprintf("%.2f GB", size/GB)
			}
			return result
		},
		"marshal": func(a any) string {
			if a == nil {
				return ""
			}
			b, _ := json.Marshal(a)
			s := string(b)
			if s == "null" {
				return ""
			}
			return s
		},
		"findInMap": func(send map[string]any, name string, id int64) any {
			m := send[name].(map[int64]any)
			return m[id]
		},
		"short": func(a any) string {
			tInt, ok := a.(int64)
			if ok {
				t := time.Unix(tInt, 0)
				return t.Format(HUMAN_SMALL)
			}
			return ""
		},
		"timeOptions": func(sa, ea float64, tz *time.Location) []template.HTML {

			saInt := int64(sa)
			eaInt := int64(ea)
			buffer := []template.HTML{}

			q := `"`
			t := time.Unix(saInt, 0).In(tz)
			now := time.Now().Unix()
			for {
				delta := now - t.Unix()
				val := fmt.Sprintf("%s%d%s", q, t.Unix(), q)
				if delta < -3600 {
					thing := fmt.Sprintf("<option value=%s>%s %s</option>", val, t.Format(HUMAN), cfg.Format(t))
					buffer = append(buffer, template.HTML(thing))
				}
				t = t.Add(time.Minute * 30)
				if t.Unix() >= eaInt {
					break
				}
			}

			return buffer
		},
		"hasSuffix": func(s any, suffix string) bool {
			val := s.(string)
			return strings.HasSuffix(val, suffix)
		},
		"null": func(s any) any {
			if s == nil {
				return ""
			}
			return s
		},
		"colorStatus": func(status, color string) map[string]any {
			m := map[string]any{}
			m["status"] = status
			m["color"] = color
			return m
		},
		"timeOptionList": func(val any) template.HTML {
			buffer := []string{}
			option := fmt.Sprintf("<option>%s</option>", "")
			buffer = append(buffer, option)
			now := time.Now()
			for i := 0; i < 9; i++ {
				option := fmt.Sprintf("<option value=\"%d\">%s</option>", now.Unix(),
					now.Format(HUMAN_SMALL))
				now = now.Add(time.Second * -86400)
				buffer = append(buffer, option)
			}
			return template.HTML(strings.Join(buffer, "\n"))
		},
		"guidOptionList": func(list []map[string]any, selectedGuid string) template.HTML {
			buffer := []string{}
			for _, item := range list {
				selected := ""
				guid := item["guid"].(string)
				if guid == selectedGuid {
					selected = `selected="true"`
				}
				option := fmt.Sprintf(`<option %s value="%s">%s</option>`,
					selected, item["guid"],
					item["name"])
				buffer = append(buffer, option)
			}
			return template.HTML(strings.Join(buffer, "\n"))
		},
		"optionList": func(s, val string) template.HTML {
			tokens := strings.Split(s[2:len(s)-2], "|")
			valTokens := strings.Split(val, ",")
			valMap := map[string]bool{}
			for _, item := range valTokens {
				valMap[item] = true
			}
			buffer := []string{}
			for _, token := range tokens {
				option := fmt.Sprintf("<option>%s</option>", token)
				if valMap[token] {
					option = fmt.Sprintf("<option selected=\"true\">%s</option>", token)
				}
				buffer = append(buffer, option)
			}
			return template.HTML(strings.Join(buffer, "\n"))
		},
		"intComma": func(a int64) string {
			if a == 0 {
				return "n/a"
			}
			return "$" + util.IntComma(a)
		},
		"breakComma": func(s string) []string { return strings.Split(s, ",") },
		"ago":        func(t time.Time) string { return cfg.Format(t) },
		"adds": func(i int64) string {
			if i != 1 {
				return "s"
			}
			return ""
		},
		"substring": func(s any, index int) string {
			text := fmt.Sprintf("%v", s)
			return text[:index]
		},
		"epoch": func(date string, hour float64, tzString string) int64 {
			dateString := fmt.Sprintf("%s %02d:00", date, int(hour))
			tz, _ := time.LoadLocation(tzString)
			t, _ := time.ParseInLocation(DATE_LAYOUT, dateString, tz)
			utc, _ := time.LoadLocation("UTC")
			return t.In(utc).Unix()
		},
		"simplePrice": func(s any) string {
			amount := fmt.Sprintf("%v", s)
			price, _ := strconv.Atoi(amount)
			if price == 0 {
				return "0.00"
			}
			if len(amount) < 3 {
				return fmt.Sprintf("00.%s", amount)
			}
			dollars := amount[0 : len(amount)-2]
			cents := amount[len(amount)-2:]
			return fmt.Sprintf("%s.%s", dollars, cents)
		},
		"price": func(s any) template.HTML {
			amount := fmt.Sprintf("%v", s)
			price, _ := strconv.Atoi(amount)
			if price == 0 {
				sp := "$0.00 USD"
				return template.HTML(sp)
			}

			if len(amount) < 3 {
				sp := fmt.Sprintf("$00.%s USD", amount)
				return template.HTML(sp)
			}
			dollars := amount[0 : len(amount)-2]
			dollarsInt, _ := strconv.ParseInt(dollars, 10, 64)
			dollars = util.IntComma(dollarsInt)
			cents := amount[len(amount)-2:]
			sp := fmt.Sprintf("$%s.%s USD", dollars, cents)
			if price < 0 {
				sp = "<span class='text-red-500'>" + sp + "</span>"
			}
			return template.HTML(sp)
		},
		"timezone": func(tz any) template.HTML {
			if tz == nil {
				return TimezoneList("UTC")
			}
			return TimezoneList(tz.(string))
		},
		"splitLines": func(thing string) template.HTML {
			tokens := strings.Split(thing, "\r")
			return template.HTML(strings.Join(tokens, "<br/>"))
		},
		"jq": func(thing string) string {
			return util.PipeToJq(thing)
		},
		"indent": func(thing any) string {
			var other any
			asBytes, ok := thing.([]byte)
			if ok {
				json.Unmarshal(asBytes, &other)
				asBytes, _ := json.MarshalIndent(other, "", "  ")
				return string(asBytes)
			}
			asString, ok := thing.(string)
			if ok {
				json.Unmarshal([]byte(asString), &other)
				asBytes, _ := json.MarshalIndent(other, "", "  ")
				return string(asBytes)
			}
			asBytes, _ = json.MarshalIndent(thing, "", "  ")
			return string(asBytes)
		},
		"chop": func(thing string) string {
			tokens := strings.Split(thing, ",")
			return strings.TrimSpace(tokens[1] + tokens[2])
		},
		"trim": func(thing string) string {
			return strings.TrimSpace(thing)
		},
		"extract": func(s, key string) any {
			var m map[string]any
			json.Unmarshal([]byte(s), &m)
			return int64(m[key].(float64))
		},
		"uuid": func() string {
			return util.PseudoUuid()
		},
		"textfield": func(name string, val any) template.HTML {
			value := ""
			if val != nil {
				value = fmt.Sprintf("%v", val)
			}
			s := fmt.Sprintf(`<input id="textfield-%s" class="nice-i border" name="%s" type="text" value="%s"/>`, name, name, value)
			return template.HTML(s)
		},
		"dob": func(thing any) string {
			tint64, ok := thing.(int64)
			if !ok {
				tfloat := thing.(float64)
				tint64 = int64(tfloat)
			}
			t := time.Unix(tint64, 0)

			return fmt.Sprintf("%d years old", calculateAge(t))
		},
		"zero": func(a any) string {
			s := fmt.Sprintf("%v", a)
			if s == "0" {
				return ""
			}
			return s
		},
		"zeroMoney": func(a any) string {
			s := fmt.Sprintf("%v", a)
			if s == "0" {
				return ""
			}
			return "$" + s
		},
		"ampm": func(f float64) string {
			i := int(f)
			if i > 12 {
				return fmt.Sprintf("%02d:00 PM", i-12)
			}
			if i == 0 {
				return "12:00 AM"
			}
			return fmt.Sprintf("%02d:00 AM", i)
		},
	}
	return fm
}

func LoadTemplates(tf template.FuncMap) *template.Template {
	t := template.New("")
	t = t.Funcs(tf)

	templateFiles, _ := EmbeddedTemplates.ReadDir("views")
	for _, file := range templateFiles {
		fileContents, _ := EmbeddedTemplates.ReadFile("views/" + file.Name())

		_, err := t.New(file.Name()).Parse(string(fileContents))
		if err != nil {
			fmt.Println(file.Name(), err)
		}
	}
	return t
}

func LoadLiveTemplates(tf template.FuncMap) *template.Template {
	t := template.New("")
	t = t.Funcs(tf)

	templateFiles, _ := ioutil.ReadDir("views")
	for _, file := range templateFiles {
		fileContents, _ := ioutil.ReadFile("views/" + file.Name())
		_, err := t.New(file.Name()).Parse(string(fileContents))
		if err != nil {
			fmt.Println(file.Name(), err)
		}
	}
	return t
}

func calculateAge(birthdate time.Time) int {
	now := time.Now()
	years := now.Year() - birthdate.Year()

	if now.YearDay() < birthdate.YearDay() {
		years--
	}

	return years
}

func TimezoneList(tz string) template.HTML {
	list := `UTC
Pacific/Honolulu
America/Anchorage
America/Los_Angeles
America/Denver
America/Chicago
America/New_York
America/Puerto_Rico
America/Santiago
America/Mexico_City
America/Bogota
America/Regina
America/Costa_Rica
America/Phoenix
America/Edmonton
America/Tijuana
America/Halifax
America/St_Johns
America/Manaus
America/Sao_Paulo
Atlantic/Cape_Verde
Europe/London
Europe/Berlin
Europe/Moscow
Asia/Kolkata
Asia/Shanghai
Asia/Tokyo
Australia/Sydney
Pacific/Auckland
Pacific/Fiji`
	tokens := strings.Split(list, "\n")
	buffer := []string{}
	for _, item := range tokens {
		selected := ""
		if item == tz {
			selected = `selected="true"`
		}
		buffer = append(buffer, fmt.Sprintf("<option %s>%s</option>", selected, item))
	}

	s := strings.Join(buffer, "\n")
	return template.HTML(s)

}
