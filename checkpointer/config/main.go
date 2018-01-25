package config

type configuration struct {
	SecretKeyPath string
	ControllerIP  string
}

var (
	config configuration
)

//GetSecretKeyPath is getter for 秘密鍵のパス
func GetSecretKeyPath() string {
	return config.SecretKeyPath
}

//SetSecretKeyPath is setter for 秘密鍵のパス
func SetSecretKeyPath(path string) {
	config.SecretKeyPath = path
}

//GetControllerIP is getter for コントローラのIP
func GetControllerIP() string {
	return config.SecretKeyPath
}

//SetControllerIP is setter for コントローラのIP
func SetControllerIP(ip string) {
	config.ControllerIP = ip
}
