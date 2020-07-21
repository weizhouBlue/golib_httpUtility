package golib_httpUtility_test
import (
    "testing"
    http "github.com/weizhouBlue/golib_httpUtility"
    "fmt"
)

//====================================

func Test_http(t *testing.T){


    severUrl:="http://10.6.185.10:11160/parcel/v1/manager/ping"
    method:=http.MethodGet
    requestHeader:= map[string][]string {
        "testHeader":  {"12345"} ,
    }
    requestBody:=""
    timeout:=0
    unixSkPath:=""

    returnCode , reponseBody , reponseHeader ,  err := http.HttpClient( severUrl   , method ,requestHeader , requestBody  , timeout ,unixSkPath , nil ) 
    if err!=nil {
        fmt.Println("err: " , err )
        t.FailNow()
    }
    fmt.Println("return code: " , returnCode)
    fmt.Println("response header: ", reponseHeader)
    fmt.Println("reponse body: " , reponseBody)


}


func Test_https(t *testing.T){

    severUrl:="https://10.6.185.10:6443/"
    method:=http.MethodGet
    requestHeader:= map[string][]string {}
    requestBody:=""
    timeout:=0
    unixSkPath:=""
    tlsConf:=  &http.TlsConf{
        IgnoreServerCa : false ,
        CaPath : "/etc/kubernetes/pki/ca.crt" ,
        CertPath : "/etc/kubernetes/pki/apiserver-kubelet-client.crt" ,
        KeyPath : "/etc/kubernetes/pki/apiserver-kubelet-client.key",
    }

    returnCode , reponseBody , reponseHeader ,  err := http.HttpClient( severUrl   , method ,requestHeader , requestBody  , timeout ,unixSkPath , tlsConf ) 
    if err!=nil {
        fmt.Println("err: " , err )
        t.FailNow()
    }
    fmt.Println("return code: " , returnCode)
    fmt.Println("response header: ", reponseHeader)
    fmt.Println("reponse body: " , reponseBody)

}


// can use : socat -d -d UNIX-LISTEN:/tmp/testunix.socket  tcp:10.6.185.10:6443
func Test_unix(t *testing.T){

    severUrl:="https://10.6.185.10:6443/"
    method:=http.MethodGet
    requestHeader:= map[string][]string {}
    requestBody:=""
    timeout:=0
    unixSkPath:="/tmp/test.socket"
    tlsConf:=  &http.TlsConf{
        IgnoreServerCa : true ,
        CaPath : "" ,
        CertPath : "" ,
        KeyPath : "",
    }

    returnCode , reponseBody , reponseHeader ,  err := http.HttpClient( severUrl   , method ,requestHeader , requestBody  , timeout ,unixSkPath , tlsConf ) 
    if err!=nil {
        fmt.Println("err: " , err )
        t.FailNow()
    }
    fmt.Println("return code: " , returnCode)
    fmt.Println("response header: ", reponseHeader)
    fmt.Println("reponse body: " , reponseBody)


}





