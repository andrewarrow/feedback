package util

func ToAnyArray(rows []map[string]any) []any {
	list := []any{}
	for _, row := range rows {
		list = append(list, row)
	}
	return list
}

func ToAny(rows []string) []any {
	list := []any{}
	for _, row := range rows {
		list = append(list, row)
	}
	return list
}

func ToMSA(rows []any) []map[string]any {
	sizes := []map[string]any{}
	for _, item := range rows {
		sizes = append(sizes, item.(map[string]any))
	}
	return sizes
}
