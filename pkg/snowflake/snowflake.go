package snowflake

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
)

var node *sf.Node

func Init(startTime string, machineId int64) (err error) {
	var st time.Time
	st,err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}

	sf.Epoch = st.UnixNano() / 1000000

	
	node, err = sf.NewNode(machineId)
	if err != nil {
		zap.L().Error("snowflake init failed", zap.Error(err))
		return
	}

	return
}

func GenerateId() (id int64) {
	return node.Generate().Int64()
}