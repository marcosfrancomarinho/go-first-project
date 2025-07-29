package valuesobject

import "errors"

type Page struct {
	page int
}

func NewPage(page int) (*Page, error) {
	if err := validatePage(page); err != nil {
		return nil, err
	}
	return &Page{page}, nil
}

func validatePage(page int) error {
	if page <= 0 {
		return errors.New("numero de pagina não pode ser 0 ou menor do que zero")
	}
	return nil
}

func (p *Page) GetValue() int {
	return p.page
}
