package songci

import (
	"time"
)

type config struct {
	// 超时
	Timeout time.Duration `json:"timeout" yaml:"timeout" xml:"timeout" toml:"timeout"`
	// 子南算法
	Zinan *zinan `json:"zinan" yaml:"zinan" xml:"zinan" toml:"zinan"`
}
