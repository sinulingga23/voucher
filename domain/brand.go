package domain

type (
	Brand struct {
		Id string
		Name string
		UrlLogo string
		Description string
		Address string
	}

	CreateBrand struct {
		Name string
		UrlLogo string
		Description string
		Address string
	}
)