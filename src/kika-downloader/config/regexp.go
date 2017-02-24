package config

const (
	RegExpVideoId    = "^html5-p[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"
	RegExpXmlDataUrl = "^return BRavFramework.register.*.setup\\(\\{dataURL:'(http?://.*.xml)'\\}\\)\\);$"
)
