/**
  @Author: dongxx
  @Since: 2026/1/15 14:19
  @Desc: //TODO
**/

package config

import (
	"fmt"
	"github.com/spf13/cast"
	"os"
	"strings"
	"unitechs.com/unios-dice/uni-base/core/config"
	"unitechs.com/unios-dice/uni-base/core/log"
)

// ReplaceByEnv 将环境变量值替换原配置文件中的值
func ReplaceByEnv() {
	configKeys := make([]string, 0)
	configKeyValue := make(map[string]interface{})
	for group, kv := range config.AllSettings() {
		for k, v := range kv.(map[string]interface{}) {
			key := fmt.Sprintf("%s.%s", group, k)
			configKeys = append(configKeys, key)
			configKeyValue[key] = v
		}
	}

	for _, key := range configKeys {
		keyEnv := strings.Replace(strings.ToUpper(key), ".", "_", -1)
		v, ok := os.LookupEnv(keyEnv)
		log.Infof("keyEnv %s value %s", keyEnv, v)
		if ok {
			oldVal := configKeyValue[key]
			switch oldVal.(type) {
			case string:
				config.Set(keyEnv, v)
			case bool:
				config.Set(keyEnv, cast.ToBool(v))
			case int:
				config.Set(keyEnv, cast.ToInt(v))
			case int64:
				config.Set(keyEnv, cast.ToInt64(v))
			}
		}
	}
}
