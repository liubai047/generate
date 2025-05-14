package RpStatusEnum

import (
	"encoding/json"
	"fmt"
	"testing"
)

type SA struct {
	A  string `json:"a"`
	B  string `json:"b"`
	C  string `json:"c"`
	En *Enum  `json:"en"`
}

type SB struct {
	RA string `json:"ra"`
	RB *Enum  `json:"rb"`
}

func TestStructWithEnum2(t *testing.T) {
	// 测试结构体序列化
	req := &SA{}
	err := json.Unmarshal([]byte(`{"a":"a","b":"b","c":"c","en":"wt01"}`), &req)
	if err != nil {
		t.Errorf("json.Unmarshal() error = %v", err)
		return
	}
	t.Logf("%#v", req.En)

	var sb = SB{
		RA: "ra",
		RB: Success,
	}
	strSb, err := json.Marshal(sb)
	fmt.Printf("%#v || %s\n", strSb, string(strSb))
}
