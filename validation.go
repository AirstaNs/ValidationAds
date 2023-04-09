package ValidationAds

import "errors"

var ErrBadTitle = errors.New("bad title")
var ErrBadText = errors.New("bad text")
var ErrBadAuthorID = errors.New("bad author id")

func ValidateTitle(title string) error {
	if title == "" {
		return ErrBadTitle
	}
	if len(title) > 100 {
		return ErrBadTitle
	}
	return nil
}
func ValidateText(text string) error {
	if text == "" {
		return ErrBadText
	}
	if len(text) > 500 {
		return ErrBadText
	}
	return nil
}

func ValidateAuthorID(adAuthorID int64, authorID int64) error {
	if adAuthorID != authorID {
		return ErrBadAuthorID
	}
	return nil
}
