package snowflake

import (
	"fmt"
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func Init(startTime string, MachineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-01", startTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(MachineID)
	return
}

func GenID() int64 {
	return node.Generate().Int64()
}

func main() {
	if err := Init("2020-08-09", 1); err != nil {
		fmt.Printf("Init Failed,err:%v\n", err)
	}
	id := GenID()
	fmt.Println(id)
}
