package dto

import (
	"encoding/json"

	"github.com/speps/go-hashids/v2"
)

type IdObject struct {
	Value    string
	ValueInt int64
}

func InitialIdObject(hashIdLocal *hashids.HashID) {
	hashId = hashIdLocal
}

var hashId *hashids.HashID

func NewIdObject(id string) (*IdObject, error) {
	values, err := hashId.DecodeInt64WithError(id)
	if err != nil {
		return nil, err
	}

	return &IdObject{
		Value:    id,
		ValueInt: values[0],
	}, nil
}

func (id *IdObject) MarshalJSON() ([]byte, error) {
	var err error
	id.Value, err = hashId.EncodeInt64([]int64{id.ValueInt})
	if err != nil {
		return nil, err
	}
	return json.Marshal(id.Value)
}

func (id *IdObject) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		// if empty value then check
		var mapId map[string]interface{}
		errConv := json.Unmarshal(data, &mapId)
		if errConv != nil {
			return err
		}

		// expect only 2 fields Value and ValueInt
		if len(mapId) != 2 {
			return err
		}

		stringValue, valueExists := mapId["Value"]
		intValue, ValueIntExists := mapId["ValueInt"]

		if !valueExists || !ValueIntExists {
			return err
		}

		id.Value = stringValue.(string)
		id.ValueInt = int64(intValue.(float64))

		return nil
	}

	id.Value = value
	values, err := hashId.DecodeInt64WithError(id.Value)
	if err != nil {
		return err
	}
	id.ValueInt = values[0]
	return nil
}

func NewIntIdObject(id int64) (*IdObject, error) {
	idObj := IdObject{ValueInt: id}
	if hashId != nil {
		var err error
		idObj.Value, err = hashId.EncodeInt64([]int64{id})
		if err != nil {
			return nil, err
		}
	}

	return &idObj, nil
}
