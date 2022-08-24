package config

var (
	Host            string
	Port            string
	User            string
	Password        string
	DBOption        string
	Appids          []string
	TableName       string
	AppidMap        map[string]struct{}
	InteractionMode bool

	IdcMap = map[string]string{
		"120.92.181.73":  "sh",
		"120.92.181.74":  "sh",
		"120.92.176.21":  "sh",
		"120.92.176.50":  "sh",
		"120.92.181.111": "sh",
		"120.92.181.137": "sh",
		"120.92.181.98":  "sh",
		"120.92.181.104": "sh",
		"120.92.181.97":  "sh",
		"120.92.181.96":  "sh",
		"120.92.181.71":  "sh",
		"120.92.181.72":  "sh",
		"120.92.181.100": "sh",
		"120.92.181.101": "sh",
		"120.92.181.102": "sh",
		"120.92.181.103": "sh",

		"172.18.4.54": "gz",
		"172.18.4.55": "gz",
		"172.18.4.56": "gz",
		"172.18.4.57": "gz",
		"172.18.4.58": "gz",
		"172.18.4.59": "gz",
		"172.18.4.33": "gz",
		"172.18.4.34": "gz",
		"172.18.4.35": "gz",
		"172.18.4.36": "gz",

		"10.11.98.24":    "zh",
		"10.11.98.61":    "zh",
		"43.156.122.8":   "sg",
		"43.156.123.152": "sg",
	}
)

func handleAppidMap() {
	AppidMap = make(map[string]struct{})
	for _, appid := range Appids {
		AppidMap[appid] = struct{}{}
	}
}

func InitConfig() {
	handleAppidMap()
}
