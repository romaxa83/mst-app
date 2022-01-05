package utils

import (
	"encoding/json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Local struct {
	Bundle *i18n.Bundle
}

func NewLocale(bundle *i18n.Bundle) *Local {
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("i18n/en.json")
	bundle.LoadMessageFile("i18n/ru.json")

	return &Local{bundle}
}

func (l *Local) GetTranslate(alias, lang string) string {
	localizer := i18n.NewLocalizer(l.Bundle, lang)

	localizeConfigWelcome := i18n.LocalizeConfig{
		MessageID: alias,
	}
	translate, _ := localizer.Localize(&localizeConfigWelcome)

	return translate
}
