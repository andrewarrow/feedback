package common

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"html/template"
	"strconv"
	"strings"
	"time"

	"github.com/andrewarrow/feedback/util"
)

const HUMAN_SMALL = "01/02/2006"

func TemplateFunctions() template.FuncMap {
	fm := template.FuncMap{
		"uuid": func() string {
			return PseudoUuid()
		},
		"short": func(a any) string {
			tInt, ok := a.(int64)
			if ok {
				t := time.Unix(tInt, 0)
				return t.Format(HUMAN_SMALL)
			}
			return ""
		},
		"add": func(i, j int) int { return i + j },
		"mul": func(i, j int64) int64 { return i * j },
		"optionList": func(s string, val any) template.HTML {
			valString := ""
			if val != nil {
				valString = val.(string)
			}
			tokens := strings.Split(s[2:len(s)-2], "|")
			valTokens := strings.Split(valString, ",")
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
		"textfield": func(name string, val any) template.HTML {
			value := ""
			if val != nil {
				value = fmt.Sprintf("%v", val)
			}
			s := fmt.Sprintf(`<input id="textfield-%s" class="nice-i border" name="%s" type="text" value="%s"/>`, name, name, value)
			return template.HTML(s)
		},
		"breakComma": func(s string) []string { return strings.Split(s, ",") },
		"price": func(s any) template.HTML {
			amount := fmt.Sprintf("%v", s)
			price, _ := strconv.Atoi(amount)
			if price == 0 {
				sp := "$0.00"
				return template.HTML(sp)
			}

			if len(amount) < 3 {
				sp := fmt.Sprintf("$00.%s", amount)
				return template.HTML(sp)
			}
			dollars := amount[0 : len(amount)-2]
			dollarsInt, _ := strconv.ParseInt(dollars, 10, 64)
			dollars = util.IntComma(dollarsInt)
			cents := amount[len(amount)-2:]
			sp := fmt.Sprintf("$%s.%s", dollars, cents)
			if price < 0 {
				sp = "<span class='text-red-500'>" + sp + "</span>"
			}
			return template.HTML(sp)
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
			//dollarsInt, _ := strconv.ParseInt(dollars, 10, 64)
			//dollars = util.IntComma(dollarsInt)
			cents := amount[len(amount)-2:]
			return fmt.Sprintf("%s.%s", dollars, cents)
		},
	}
	return fm
}

func PseudoUuid() string {

	b := make([]byte, 16)
	rand.Read(b)

	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
