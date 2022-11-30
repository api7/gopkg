package id

import (
	"github.com/pkg/errors"
	"github.com/sony/sonyflake"
)

type snowflake sonyflake.Sonyflake

func (s *snowflake) NextID() ID {
	uid, err := (*sonyflake.Sonyflake)(s).NextID()
	if err != nil {
		panic("get sony flake uid failed:" + err.Error())
	}
	return ID(uid)
}

// NewIDGenerator returns an IDGenerator object.
func NewIDGenerator() (IDGenerator, error) {
	ips, err := getLocalIPs()
	if err != nil {
		panic(err)
	}
	sf := (*snowflake)(sonyflake.NewSonyflake(sonyflake.Settings{
		MachineID: func() (u uint16, e error) {
			return sumIPs(ips), nil
		},
	}))
	if sf == nil {
		return nil, errors.New("failed to new snoyflake object")
	}
	return sf, nil
}
