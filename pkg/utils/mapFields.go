package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

func MapFields(oldRecord interface{}, newRecord interface{}) ([]byte, error) {
	var result map[string]interface{}
	oldRecordJson, _ := json.Marshal(oldRecord)
	json.Unmarshal(oldRecordJson, &result)

	var newRecordMap map[string]interface{}
	newRecordJson, _ := json.Marshal(newRecord)
	json.Unmarshal(newRecordJson, &newRecordMap)

	fmt.Println("MapFields")
	for key, newValue := range newRecordMap {
		switch c := newValue.(type) {
		case nil:
			fmt.Println("nil")
		case bool:
			if newRecordMap[key] != result[key] {
				result[key] = newRecordMap[key]
			}
		case string, int, int8, int16, int32, int64, float32, float64:
			if newRecordMap[key] != "" && newRecordMap[key] != result[key] {
				result[key] = newRecordMap[key]
			}
		case map[string]interface{}:
			//-- Check map is a sql.NullTime
			if _, isTime := c["Time"]; isTime {
				if _, isSqlNullTime := c["Valid"]; isSqlNullTime {
					if c["Valid"].(bool) && newRecordMap[key].(map[string]interface{})["Time"] != result[key].(map[string]interface{})["Time"] {
						result[key] = newRecordMap[key]
					}
				}
			}
		case time.Time:
		case *time.Time:
			fmt.Println("time.Time")
			fmt.Println("*time.Time")
			// //-- Check map is a sql.NullTime
			// if _, isTime := c["Time"]; isTime {
			// 	if _, isSqlNullTime := c["Valid"]; isSqlNullTime {
			// 		if c["Valid"].(bool) && newRecordMap[key].(map[string]interface{})["Time"] != oldRecordMap[key].(map[string]interface{})["Time"] {
			// 			oldRecordMap[key] = newRecordMap[key]
			// 		}
			// 	}
			// }
		default:
			return nil, errors.New(fmt.Sprintf("There is no map for key %v :: Value %v", key, c))
		}
	}
	fmt.Println("MapFields End.")
	return json.Marshal(result)
}
