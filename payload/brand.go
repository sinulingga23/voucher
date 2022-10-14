package payload

type (
	Brand struct {
		Id string `json:"id"`
		Name string `json:"name"`
		UrlLogo string `json:"urlLogo"`
		Description string `json:"description"`
		Address string `json:"address"`
	}

	CreateBrand struct {
		Name string `json:"name"`
		UrlLogo string `json:"urlLogo"`
		Description string `json:"description"`
		Address string `json:"address"`
	}

)