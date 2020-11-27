package gm

import (
	"errors"

	"github.com/jxu86/fabric-sdk-go-gm/internal/github.com/hyperledger/fabric/bccsp"
)
//todo:国密 gosdk：增加gm
//定义国密 Key的驱动 ，实现 KeyDeriver 接口
type smPublicKeyKeyDeriver struct{}

func (kd *smPublicKeyKeyDeriver) KeyDeriv(k bccsp.Key, opts bccsp.KeyDerivOpts) (dk bccsp.Key, err error) {
	return nil, errors.New("Not implemented")
}
