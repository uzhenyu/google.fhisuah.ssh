package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"zg5/lmonth/framework/vipers"
)

var client config_client.IConfigClient

func getClient(fileName string) error {
	var err error
	err = vipers.GetViper(fileName)
	if err != nil {
		return err
	}
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(viper.GetString("Nacos.Host"), uint64(viper.GetInt("Nacos.Port")), constant.WithContextPath("/nacos")),
	}
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(""), //当namespace是public时，此处填空字符串。
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)
	client, err = clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func GetConfig(fileName string) (string, error) {
	if client == nil {
		err := getClient(fileName)
		if err != nil {
			return "", err
		}
	}
	err := vipers.GetViper(fileName)
	if err != nil {
		return "", err
	}
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: viper.GetString("Nacos.DataId"),
		Group:  viper.GetString("Nacos.Group"),
	})
	if err != nil {
		return "", err
	}
	return content, nil
}
