package golib_httpUtility_test
import (
    "testing"
    http "github.com/weizhouBlue/golib_httpUtility"
    "fmt"
)

//====================================

func Test_http(t *testing.T){


    severUrl:="http://10.6.185.160:6081/parcel/v1/manager/ping"
    method:=http.MethodGet
    requestHeader:= map[string][]string {
        "testHeader":  {"12345"} ,
    }
    requestBody:=""
    timeout:=0

    returnCode , reponseBody , reponseHeader ,  err := http.HttpClient( severUrl   , method ,requestHeader , requestBody  , timeout  ) 
    if err!=nil {
        fmt.Println("err: " , err )
        t.FailNow()
    }
    fmt.Println("return code: " , returnCode)
    fmt.Println("response header: ", reponseHeader)
    fmt.Println("reponse body: " , reponseBody)


}

