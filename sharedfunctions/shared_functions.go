package sharedfunctions

import (
	"encoding/json"
)

func GetMap(data map[string]any, key string) map[string]any {
	if val, ok := data[key].(map[string]any); ok {
		return val
	}
	return map[string]any{}
}

func GetList(data map[string]any, key string) []map[string]any {
	if val, ok := data[key]; ok {
		if list, ok := val.([]map[string]any); ok {
			return list
		}
	}
	return []map[string]any{}
}

func GetListString(data map[string]any, key string) []string {
	if val, ok := data[key]; ok {
		if interfaceList, ok := val.([]any); ok {
			strList := make([]string, 0)
			for _, v := range interfaceList {
				if str, ok := v.(string); ok {
					strList = append(strList, str)
				}
			}
			return strList
		}
	}
	return []string{}
}

func GetListAny(data map[string]any, key string) []map[string]any {
	if val, ok := data[key]; ok {
		if list, ok := val.([]any); ok {
			result := make([]map[string]any, 0, len(list))
			for _, item := range list {
				if itemMap, ok := item.(map[string]any); ok {
					result = append(result, itemMap)
				}
			}
			return result
		}
	}
	return []map[string]any{}
}

func ConvertStringToJSONMap(data map[string]any) {
	for key, value := range data {
		if str, ok := value.(string); ok {
			var jsonObject map[string]any
			if err := json.Unmarshal([]byte(str), &jsonObject); err == nil {
				data[key] = jsonObject
				continue
			}
			var jsonArray []map[string]any
			if err := json.Unmarshal([]byte(str), &jsonArray); err == nil {
				data[key] = jsonArray
			}
		}
	}
}

func ConvertStringToJSONList(data []map[string]any) {
	for i := range data {
		for key, value := range data[i] {
			if str, ok := value.(string); ok {
				var jsonValue map[string]any
				if err := json.Unmarshal([]byte(str), &jsonValue); err == nil {
					data[i][key] = jsonValue
				}
			}
			if str, ok := value.(string); ok {
				var jsonValue []map[string]any
				if err := json.Unmarshal([]byte(str), &jsonValue); err == nil {
					data[i][key] = jsonValue
				}
			}
		}
	}

}

func GetMapAtListAny(data []any, index int) map[string]any {
	if index < 0 || index >= len(data) {
		return map[string]any{}
	}

	if result, ok := data[index].(map[string]any); ok {
		return result
	}

	return map[string]any{}
}
