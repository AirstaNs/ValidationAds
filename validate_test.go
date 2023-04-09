package ValidationAds

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestBadValidateTitle(t *testing.T) {
	err := ValidateTitle("")
	if is := assert.ErrorIs(t, err, ErrBadTitle); !is {
		t.Errorf("should be bad title")
	}
}
func TestBadValidateTitleLen(t *testing.T) {
	title := strings.Repeat("a", 101)
	err := ValidateTitle(title)
	if is := assert.ErrorIs(t, err, ErrBadTitle); !is {
		t.Errorf("should be bad title")
	}

}
func TestGoodValidateTitle(t *testing.T) {
	err := ValidateTitle("1")

	if is := assert.ErrorIs(t, err, nil); !is {
		t.Errorf("title  should be passes the test")
	}
}

func TestBadValidateTextEmpty(t *testing.T) {
	err := ValidateText("")
	if is := assert.ErrorIs(t, err, ErrBadText); !is {
		t.Errorf("should be bad text")
	}
}
func TestBadValidateTextLen(t *testing.T) {
	text := strings.Repeat("a", 501)
	err := ValidateText(text)
	if is := assert.ErrorIs(t, err, ErrBadText); !is {
		t.Errorf("should be bad text")
	}
}
func TestGoodValidateText(t *testing.T) {
	err := ValidateText("1")
	if is := assert.ErrorIs(t, err, nil); !is {
		t.Errorf("text should be passes the test")
	}
}

func TestBadValidateAuthorID(t *testing.T) {
	err := ValidateAuthorID(1, 2)
	if is := assert.ErrorIs(t, err, ErrBadAuthorID); !is {
		t.Errorf("should be bad author id")
	}
}
func TestGoodValidateAuthorID(t *testing.T) {
	err := ValidateAuthorID(1, 1)
	if is := assert.ErrorIs(t, err, nil); !is {
		t.Errorf("author id should be passes the test")
	}
}
