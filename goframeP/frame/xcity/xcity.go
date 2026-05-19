package xcity

// 直辖市及特别行政区的行政区划代码
func ToSwitchCityCode(cityCode string) string {
	switch cityCode {
	case "110000":
		return "110100"
	case "310000":
		return "310100"
	case "120000":
		return "120100"
	case "500000":
		return "500100"
	case "810000":
		return "810100"
	}
	return cityCode
}
