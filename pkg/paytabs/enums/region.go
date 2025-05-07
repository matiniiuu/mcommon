package enums

type PaytabsRegion string

const (
	ARE    PaytabsRegion = "ARE"
	SAU    PaytabsRegion = "SAU"
	OMN    PaytabsRegion = "OMN"
	JOR    PaytabsRegion = "JOR"
	EGY    PaytabsRegion = "EGY"
	KWT    PaytabsRegion = "KWT"
	GLOBAL PaytabsRegion = "GLOBAL"
)

func (pg PaytabsRegion) GetURL() string {
	urls := map[PaytabsRegion]string{
		ARE:    "https://secure.paytabs.com",
		SAU:    "https://secure.paytabs.sa",
		OMN:    "https://secure-oman.paytabs.com",
		JOR:    "https://secure-jordan.paytabs.com",
		EGY:    "https://secure-egypt.paytabs.com",
		KWT:    "https://secure-kuwait.paytabs.com",
		GLOBAL: "https://secure-global.paytabs.com",
	}
	return urls[pg]
}
