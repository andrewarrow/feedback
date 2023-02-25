package persist

type Database interface {
	exec(string) []map[string]string
}

type InMemory struct {
}

func NewInMemory() *InMemory {
	i := InMemory{}
	return &i
}

func (i *InMemory) exec(sql string) []map[string]string {
	m := map[string]string{}
	m["foo1"] = "a123"
	m["bar1"] = "a456"

	list := []map[string]string{}
	list = append(list, m)

	m = map[string]string{}
	m["foo2"] = "b123"
	m["bar2"] = "b456"
	list = append(list, m)

	return list
}
