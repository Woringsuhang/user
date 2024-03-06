package global

var (
	ConfigAll   *Config
	NacosConfig *ConfigNa
)

type Config struct {
	Mysql struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Library  string `json:"library"`
	} `json:"mysql"`
	Grpc struct {
		Agreement string `json:"agreement"`
		Port      string `json:"port"`
	} `json:"grpc"`
}
type ConfigNa struct {
	NamespaceId string `mapstructure:"namespaceId"`
	IpAddr      string `mapstructure:"ipAddr"`
	Port        int    `mapstructure:"port"`
	DataId      string `mapstructure:"dataId"`
	Group       string `mapstructure:"group"`
}
