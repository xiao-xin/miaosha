package common

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// 根据结构体中的sql标签，把数据转换成结构体
func DataToStructByTagSql(data map[string]string ,obj interface{}){
	  objValue :=reflect.ValueOf(obj).Elem()
	  for i := 0 ; i<objValue.NumField() ; i++ {
	  	// 获取结构体sql标签名在data中的值
	  	value,ok := data[objValue.Type().Field(i).Tag.Get("sql")]
	  	if !ok {
	  		continue
		}
	  	// 获取结构体i字段的名称
	  	name := objValue.Type().Field(i).Name
	  	// 获取结构体第i个字段的类型
	  	structFieldType := objValue.Field(i).Type()
	  	// 其实value的类型就是string
	  	val:= reflect.ValueOf(value)
	  	var err error
	  	if structFieldType != val.Type() {
			val,err = TypeConversion(value,structFieldType.Name())
			if err != nil {
				fmt.Println(err,val.Type())
				break;
			}
		}
		//index :=[1]int{1}
		objValue.FieldByName(name).Set(val)
 	  }
}

// 类型转换，val值转换为toType类型
func TypeConversion(val string ,toType string ) (reflect.Value,error){
	if toType == "string " {
		return  reflect.ValueOf(val),nil
	}else if  toType =="time.Time"{
		t, err := time.ParseInLocation("2006-01-02 15:04:05",val,time.Local)
		return reflect.ValueOf(t),err
	} else if toType =="Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05",val,time.Local)
		return reflect.ValueOf(t),err
	} else if toType =="int" {
		v, err := strconv.Atoi(val)
		return reflect.ValueOf(v),err
	}else if toType =="int8" {
		v, err := strconv.ParseInt(val,10,64)
		return reflect.ValueOf(int8(v)),err
	} else if toType =="int32" {
		t, err := strconv.ParseInt(val,10,64)
		return reflect.ValueOf(int32(t)),err
	} else if toType =="int64" {
		t, err := strconv.ParseInt(val,10,64)
		return reflect.ValueOf(t),err
	} else if toType == "float32" {
		i, err := strconv.ParseFloat(val,64)
		return  reflect.ValueOf(float32(i)),err
	}else if toType == "float64" {
		i, err := strconv.ParseFloat(val,64)
		return  reflect.ValueOf(i),err
	}

	return reflect.ValueOf(val),errors.New("未知类型:"+toType)
}
