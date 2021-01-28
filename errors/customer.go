package errors

//type HttpError struct {
//	Status int
//	Method string
//}
//func (httpError *HttpError) Error() string {
//	return fmt.Sprintf("Something went wrong with the %v request. Server returned %v status.",
//		httpError.Method, httpError.Status)
//}

type ConstError string

func (err ConstError) Error() string{
	return string(err)
}

const (
	DBError = ConstError("Something went wrong with DataBase")
	HttpError=ConstError("Something went wrong with Request")
	NotFound=ConstError("Id not found")
	OutputFormat=ConstError("unexpected output format")
	BadRequest=ConstError("value of attributes is missing")
)
