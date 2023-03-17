package models

func FindField(model *Model, id string) *Field {
	for _, f := range model.Fields {
		if f.Name == id {
			return f
		}
	}

	return nil
}
