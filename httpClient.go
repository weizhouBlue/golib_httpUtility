package golib_httpUtility
import (
	"net"
    "net/http"
    "context"
    "strings"
    "fmt"
    "io"
    "time"
    "io/ioutil"

    "crypto/x509"
    "crypto/tls"
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

type TlsConf struct {
	IgnoreServerCa bool
	CaPath string
	CertPath string
	KeyPath string
}

func HttpClient( severUrl string  , method HttpMethod , requestHeader map[string][]string, requestBody string , timeout int , unixSkPath string , tlsConf *TlsConf ) (  returnCode int , reponseBody string , reponseHeader map[string][]string,  err error ){

	var ctx context.Context
	//var cancel context.CancelFunc

	var request *http.Request
	var er error
	var msg io.Reader
	var response *http.Response

	err=nil
	returnCode=0
	reponseBody=""
	reponseHeader=nil

	if timeout==0 {
		timeout=60
	}


	// https://godoc.org/net/http#Transport
	tr:=&http.Transport{}

	if tlsConf!=nil {
		// https://godoc.org/crypto/tls#Config
	    tr.TLSClientConfig=&tls.Config{ InsecureSkipVerify: tlsConf.IgnoreServerCa }
		if len(tlsConf.CaPath)!=0 {
		    pool := x509.NewCertPool()
		    caCrt, e1 := ioutil.ReadFile(tlsConf.CaPath)
		    if e1 != nil {
				err=fmt.Errorf(  "ReadFile err: %v", e1  )
		        return
		    }
		    pool.AppendCertsFromPEM(caCrt)
		    tr.TLSClientConfig.RootCAs=pool
		}

		if len(tlsConf.CertPath)!=0 && len(tlsConf.KeyPath)!=0  {
		    cliCrt, e2 := tls.LoadX509KeyPair( tlsConf.CertPath , tlsConf.KeyPath )
		    if e2 != nil {
				err=fmt.Errorf(  "Loadx509keypair err: %v ", e2  )
		        return
		    }
		    tr.TLSClientConfig.Certificates=[]tls.Certificate{cliCrt}
		}
	}

	if len(unixSkPath)!=0 {
		tr.DialContext = func( _ context.Context, _, _ string) (net.Conn, error) {
			return net.DialTimeout("unix", unixSkPath , time.Duration(timeout)*time.Second )
		}
	}

	// 必须设置如下 其一，否则，会发现client端 一直keepalive connection，使得client和server端一直 携程泄漏，connection泄漏
	tr.DisableKeepAlives=true
	defer tr.CloseIdleConnections()

	//tr.ForceAttemptHTTP2=true
	// maximum amount of time an idle (keep-alive) connection will remain idle before closing itself
	//tr.IdleConnTimeout=3 0 * time.Second

	// https://godoc.org/net/http#Client
	httpClient := &http.Client{Transport: tr}

	if len(severUrl) == 0 {
		err=fmt.Errorf( "empty url " )
		return
	}


	ctx, _ = context.WithTimeout(context.Background(), time.Duration(timeout) * time.Second )

	if len(requestBody)>0{
		msg=strings.NewReader(requestBody)
	}else{
		msg=nil
	}

	// generate a request : https://godoc.org/net/http#Request
	request, er=http.NewRequestWithContext( ctx , string(method) , severUrl , msg )
	if er!=nil {
		err=fmt.Errorf( "%v", er )
		return
	}

	if len(requestHeader)>0{
		request.Header=requestHeader
	}

	// send request 
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




