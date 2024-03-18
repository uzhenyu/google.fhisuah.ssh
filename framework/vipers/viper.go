package vipers

import "github.com/spf13/viper"

func GetViper(fileName string) error {
	viper.SetConfigFile(fileName)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
