/*
 * 版权属于：yitter(yitter@126.com)
 * 代码编辑：guoyahao
 * 代码修订：yitter
 * 开源地址：https://github.com/yitter/idgenerator
 */

package idgen

import (
	"fmt"
	"strconv"
)

type SnowWorkerM2 struct {
	*SnowWorkerM1
}

func NewSnowWorkerM2(options *IdGeneratorOptions) ISnowWorker {
	return &SnowWorkerM2{
		NewSnowWorkerM1(options).(*SnowWorkerM1),
	}
}

func (m2 SnowWorkerM2) NextId() int64 {
	m2.Lock()
	defer m2.Unlock()
	currentTimeTick := m2.GetCurrentTimeTick()
	if m2._LastTimeTick == currentTimeTick {
		m2._CurrentSeqNumber++
		if m2._CurrentSeqNumber > m2.MaxSeqNumber {
			m2._CurrentSeqNumber = m2.MinSeqNumber
			currentTimeTick = m2.GetNextTimeTick()
		}
	} else {
		m2._CurrentSeqNumber = m2.MinSeqNumber
	}
	if currentTimeTick < m2._LastTimeTick {
		fmt.Println("Time error for {0} milliseconds", strconv.FormatInt(m2._LastTimeTick-currentTimeTick, 10))
	}
	m2._LastTimeTick = currentTimeTick
	result := int64(currentTimeTick<<m2._TimestampShift) + int64(m2.WorkerId<<m2.SeqBitLength) + int64(m2._CurrentSeqNumber)
	return result
}
