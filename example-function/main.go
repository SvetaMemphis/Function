package main

import (
	"encoding/json"
	"fmt"

	"github.com/memphisdev/memphis.go"
	"strconv"
)

func CheckSeverity(message []byte, headers map[string]string, inputs map[string]string) ([]byte, map[string]string, error) {
	numInt := strconv.Atoi(inputs["num"])
	fmt.Println("input", 5/numInt)
	// var msg_map map[string]interface{} = make(map[string]interface{})

	// if err := json.Unmarshal(message, &msg_map); err != nil {
	// 	return nil, nil, err
	// }

	// last_produce_from_message_ms, ok := ((msg_map)["time_since_last_produce"]).(float64)
	// if !ok {
	// 	fmt.Println("last_produce_from_message is not an int")
	// 	err_str := "time_since_last_produce is not a valid key, or was not parsed correctly as a fload by json.Unmarshal"
	// 	return nil, nil, fmt.Errorf(err_str)
	// }
	// last_produce_from_message_ms_int := int(last_produce_from_message_ms)

	// if last_produce_from_message_ms_int >= 10000 {
	// 	msg_map["severity"] = "critical"
	// } else if last_produce_from_message_ms_int >= 5000 {
	// 	msg_map["severity"] = "high"
	// } else {
	// 	msg_map["severity"] = "low"
	// }

	// if msgStr, err := json.Marshal(msg_map); err != nil {
	// 	return nil, nil, err
	// } else {
	// 	return msgStr, inputs, nil
	// }
	// return fmt.Errof("test error")
	return message, headers, nil
}

func main() {
	fmt.Println("hello world")
	memphis.CreateFunction(CheckSeverity)
}
