package utils

import (
	"encoding/json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Local struct {
	Bundle *i18n.Bundle
}

type Template map[string]string

func NewLocale(bundle *i18n.Bundle) *Local {
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("i18n/en.json")
	bundle.LoadMessageFile("i18n/ru.json")

	return &Local{bundle}
}

func (l *Local) GetTranslate(lang, alias string, template Template) string {
	localizer := i18n.NewLocalizer(l.Bundle, lang)

	localizeConfigWelcome := i18n.LocalizeConfig{
		MessageID:    alias,
		TemplateData: template,
		PluralCount:  1,
	}
	translate, _ := localizer.Localize(&localizeConfigWelcome)

	return translate
}
