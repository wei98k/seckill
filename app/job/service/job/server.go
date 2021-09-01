package job

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/go-kratos/kratos/v2/transport"
	event2 "github.com/peter-wow/seckill/app/job/service/event"
	"github.com/peter-wow/seckill/app/job/service/internal/biz"
	"github.com/peter-wow/seckill/app/job/service/internal/conf"
	"github.com/peter-wow/seckill/app/job/service/internal/service"
	"github.com/segmentio/kafka-go"
)

var _ transport.Server = (*Server)(nil)
var _ event2.Message = (*Message)(nil)

type Server struct {
	reader *kafka.Reader
	topic  string
	uo     *service.OrderService
}

type Message struct {
	key    string
	value  []byte
	header map[string]string
}

func (m *Message) Key() string {
	return m.key
}

func (m *Message) Value() []byte {
	return m.value
}

func (m *Message) Header() map[string]string {
	return m.header
}

func NewMessage(key string, value []byte, header map[string]string) event2.Message {
	return &Message{
		key:    key,
		value:  value,
		header: header,
	}
}

func (s Server) Receive(ctx context.Context, handler event2.Handler) error {
	go func() {
		for {
			m, err := s.reader.FetchMessage(context.Background())

			if err != nil {
				break
			}
			h := make(map[string]string)
			if len(m.Headers) > 0 {
				for _, header := range m.Headers {
					h[header.Key] = string(header.Value)
				}
			}
			err = handler(context.Background(), &Message{
				key:    string(m.Key),
				value:  m.Value,
				header: h,
			})
			if err != nil {
				log.Fatal("message handing exception:", err)
			}
			if err := s.reader.CommitMessages(ctx, m); err != nil {
				log.Fatal("failed to commit message:", err)
			}
		}
	}()
	return nil
}

func (s Server) Close() error {
	err := s.reader.Close()
	if err != nil {
		return err
	}
	return nil
}

func NewJOBServer(c *conf.Server, uo *service.OrderService) *Server {

	// []string{"192.168.2.27:9092"}, "order"

	address := []string{"192.168.0.111:9092"}
	topic := "order"

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  address,
		GroupID:  "group-d",
		Topic:    topic,
		MinBytes: 1,    // 10kb
		MaxBytes: 10e6, // 10mb
	})
	return &Server{reader: r, topic: topic, uo: uo}
}

func (s Server) Start(ctx context.Context) error {
	fmt.Printf("job-job start")

	//in := &v1.HelloRequest{
	//	Name: "kratos",
	//}
	//s.t.SayHello(ctx, in)

	s.Receive(ctx, func(ctx context.Context, message event2.Message) error {
		//TODO::路由解析 根据不同的key调用不同的业务逻辑处理

		msg := message.Header()

		//fmt.Println(msg["uid"])

		uid, err := strconv.ParseInt(msg["uid"], 10, 64)
		gid, err := strconv.ParseInt(msg["goods_id"], 10, 64)
		oid, err := strconv.ParseInt(msg["order_id"], 10, 64)

		if err != nil {

		}

		in := &biz.SeckillOrder{
			OrderId: oid,
			UserId:  uid,
			GoodsId: gid,
		}
		s.uo.Create(ctx, in)
		fmt.Printf("key:%s, value:%s, header:%s\n", message.Key(), message.Value(), message.Header())
		return nil
	})

	return nil
}

func (s Server) Stop(ctx context.Context) error {
	err := s.reader.Close()
	if err != nil {
		return err
	}
	return nil
}
