package models

import (
	"strings"
	"time"
)

type BaseModel struct {
	Item map[string]any
}

func NewBase(item map[string]any) *BaseModel {
	b := BaseModel{}
	b.Item = item
	return &b
}

func (b *BaseModel) GetBytes(name string) []byte {
	v, _ := b.Item[name].(string)
	return []byte(v)
}
func (b *BaseModel) GetFloat(name string) float64 {
	v, _ := b.Item[name].(float64)
	return v
}
func (b *BaseModel) GetInt(name string) int64 {
	v, _ := b.Item[name].(int64)
	return v
}
func (b *BaseModel) GetString(name string) string {
	v, _ := b.Item[name].(string)
	return strings.TrimSpace(v)
}
func (b *BaseModel) GetStringOk(name string) (string, bool) {
	if b.Item[name] == nil {
		return "", false
	}
	v, _ := b.Item[name].(string)
	return v, true
}
func (b *BaseModel) GetMap(name string) map[string]any {
	v, _ := b.Item[name].(map[string]any)
	return v
}
func (b *BaseModel) GetBool(name string) (bool, bool) {
	if b.Item[name] == nil {
		return false, false
	}
	v := b.Item[name].(bool)
	if v == false {
		return false, true
	}
	return true, true
}
func (b *BaseModel) GetSimpleBool(name string) bool {
	v, _ := b.Item[name].(bool)
	return v
}
func (b *BaseModel) GetList(name string) []any {
	v, _ := b.Item[name].([]any)
	return v
}
func (b *BaseModel) GetTime(name string) time.Time {
	v, ok := b.Item[name].(time.Time)
	if ok {
		return v
	}
	return time.Now()
}
