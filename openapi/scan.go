package openapi

type OpenAPI struct {
}

func ScanDir(dir string) *OpenAPI {
	oa := OpenAPI{}
	return &oa
}
