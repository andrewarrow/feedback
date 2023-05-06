package router

import (
	"embed"
	"fmt"
	"html/template"
	"strings"
	"time"

	"github.com/xeonx/timeago"
)

var EmbeddedTemplates embed.FS

const DATE_LAYOUT = "Monday, January 2, 2006 15:04"
const HUMAN = "Monday, January 2, 2006 3:04 PM"

func TemplateFunctions() template.FuncMap {
	cfg := timeago.English
	cfg.Max = 9223372036854775807
	fm := template.FuncMap{
		"mod":    func(i, j int) bool { return i%j == 0 },
		"tokens": func(s string, i int) string { return strings.Split(s, ".")[i] },
		"add":    func(i, j int) int { return i + j },
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
		"ago": func(t time.Time) string { return cfg.Format(t) },
		"adds": func(i int64) string {
			if i != 1 {
				return "s"
			}
			return ""
		},
		"epoch": func(date string, hour float64, tzString string) int64 {
			dateString := fmt.Sprintf("%s %02d:00", date, int(hour))
			tz, _ := time.LoadLocation(tzString)
			t, _ := time.ParseInLocation(DATE_LAYOUT, dateString, tz)
			utc, _ := time.LoadLocation("UTC")
			return t.In(utc).Unix()
		},
		"price": func(thing any) template.HTML {
			thingInt64, ok := thing.(int64)
			if !ok {
				thingFloat64 := thing.(float64)
				thingInt64 = int64(thingFloat64)
			}
			if thingInt64 == 0 {
				s := "$0.00 USD"
				return template.HTML(s)
			}

			amount := fmt.Sprintf("%d", thingInt64)
			if len(amount) < 3 {
				s := fmt.Sprintf("$00.%s USD", amount)
				return template.HTML(s)
			}
			dollars := amount[0 : len(amount)-2]
			cents := amount[len(amount)-2:]
			s := fmt.Sprintf("$%s.%s USD", dollars, cents)
			if thingInt64 < 0 {
				s = "<span class='text-red-500'>" + s + "</span>"
			}
			return template.HTML(s)
		},
		"timezone": func(tz any) template.HTML {
			if tz == nil {
				return TimezoneList("UTC")
			}
			return TimezoneList(tz.(string))
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

func LoadTemplates() *template.Template {
	t := template.New("")
	t = t.Funcs(TemplateFunctions())

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
