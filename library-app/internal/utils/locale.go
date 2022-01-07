package utils

import (
	"encoding/json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Local struct {
	Bundle *i18n.Bundle
}

type Replace map[string]string

func NewLocale(bundle *i18n.Bundle) *Local {
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("i18n/en.json")
	bundle.LoadMessageFile("i18n/ru.json")

	return &Local{bundle}
}

func (l *Local) GetTranslate(lang, alias string) string {
	localizer := i18n.NewLocalizer(l.Bundle, lang)

	localizeConfigWelcome := i18n.LocalizeConfig{
		MessageID: alias,
	}
	translate, _ := localizer.Localize(&localizeConfigWelcome)

	return translate
}

func (l *Local) GetTranslateWithReplace(lang, alias string, replace Replace) string {
	localizer := i18n.NewLocalizer(l.Bundle, lang)

	localizeConfigWelcome := i18n.LocalizeConfig{
		MessageID:    alias,
		TemplateData: replace,
		PluralCount:  1,
	}
	translate, _ := localizer.Localize(&localizeConfigWelcome)

	return translate
}
