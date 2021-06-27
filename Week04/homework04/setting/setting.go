package setting

import (
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}


// NewSetting 读取普通的配置
func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")  // config.yaml
	vp.AddConfigPath("./configs")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}