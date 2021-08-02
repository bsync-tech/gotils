package upstream

import (
	"container/list"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"strings"

	//	"context"
	"io"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"

	"gitlab.intsig.net/qxb_cus/common/proxy"

	log "github.com/cihub/seelog"
)

type Peer struct {
	host   string
	weight uint
}

type Upstream struct {
	mutex       sync.Mutex
	totalweight uint
	index       uint
	peers       *list.List
	//	Client      *http.Client
	Client *HttpClient
}

func NewUpstream() *Upstream {
	client := &HttpClient{
		Client: http.Client{
			Transport: &http.Transport{
				Dial: func(network, addr string) (net.Conn, error) {
					dial := net.Dialer{
						Timeout:   30 * time.Second,
						KeepAlive: 30 * time.Second,
					}

					conn, err := dial.Dial(network, addr)
					if err != nil {
						return nil, err
					}

					log.Debugf("Connect done, [local] %s [remote] %s", conn.LocalAddr().String(), conn.RemoteAddr().String())
					return conn, err
				},
				MaxIdleConnsPerHost:   500,
				ResponseHeaderTimeout: time.Second * 60,
				Proxy:                 proxy.GetInstance().GetProxy(),
			},
			Timeout: time.Duration(300 * time.Second),
		},
	}

	return &Upstream{
		peers:       list.New(),
		index:       0,
		totalweight: 0,
		Client:      client,
	}
}

// 不需要走代理的UpStream
func NewUpstreamWithoutProxy() *Upstream {
	client := &HttpClient{
		Client: http.Client{
			Transport: &http.Transport{
				Dial: func(network, addr string) (net.Conn, error) {
					dial := net.Dialer{
						Timeout:   30 * time.Second,
						KeepAlive: 30 * time.Second,
					}

					conn, err := dial.Dial(network, addr)
					if err != nil {
						return nil, err
					}

					log.Debugf("Connect done, [local] %s [remote] %s", conn.LocalAddr().String(), conn.RemoteAddr().String())
					return conn, err
				},
				MaxIdleConnsPerHost:   500,
				ResponseHeaderTimeout: time.Second * 60,
			},
			Timeout: time.Duration(60 * time.Second),
		},
	}

	return &Upstream{
		peers:       list.New(),
		index:       0,
		totalweight: 0,
		Client:      client,
	}
}

// 测试时，caCrtPath可以设置为""，客户端不会检验服务端发送过来的任何ca证书
// 线上时，建议caCrtPath设置为服务端下发的证书的本地路径，否则，可能遭受“中间人攻击”
func NewHTTPSUpstream(caCrtPath string) *Upstream {
	var (
		pool               *x509.CertPool
		InsecureSkipVerify bool = true
	)

	if "" != caCrtPath {
		//管理数字证书
		pool = x509.NewCertPool()
		caCrt, err := ioutil.ReadFile(caCrtPath)
		if err != nil {
			fmt.Printf("ReadFile: %s err: %v\n", caCrtPath, err)
			return nil
		}
		//将生成的数字证书添加到数字证书集合中
		if ok := pool.AppendCertsFromPEM(caCrt); !ok {
			fmt.Printf("append cert from file: %s fail\n", caCrtPath)
			return nil
		}

		InsecureSkipVerify = false
	}

	client := &HttpClient{
		Client: http.Client{
			Transport: &http.Transport{
				TLSHandshakeTimeout: time.Second * 30,
				TLSClientConfig: &tls.Config{
					RootCAs:            pool,
					InsecureSkipVerify: InsecureSkipVerify,
					Certificates:       []tls.Certificate{tls.Certificate{}},
				},
				Dial: func(network, addr string) (net.Conn, error) {
					dial := net.Dialer{
						Timeout:   30 * time.Second,
						KeepAlive: 30 * time.Second,
					}

					conn, err := dial.Dial(network, addr)
					if err != nil {
						return nil, err
					}

					//					log.Debugf("Connect done, [local] %s [remote] %s", conn.LocalAddr().String(), conn.RemoteAddr().String())
					return conn, err
				},
				MaxIdleConnsPerHost:   500,
				ResponseHeaderTimeout: time.Second * 60,
				Proxy:                 proxy.GetInstance().GetProxy(),
			},
			Timeout: time.Duration(60 * time.Second),
		},
	}

	return &Upstream{
		peers:       list.New(),
		index:       0,
		totalweight: 0,
		Client:      client,
	}
}

func NewUnixUpstream(sock string) *Upstream {
	client := &HttpClient{
		Client: http.Client{
			Transport: &http.Transport{
				Dial: func(network, addr string) (net.Conn, error) {
					dial := net.Dialer{
						Timeout:   30 * time.Second,
						KeepAlive: 30 * time.Second,
					}

					conn, err := dial.Dial("unix", sock)
					if err != nil {
						return nil, err
					}

					//					log.Debugf("Connect done, [local] %s [remote] %s", conn.LocalAddr().String(), conn.RemoteAddr().String())
					return conn, err
				},
				MaxIdleConnsPerHost:   500,
				ResponseHeaderTimeout: time.Second * 30,
			},
			Timeout: time.Duration(60 * time.Second),
		},
	}

	return &Upstream{
		peers:       list.New(),
		index:       0,
		totalweight: 0,
		Client:      client,
	}
}

func NewUpstreamWithTimeout(dial, read, write, keepalive time.Duration) *Upstream {
	if dial < 0 {
		dial = 30
	}
	if read < 0 {
		read = 30
	}

	if write < 0 {
		write = 30
	}

	if keepalive < 0 {
		keepalive = 30
	}

	totalTimeout := dial + read + write
	client := &HttpClient{
		Client: http.Client{
			Transport: &http.Transport{
				Dial: func(network, addr string) (net.Conn, error) {
					dial := net.Dialer{
						Timeout:   dial * time.Second,
						KeepAlive: keepalive * time.Second,
					}

					conn, err := dial.Dial(network, addr)
					if err != nil {
						return nil, err
					}

					//					log.Debugf("Connect done, [local] %s [remote] %s", conn.LocalAddr().String(), conn.RemoteAddr().String())
					return conn, err
				},
				MaxIdleConnsPerHost:   500,
				ResponseHeaderTimeout: time.Second * read,
			},
			Timeout: time.Duration(totalTimeout * time.Second),
		},
	}

	return &Upstream{
		peers:       list.New(),
		index:       0,
		totalweight: 0,
		Client:      client,
	}
}

//func NewUpstreamWithProxy() *Upstream {
//	client := &HttpClient{
//		Client: http.Client{
//			Transport: &http.Transport{
//				Dial: func(network, addr string) (net.Conn, error) {
//					dial := net.Dialer{
//						Timeout:   30 * time.Second,
//						KeepAlive: 30 * time.Second,
//					}

//					conn, err := dial.Dial(network, addr)
//					if err != nil {
//						return nil, err
//					}

//					//					log.Debugf("Connect done, [local] %s [remote] %s", conn.LocalAddr().String(), conn.RemoteAddr().String())
//					return conn, err
//				},
//				MaxIdleConnsPerHost:   500,
//				ResponseHeaderTimeout: time.Second * 30,
//				Proxy:                 http.ProxyFromEnvironment,
//			},
//			Timeout: time.Duration(60 * time.Second),
//		},
//	}

//	return &Upstream{
//		peers:       list.New(),
//		index:       0,
//		totalweight: 0,
//		Client:      client,
//	}
//}

func (us *Upstream) AddPeer(host string, weight uint) {
	us.mutex.Lock()
	defer us.mutex.Unlock()
	if !strings.HasPrefix(host, "http") {
		host = "http://" + host
	}

	us.peers.PushBack(&Peer{
		host:   host,
		weight: weight,
	})
	us.totalweight = us.totalweight + weight
}

func (us *Upstream) RemovePeer(host string) {
	us.mutex.Lock()
	defer us.mutex.Unlock()
	for e := us.peers.Front(); e != nil; e = e.Next() {
		if e.Value.(*Peer).host == host {
			us.totalweight = us.totalweight - e.Value.(*Peer).weight
			us.peers.Remove(e)
			break
		}
	}
}

func (us *Upstream) Gethost() string {
	us.mutex.Lock()
	defer us.mutex.Unlock()
	if 0 == us.totalweight {
		return ""
	}
	us.index = us.index + 1
	if us.index > us.totalweight {
		us.index = 1
	}
	tmp_weight := uint(0)
	for e := us.peers.Front(); e != nil; e = e.Next() {
		tmp_weight = tmp_weight + e.Value.(*Peer).weight
		if us.index <= tmp_weight {
			return e.Value.(*Peer).host
		}
	}
	return ""
}

func (us *Upstream) Close() {
	us.Client.Transport.(*http.Transport).CloseIdleConnections()
}

func (us *Upstream) Getclient() *HttpClient {
	return us.Client
}

type HttpClient struct {
	http.Client
}

func (cli *HttpClient) Do(req *http.Request) (resp *http.Response, err error) {
	resp, err = cli.Client.Do(req)
	if urlErr, ok := err.(*url.Error); ok && (urlErr.Temporary() || urlErr.Err == io.EOF) {
		resp, err = cli.Client.Do(req)
	}
	return
}
