package utility

import (
	"fmt"
)

var (
	// For all services.
	ErrInternalError = fmt.Errorf("Sorry. Cant' process your request now.")

	// For Brand service.
	ErrBrandNameIsEmpty = fmt.Errorf("Make sure the name of brand is not empty.")
	ErrBrandLogoIsEmpty = fmt.Errorf("Make sure the url logo of brand is not empty.") 
)