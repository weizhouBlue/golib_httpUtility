package golib_httpUtility
import (
    "net/http"
    "context"
    "strings"
    "fmt"
    "io"
    "time"
    "io/ioutil"
)

type HttpMethod string
const (
	MethodPost HttpMethod = "POST"
	MethodGet HttpMethod = "GET"
	MethodDel HttpMethod = "DELETE"
	MethodHead HttpMethod = "HEAD"
	MethodPatch HttpMethod = "PATCH"
	MethodPut HttpMethod = "PUT"

)

func HttpClient( severUrl string  , method HttpMethod , requestHeader map[string][]string, requestBody string , timeout int ) (  returnCode int , reponseBody string , reponseHeader map[string][]string,  err error ){

	var ctx context.Context
	//var cancel context.CancelFunc

	var request *http.Request
	var er error
	var msg io.Reader
	var response *http.Response
	var httpClient = &http.Client{}

	err=nil
	returnCode=0
	reponseBody=""
	reponseHeader=nil


	if timeout > 0 {
		ctx, _ = context.WithTimeout(context.Background(), time.Duration(timeout) * time.Second )
	}else{
		ctx, _ = context.WithCancel(context.Background() )
	}

	if len(requestBody)>0{
		msg=strings.NewReader(requestBody)
	}else{
		msg=nil
	}

	request, er=http.NewRequestWithContext( ctx , string(method) , severUrl , msg )
	if er!=nil {
		err=fmt.Errorf( "%v", er )
		return
	}

	if len(requestHeader)>0{
		request.Header=requestHeader
	}

	response , er=httpClient.Do(request)
	if er!=nil{
		err=fmt.Errorf( "%v", er )
		return
	}

	returnCode=response.StatusCode
	reponseHeader=map[string][]string (response.Header)


	defer response.Body.Close()
    if body, er := ioutil.ReadAll(response.Body) ; er!=nil {
    	if er!=io.EOF{
			err=fmt.Errorf( "%v", er )
			return    		
    	}
    }else{
    	if len(body)>0 {
    		reponseBody=string(body)
    	}
    }


	return
}












