package hello

const (
	englishPrefix   = "Hello, "
	spanishPrefix   = "Hola, "
	frenchPrefix    = "Bonjour, "
	languageEnglish = "english"
	languageSpanish = "spanish"
	languageFrench  = "french"
)

func Hello(name, languageName string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(languageName) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case languageFrench:
		prefix = frenchPrefix
	case languageSpanish:
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}
	return
}

// func Hello(name, languageName string) string {
// 	if name == "" {
// 		name = "World"
// 	}
// 	if languageName == languageSpanish {
// 		return spanishPrefix + name
// 	}
// 	if languageName == languageFrench {
// 		return frenchPrefix + name
// 	}
//
// 	return englishPrefix + name
// }
