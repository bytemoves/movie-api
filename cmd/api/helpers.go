package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)


type envelope map[string]any



func (app  *application) writeJSON (w  http.ResponseWriter , status int , data envelope, headers http.Header) error {

	js, err := json.MarshalIndent(data,"","\t")

	if err != nil {
		return err
	}
	js = append(js, '\n')
	//loop thriugh the header map and add each header to thr http.RespWr

	for key,value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type","application/json")

	w.WriteHeader(status)
	w.Write(js)
	return nil

}


func(app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	err := json.NewDecoder(r.Body).Decode(dst)

	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnMarshalError *json.InvalidUnmarshalError

		switch {
			//error.As to check == *json.Syntax err

		case errors.As(err,&syntaxError):
			return fmt.Errorf("bosy contains basly-formed JSON (at character %d)",syntaxError.Offset)

		case errors.Is(err,io.ErrUnexpectedEOF):
			return errors.New("bosy contains badly-formed JSON")

		case errors.As(err,&unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrecr JSON type for fiels %q",unmarshalTypeError)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		case errors.Is(err,io.EOF):
			return errors.New("bosy must be empty")

		case errors.As(err,&invalidUnMarshalError):
			panic(err)

		default:
			return err

		}
	}

	return nil
}


