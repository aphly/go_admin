package helper

import (
	"go_admin/app/core"
	"math"
	"sync"
	"time"
)

var NewSnowflake *Snowflake

func init() {
	epoch := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC).UnixNano() / 1e6
	var nodeId int64 = 1
	if nodeId < 0 || nodeId > 1023 {

	}
	NewSnowflake = &Snowflake{
		epoch:        epoch,
		nodeBits:     10,
		nodeMax:      int64(math.Pow(2, float64(10)) - 1),
		nodeId:       nodeId,
		sequenceBits: 12, // 12 bits for sequence
		sequenceMax:  int64(math.Pow(2, float64(12)) - 1),
	}
}

type Snowflake struct {
	epoch         int64
	nodeBits      uint
	nodeMax       int64
	nodeId        int64
	sequenceBits  uint
	sequenceMax   int64
	lastTimestamp int64
	mutex         sync.Mutex
	sequence      int64
}

func (s *Snowflake) NextId() core.Uint {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	currentTimestamp := time.Now().UnixNano() / 1e6
	if s.lastTimestamp == currentTimestamp {
		s.sequence++
		if s.sequence > s.sequenceMax {
			//for currentTimestamp <= s.lastTimestamp {
			//	currentTimestamp = time.Now().UnixNano() / 1e6
			//}
			time.Sleep(time.Millisecond)
			currentTimestamp = time.Now().UnixNano() / 1e6
			s.sequence = 0
		}
	} else {
		s.sequence = 0
	}
	s.lastTimestamp = currentTimestamp
	return core.Uint(((currentTimestamp - s.epoch) << (s.nodeBits + s.sequenceBits)) | (s.nodeId << s.sequenceBits) | s.sequence)
}
