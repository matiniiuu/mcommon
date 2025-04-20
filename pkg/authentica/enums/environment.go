package enums

type Environment string

const (
	MockServer     Environment = "mock_server"
	DebuggingProxy Environment = "debugging_proxy"
	Production     Environment = "production"
)

func (e Environment) GetURL() string {
	switch e {
	case MockServer:
		return "https://private-anon-0693cbcce7-authenticasa.apiary-mock.com/api/v1"
	case DebuggingProxy:
		return "https://private-anon-0693cbcce7-authenticasa.apiary-proxy.com/api/v1"
	case Production:
		return "https://api.authentica.sa/api/v1"
	default:
		return "https://api.authentica.sa/api/v1"
	}
}
