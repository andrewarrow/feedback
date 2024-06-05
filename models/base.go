package models

type BaseModel struct {
	Item map[string]any
}

func (b *BaseModel) GetInt(name string) int64 {
	v, _ := b.Item[name].(int64)
	return v
}
func (b *BaseModel) GetString(name string) string {
	v, _ := b.Item[name].(string)
	return v
}
