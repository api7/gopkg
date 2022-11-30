// Copyright 2022 API7.ai, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package id

import (
	"github.com/pkg/errors"
	"github.com/sony/sonyflake"
)

type snowflake sonyflake.Sonyflake

// NextID generates an ID.
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
