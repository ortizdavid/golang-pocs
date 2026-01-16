package i18n

import (
	"encoding/json"
	"os"
	"sync"
)

var (
	translations = make(map[string]map[string]string)
	once         sync.Once
)

func LoadTranslations() {
	files := []string{"en", "pt"}
	for _, lang := range files {
		data, _ := os.ReadFile("i18n/locales/" + lang + ".json")
		var msgs map[string]string
		json.Unmarshal(data, &msgs)
		translations[lang] = msgs
	}
}
