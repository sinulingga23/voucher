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
	ErrBrandEmpty = fmt.Errorf("Brands is empty.")
	
	// For Voucher service.
	ErrVoucherNotExists = fmt.Errorf("The Voucher is not exists.")
	ErrVoucherIsNotEnough = fmt.Errorf("The Voucher is not enough.")
	ErrVoucherNameIsEmpty = fmt.Errorf("Make sure the name of voucher is not empty.")
	ErrVoucherCostInInPointInvalid = fmt.Errorf("Make sure the cost in point of voucher is greater than 0.")
	ErrVoucherStockInvalid = fmt.Errorf("Make sure the stock of voucher is greater than 0.")
	ErrVoucherExpiratonDateInvalid = fmt.Errorf("Make sure the expiration date is valid.")
	ErrVoucherBrandIdNotExists = fmt.Errorf("Make sure the brand is exists.")
)