package starter

import (
	"github.com/pangum/pangu"
)

// Starter TODO 根据第三方入口的名字来命名
type Starter struct {
	// TODO 从第三方要接入的主入口继承
}

func newAgent(config *pangu.Config) (starter *Starter, err error) {
	_config := new(panguConfig)
	if err = config.Load(_config); nil != err {
		return
	}

	return
}
