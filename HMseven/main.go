package main

import (
"log"
"reflect"
)

func PrintStruct(in interface{}){
val := reflect.ValueOf(in)
if val.Kind() == reflect.Ptr {
val = val.Elem()
}

if val.Kind() != reflect.Struct {
return
}

for i := 0; i < val.NumField(); i++ {
typeField := val.Type().Field(i)

if typeField.Type.Kind() == reflect.Struct {
log.Printf("nested field:%v", typeField.Name)
PrintStruct(val.Field(i).Interface())
continue
}
log.Printf("name=%s, type=%s, value=%v, tag = `%s`\n",
typeField.Name,
typeField.Type,
val.Field(i),
typeField.Tag,
)
}
}

func main(){


v := struct{
FieldString string `json:"field_string"`
FieldInt   int
Slice []int
Key  map[string]interface{}
}{
FieldString: "qwerty",
FieldInt: 42,
Key: [key] {1,2,3},
}
PrintStruct(v)
}
