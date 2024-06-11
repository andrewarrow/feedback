package models

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
	return v
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
func (b *BaseModel) GetBoolAsInt(name string) int {
	if b.Item[name] == nil {
		return 2
	}
	v := b.Item[name].(bool)
	if v == false {
		return 0
	}
	return 1
}
