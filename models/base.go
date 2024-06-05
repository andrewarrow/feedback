package models

type BaseModel struct {
	item map[string]any
}

func (b *BaseModel) GetInt(name string) int64 {
	v, _ := b.item[name].(int64)
	return v
}
func (b *BaseModel) GetString(name string) string {
	v, _ := b.item[name].(string)
	return v
}
