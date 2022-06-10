package client

import (
	"context"
	"time"

	"github.com/rotationalio/agenda/pkg/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	cc  *grpc.ClientConn
	rpc api.AgendaClient
}

func New(endpoint string, opts ...grpc.DialOption) (c *Client, err error) {
	c = &Client{}

	if len(opts) == 0 {
		// Production connection opts (TLS, etc.)
		opts = make([]grpc.DialOption, 0)
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	if c.cc, err = grpc.Dial(endpoint, opts...); err != nil {
		return nil, err
	}
	c.rpc = api.NewAgendaClient(c.cc)

	return c, nil
}

func (c *Client) Close() error {
	return c.cc.Close()
}

func (c *Client) Schedule(title, description string, start, end time.Time) error {
	item := &api.Item{
		Title:       title,
		Description: description,
		Date:        start.Format("2006-01-02"),
		Start:       start.Format("15:04"),
		End:         end.Format("15:04"),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := c.rpc.Schedule(ctx, item); err != nil {
		return err
	}
	return nil
}
