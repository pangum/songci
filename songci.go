package songci

import (
	"github.com/goexl/songci"
	"github.com/pangum/pangu"
)

// Songci 入口封装，由于语言特性，最好是直接继承的方式，通过显示使用主入口来执行初始化流程
type Songci struct {
	*songci.Songci
}

func newSongci(config *pangu.Config) (songci *Songci, err error) {
	_config := new(panguConfig)
	if err = config.Load(_config); nil != err {
		return
	}

	songci = new(Songci)
	songci.Songci = buildSongci(_config.Songci)

	return
}

func buildSongci(config config) *songci.Songci {
	builder := songci.New()
	if 0 != config.Timeout {
		builder.Timeout(config.Timeout)
	}
	if nil != config.Zinan {
		builder.Zinan().Scheme(config.Zinan.Scheme)
	}

	return builder.Build()
}
