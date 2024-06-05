package models

type BaseModel struct {
	Item map[string]any
}

func NewBase(item map[string]any) *BaseModel {
	b := BaseModel{}
	b.Item = item
	return &b
}

func (b *BaseModel) GetInt(name string) int64 {
	v, _ := b.Item[name].(int64)
	return v
}
func (b *BaseModel) GetString(name string) string {
	v, _ := b.Item[name].(string)
	return v
}
func (b *BaseModel) GetMap(name string) map[string]any {
	v, _ := b.Item[name].(map[string]any)
	return v
}
