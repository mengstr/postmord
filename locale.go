package postmord

type Locale string

const (
	LocaleEnglish   Locale = "en"
	LocaleSwedish   Locale = "sv"
	LocaleDanish    Locale = "dk"
	LocaleNorwegian Locale = "no"
	LocaleFinnish   Locale = "fi"

	LocaleDefault = LocaleEnglish
)

func (l Locale) String() string { return string(l) }
