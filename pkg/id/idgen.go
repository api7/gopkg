// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
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
	"encoding/json"
	"fmt"
	"hash/crc32"
	"reflect"
	"strconv"
	"unsafe"
)

// GenID generates an ID according to the raw material.
func GenID(raw string) string {
	if raw == "" {
		return ""
	}
	sh := &reflect.SliceHeader{
		Data: (*reflect.StringHeader)(unsafe.Pointer(&raw)).Data,
		Len:  len(raw),
		Cap:  len(raw),
	}
	p := *(*[]byte)(unsafe.Pointer(sh))

	res := crc32.ChecksumIEEE(p)
	return fmt.Sprintf("%x", res)
}

// ID is the type of the id field used for any entities
type ID uint64

// String indicates how to convert ID to a string.
func (id ID) String() string {
	return strconv.FormatUint(uint64(id), 10)
}

// MarshalJSON is the way to encode ID to JSON string.
func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.FormatUint(uint64(id), 10))
}

// UnmarshalJSON is the way to decode ID from JSON string.
func (id *ID) UnmarshalJSON(data []byte) error {
	var value interface{}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	switch v := value.(type) {
	case string:
		u, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return err
		}
		*id = ID(u)
	default:
		panic("unknown type")
	}
	return nil
}

// IDGenerator is an interface for generating IDs.
type IDGenerator interface {
	// NextID generates an ID.
	NextID() ID
}
