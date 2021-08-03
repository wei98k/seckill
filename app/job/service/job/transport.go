package job

import "github.com/go-kratos/kratos/v2/transport"

var (
	_ transport.Transporter = &Transport{}
)


type Transport struct {

}

func (t Transport) Kind() transport.Kind {
	panic("implement me")
}

func (t Transport) Endpoint() string {
	panic("implement me")
}

func (t Transport) Operation() string {
	panic("implement me")
}

func (t Transport) RequestHeader() transport.Header {
	panic("implement me")
}

func (t Transport) ReplyHeader() transport.Header {
	panic("implement me")
}

