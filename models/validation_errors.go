package models

// import (
// 	"github.com/go-playground/validator/v10"
// )

// type ValidationError struct {
// 	Message string              `json:"message"`
// 	Details []ValidationDetails `json:"details"`
// }

// type ValidationDetails struct {
// 	Field      string `json:"field"`
// 	Value      string `json:"value"`
// 	Constraint string `json:"constraint"`
// }

// type GinErrors interface {
// 	ListAllErrors(model interface{}, err error) map[string]string
// }

// type ginErrors struct {
// 	errorMaps map[string]string
// }

// type ErrorResult struct {
// 	Field   string
// 	JsonTag string
// 	Message string
// }

// func NewShyGinErrors(errors map[string]string) GinErrors {
// 	return &ginErrors{
// 		errorMaps: errors,
// 	}
// }

// func (ge ginErrors) ListAllErrors(model interface{}, err error) map[string]string {
// 	errors := map[string]string{}
// 	fields := map[string]ErrorResult{}

// 	if _, ok := err.(validator.ValidationErrors); ok {
// 		// resolve all json tags for the struct
// 		types := reflect.TypeOf(model)
// 		values := reflect.ValueOf(model)

// 		for i := 0; i < types.NumField(); i++ {
// 			field := types.Field(i)
// 			value := values.Field(i).Interface()
// 			jsonTag := field.Tag.Get("json")
// 			if jsonTag == "" {
// 				jsonTag = field.Name
// 			}
// 			messageTag := field.Tag.Get("msg")
// 			msg := ge.getErrorMessage(messageTag)

// 			fmt.Printf("%s: %v = %v, tag= %v\n", field.Name, field.Type, value, jsonTag)
// 			fields[field.Name] = ErrorResult{
// 				Field:   field.Name,
// 				JsonTag: jsonTag,
// 				Message: msg,
// 			}
// 		}

// 		for _, e := range err.(validator.ValidationErrors) {
// 			if field, ok := fields[e.Field()]; ok {
// 				if field.Message != "" {
// 					errors[field.JsonTag] = field.Message
// 				} else {
// 					errors[field.JsonTag] = e.Error()
// 				}
// 			}
// 		}
// 	} else {
// 		errors["0"] = err.Error()
// 	}

// 	return errors
// }

// func (ge ginErrors) getErrorMessage(key string) string {
// 	if value, ok := ge.errorMaps[key]; ok {
// 		return value
// 	}
// 	return key
// }

// func FromBindingError(data interface{}, e error) ValidationError {
// 	errors := ValidationError{
// 		Message: "validation error",
// 	}
// 	ve := e.(validator.ValidationErrors)
// 	// InvalidFields := make([]map[string]string, 0)

// 	for _, e := range ve {
// 		fmt.Println(e)
// 		vr := ValidationDetails{}
// 		// errors := map[string]string{}
// 		// field, _ := reflect.TypeOf(data).Elem().FieldByName(e.Type().Name())
// 		// jsonTag := string(field.Tag.Get("json"))
// 		vr.Field = e.Field()
// 		vr.Constraint = e.Tag()
// 		// vr.Constraint =
// 		// errors[jsonTag] = e.Tag
// 		errors.Details = append(errors.Details, vr)
// 	}

// 	return errors
// }
