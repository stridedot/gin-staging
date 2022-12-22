package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"go_code/gintest/bootstrap"
	"go_code/gintest/bootstrap/glog"
	"time"
)

var node *snowflake.Node

func createNode(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}

	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineID)
	return
}

func GenID() int64 {
	err := createNode(bootstrap.Config.StartTime, bootstrap.Config.MachineId)
	if err != nil {
		glog.SL.Error("Failed to create snowflake node, ", err)
		return 0
	}

	return node.Generate().Int64()
}
