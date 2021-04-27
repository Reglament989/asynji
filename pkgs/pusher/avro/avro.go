package avro

import (
	"fmt"

	"github.com/linkedin/goavro/v2"
)

var (
	codec *goavro.Codec
)

func init() {
	//Create Schema Once
	var err error
	codec, err = goavro.NewCodec(`
	{
		"type": "record",
		"name": "Message",
		"fields" : [
			{"name": "Hello", "type": "string"}
		]
	}
	`)
	if err != nil {
		panic(err)
	}
}

func UnmarshalMessage(buf []byte) (*PushMessage, error) {
	native, _, err1 := codec.NativeFromTextual(buf)
	if err1 != nil {
		fmt.Printf("%v\n", err1)
		return nil, err1
	}

	return StringMapToPushMessage(native.(map[string]interface{})), nil
}

func (p *PushMessage) MarshalMessage() ([]byte, error) {
	binary, err := codec.TextualFromNative(nil, p.ToStringMap())
	if err != nil {
		return nil, err
	}
	return binary, nil
}

type PushMessage struct {
	Hello string
}

func (u *PushMessage) ToStringMap() map[string]interface{} {
	datumIn := map[string]interface{}{
		"Hello": string(u.Hello),
	}

	// if len(u.Errors) > 0 {
	// 	datumIn["Errors"] = goavro.Union("array", u.Errors)
	// } else {
	// 	datumIn["Errors"] = goavro.Union("null", nil)
	// }
	return datumIn
}

// StringMapToUser returns a User from a map representation of the User.
func StringMapToPushMessage(data map[string]interface{}) *PushMessage {

	ind := &PushMessage{}
	for k, v := range data {

		switch k {
		case "Hello":
			ind.Hello = v.(string)
		}
	}
	return ind

}
