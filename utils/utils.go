package utils

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// NumCheck 数値かどうかチェック
func NumCheck(value any) (rs bool) {
	if _, ok := value.(int); ok {
		return true
	} else if _, ok := value.(float64); ok {
		return true
	} else if _, ok := value.(float32); ok {
		return true
	}
	return false
}

// FileExistCheck ファイルが存在するかどうかチェック
func FileExistCheck(filename string) bool {
	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}

	return true
}

// NumBorderCheck 数値が指定した値を超えているかのチェック
func NumBorderCheck(values interface{}, lineNum float64) bool {
	var array []float64
	switch v := values.(type) {
	case []float64:
		array = v
	case map[string]float64:
		for _, value := range v {
			array = append(array, value)
		}
	case []int:
		for _, value := range v {
			array = append(array, float64(value))
		}
	case map[string]int:
		for _, value := range v {
			array = append(array, float64(value))
		}
	default:
		Error(fmt.Sprintf("type is not supported: %v", reflect.TypeOf(values)), 1)
		return false
	}
	var sum float64
	for _, value := range array {
		sum += value
	}

	if sum >= lineNum {
		return true
	}
	return false
}

func EncodeJSON(data interface{}) ([]byte, error) {
	if data == nil {
		return nil, nil
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to encode JSON: %v", err)
	}
	return jsonData, nil
}

type stringer interface {
	String() string
}

func JoinArr[T stringer](arr []T, sep string) string {
	var strArr []string
	for i, v := range arr {
		strArr[i] = v.String()
	}
	return strings.Join(strArr, sep)
}

func UnmarshalJSON(data []byte, v interface{}) error {
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("failed to decode JSON: %v", err)
	}
	return nil
}

func Contains(list interface{}, elem interface{}) bool {
	listV := reflect.ValueOf(list)

	if listV.Kind() == reflect.Slice {
		for i := 0; i < listV.Len(); i++ {
			item := listV.Index(i).Interface()
			if !reflect.TypeOf(elem).ConvertibleTo(reflect.TypeOf(item)) {
				continue
			}
			target := reflect.ValueOf(elem).Convert(reflect.TypeOf(item)).Interface()
			if ok := reflect.DeepEqual(item, target); ok {
				return true
			}
		}
	}
	return false
}

func SortMap(arr []map[string]string, targetColumn, sortType string) []map[string]string {
	sort.Slice(arr, func(i, j int) bool {
		valI, _ := strconv.Atoi(arr[i][targetColumn])
		valJ, _ := strconv.Atoi(arr[j][targetColumn])
		if sortType == "asc" {
			return valI < valJ
		} else {
			return valI > valJ
		}
	})
	return arr
}

func MakeRandomStr(digit uint) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	var result string
	for _, v := range b {
		result += string(letters[v%byte(len(letters))])
	}
	return result, nil
}
