package mutils

import (
	"time"

	"github.com/jinzhu/copier"
	"github.com/matiniiuu/mcommon/pkg/derrors"
	"github.com/matiniiuu/mcommon/pkg/translator/messages"
	"go.mongodb.org/mongo-driver/v2/bson"
	"google.golang.org/protobuf/types/known/structpb"
)

func GetCopier(toValue interface{}, fromValue interface{}, options ...copier.TypeConverter) (err error) {
	converters := []copier.TypeConverter{ // String to *bson.ObjectID
		//* string to bson.ObjectID to
		{
			SrcType: string(""),
			DstType: bson.ObjectID{},
			Fn: func(src interface{}) (interface{}, error) {
				return bson.ObjectIDFromHex(src.(string))
			},
		},
		//* bson.ObjectID to string
		{
			SrcType: bson.ObjectID{},
			DstType: string(""),
			Fn: func(src interface{}) (interface{}, error) {
				objectID := src.(bson.ObjectID)
				return objectID.Hex(), nil // Convert ObjectID to string
			},
		},
		//* string to *bson.ObjectID to
		{
			SrcType: string(""),
			DstType: &bson.ObjectID{},
			Fn: func(src interface{}) (interface{}, error) {
				objectID, err := bson.ObjectIDFromHex(src.(string))
				if err != nil {
					return nil, derrors.New(derrors.KindInvalid, messages.ParseQueryError)
				}
				return &objectID, nil
			},
		},
		//* *bson.ObjectID to string
		{
			SrcType: &bson.ObjectID{},
			DstType: string(""),
			Fn: func(src interface{}) (interface{}, error) {
				objectIDPtr := src.(*bson.ObjectID)
				if objectIDPtr == nil {
					return "", nil
				}
				return objectIDPtr.Hex(), nil
			},
		},
		//* []string to []*bson.ObjectID
		{
			SrcType: []string{},
			DstType: []*bson.ObjectID{},
			Fn: func(src interface{}) (interface{}, error) {
				strSlice := src.([]string)
				objectIDPtrs := make([]*bson.ObjectID, len(strSlice))
				for i, str := range strSlice {
					objectID, err := bson.ObjectIDFromHex(str)
					if err != nil {
						return nil, derrors.New(derrors.KindInvalid, messages.ParseQueryError)
					}
					objectIDPtrs[i] = &objectID
				}
				return objectIDPtrs, nil
			},
		},
		//* []*bson.ObjectID to []string
		{
			SrcType: []*bson.ObjectID{},
			DstType: []string{},
			Fn: func(src interface{}) (interface{}, error) {
				objectIDPtrs := src.([]*bson.ObjectID)
				strSlice := make([]string, len(objectIDPtrs))
				for i, objectIDPtr := range objectIDPtrs {
					if objectIDPtr == nil {
						strSlice[i] = "" // Handle nil pointers
					} else {
						strSlice[i] = objectIDPtr.Hex()
					}
				}
				return strSlice, nil
			},
		},
		//* map[string]*structpb.Value{} to map[string]interface{}{}
		{
			SrcType: map[string]*structpb.Value{},
			DstType: map[string]interface{}{},
			Fn: func(src interface{}) (interface{}, error) {
				return convertStructpbToInterface(src)
			},
		},
		{
			SrcType: map[string]interface{}{},
			DstType: map[string]*structpb.Value{},
			Fn: func(src interface{}) (interface{}, error) {
				return convertInterfaceToStructpb(src)
			},
		},
		//* interface{} to  *structpb.Value{}
		{
			DstType: structpb.NewNullValue(),
			Fn: func(src interface{}) (interface{}, error) {
				return structpb.NewValue(src)
			},
		},
		//* *time.Time to String
		{
			SrcType: &time.Time{},
			DstType: string(""),
			Fn: func(src interface{}) (interface{}, error) {
				if src == nil {
					return "", nil
				}
				t, ok := src.(*time.Time)
				if !ok {
					return "", derrors.New(derrors.KindInvalid, messages.ParseQueryError)
				}
				return t.Format(time.RFC3339), nil
			},
		},
	}
	converters = append(converters, options...)
	if err := copier.CopyWithOption(toValue, fromValue, copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
		Converters:  converters}); err != nil {
		return derrors.New(derrors.KindInvalid, messages.ParseQueryError)
	}
	return nil
}

// Convertor function to convert from map[string]*structpb.Value to map[string]interface{}
func convertStructpbToInterface(src interface{}) (interface{}, error) {
	srcMap, ok := src.(map[string]*structpb.Value)
	if !ok {
		return nil, derrors.New(derrors.KindInvalid, messages.ParseQueryError)
	}

	dstMap := make(map[string]interface{})

	for key, value := range srcMap {
		switch v := value.Kind.(type) {
		case *structpb.Value_StringValue:
			dstMap[key] = v.StringValue
		case *structpb.Value_NumberValue:
			dstMap[key] = v.NumberValue
		case *structpb.Value_BoolValue:
			dstMap[key] = v.BoolValue
		case *structpb.Value_ListValue:
			dstMap[key] = v.ListValue
		case *structpb.Value_StructValue:
			dstMap[key], _ = convertStructpbToInterface(v.StructValue) // Recursively convert nested map
		default:
			return nil, derrors.New(derrors.KindInvalid, messages.ParseQueryError)
		}
	}

	return dstMap, nil
}

// Convertor function to convert from map[string]interface{} to map[string]*structpb.Value
func convertInterfaceToStructpb(src interface{}) (interface{}, error) {
	srcMap, ok := src.(map[string]interface{})
	if !ok {
		return nil, derrors.New(derrors.KindInvalid, messages.ParseQueryError)
	}

	dstMap := make(map[string]*structpb.Value)

	for key, value := range srcMap {
		switch v := value.(type) {
		case string:
			dstMap[key] = structpb.NewStringValue(v)
		case float64:
			dstMap[key] = structpb.NewNumberValue(v)
		case bool:
			dstMap[key] = structpb.NewBoolValue(v)
		case []interface{}:
			// Convert list of interface{} to ListValue
			listValues := make([]*structpb.Value, len(v))
			for i, item := range v {
				convertedItem, err := convertInterfaceToStructpb(item)
				if err != nil {
					return nil, derrors.New(derrors.KindInvalid, messages.ParseQueryError)
				}
				listValues[i] = convertedItem.(*structpb.Value)
			}
			dstMap[key] = structpb.NewListValue(&structpb.ListValue{Values: listValues})
		case map[string]interface{}:
			// Recursively convert nested map to structpb.Struct
			convertedMap, err := convertInterfaceToStructpb(v)
			if err != nil {
				return nil, derrors.New(derrors.KindInvalid, messages.ParseQueryError)
			}
			// Convert the inner map into a structpb.Struct
			innerStruct := convertedMap.(*structpb.Struct)
			dstMap[key] = structpb.NewStructValue(innerStruct)
		default:
			return nil, derrors.New(derrors.KindInvalid, messages.ParseQueryError)
		}
	}

	// Return the final converted map
	return dstMap, nil
}

func EnumCopier[T ~string, PB ~int32](val T, pbMap map[string]int32, dft PB) *PB {
	v, ok := pbMap[string(val)]
	if !ok {
		return &dft
	}
	result := PB(v)
	return &result
}
