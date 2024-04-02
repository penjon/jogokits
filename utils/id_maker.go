package utils

import (
	"errors"
	"fmt"
)

type IDMaker struct {
	workerID           int64
	twepoch            int64
	sequence           int64
	workerIdBits       int64
	maxWorkerID        int64
	sequenceBits       int64
	workerIDShift      int64
	timeStampLeftShift int64
	sequenceMask       int64
	lastTimeStamp      int64
}

/*func NewMaker(workerID int64) *IDMaker {

	maker := &IDMaker{
		workerID:           workerID,
		twepoch:            1288834974657,
		sequence:           0,
		workerIdBits:       4,
		maxWorkerID:        -1 ^ -1<<4,
		sequenceBits:       10,
		workerIDShift:      10,
		timeStampLeftShift: 14,
		sequenceMask:       -1 ^ -1<<10,
		lastTimeStamp:      -1,
	}

	return maker
}*/

func (i *IDMaker) NextID() (int64, error) {
	timestamp := GetTimeMillis()
	if i.lastTimeStamp == timestamp {
		i.sequence = (i.sequence + 1) & i.sequenceMask
		if i.sequence == 0 {
			timestamp = tilNextMillis(i.lastTimeStamp)
		}
	} else {
		i.sequence = 0
	}

	if timestamp < i.lastTimeStamp {
		return 0, errors.New(fmt.Sprintf("Clock moved backwards.  Refusing to generate id for %d milliseconds", i.lastTimeStamp-timestamp))
	}
	i.lastTimeStamp = timestamp
	return (timestamp - i.twepoch<<i.timeStampLeftShift) | (i.workerID << i.workerIDShift) | (i.sequence), nil
}

func tilNextMillis(lastTimeStamp int64) int64 {
	timestamp := GetTimeMillis()
	for {
		if timestamp > lastTimeStamp {
			return timestamp
		}
		timestamp = GetTimeMillis()
	}
}
