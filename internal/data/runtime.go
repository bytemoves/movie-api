package data

import (
	"fmt"
	"strconv"
)


type Runtime int32 


//used to return the json encoded value for the movie

func ( r Runtime) MarshalJSON ([]byte,error) {
	//gen a string containing the movie runtime 
	jsonValue := fmt.Sprintf("%d mins",r)

	quotedJSONValue := strconv.Quote(jsonValue)

	return  []byte(quotedJSONValue),nil

}