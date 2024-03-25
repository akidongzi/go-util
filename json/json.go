package json

import "encoding/json"

func Encode(v interface{}) (string, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func Decode(data []byte, val interface{}) error {
	return json.Unmarshal(data, val)
}


// ObjToJSON ...
func ObjToJSON(msg interface{}) string {
	jso, error := json.Marshal(msg)
	if error == nil {
		return string(jso)
	}
	return ""
}
