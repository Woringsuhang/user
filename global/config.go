package global

var ConfigAll Config

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
	} `json:"register"`
}
