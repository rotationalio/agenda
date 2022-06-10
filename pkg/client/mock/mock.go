package mock

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/rotationalio/agenda/pkg/api/v1"
	"github.com/rotationalio/agenda/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	DailyRPC    = "agenda.v1.Agenda/Daily"
	ScheduleRPC = "agenda.v1.Agenda/Schedule"
)

// New creates a new mock RemotePeer. If bufnet is nil, one is created for the user.
func New(bufnet *utils.Listener) *AgendaServer {
	if bufnet == nil {
		bufnet = utils.New()
	}

	remote := &AgendaServer{
		bufnet: bufnet,
		srv:    grpc.NewServer(),
		Calls:  make(map[string]int),
	}

	api.RegisterAgendaServer(remote.srv, remote)
	go remote.srv.Serve(remote.bufnet.Sock())
	return remote
}

type AgendaServer struct {
	api.UnimplementedAgendaServer
	bufnet     *utils.Listener
	srv        *grpc.Server
	Calls      map[string]int
	OnSchedule func(context.Context, *api.Item) (*api.Item, error)
	OnDaily    func(context.Context, *api.Day) (*api.Docket, error)
}

func (s *AgendaServer) Channel() *utils.Listener {
	return s.bufnet
}

func (s *AgendaServer) Shutdown() {
	s.srv.GracefulStop()
	s.bufnet.Close()
}

func (s *AgendaServer) Reset() {
	for key := range s.Calls {
		s.Calls[key] = 0
	}

	s.OnDaily = nil
	s.OnSchedule = nil
}

// UseFixture loadsa a JSON fixture from disk (usually in a testdata folder) to use as
// the protocol buffer response to the specified RPC, simplifying handler mocking.
func (s *AgendaServer) UseFixture(rpc, path string) (err error) {
	var data []byte
	if data, err = ioutil.ReadFile(path); err != nil {
		return fmt.Errorf("could not read fixture: %v", err)
	}

	jsonpb := &protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: true,
	}

	switch rpc {
	case ScheduleRPC:
		out := &api.Item{}
		if err = jsonpb.Unmarshal(data, out); err != nil {
			return fmt.Errorf("could not unmarshal json into %T: %v", out, err)
		}
		s.OnSchedule = func(context.Context, *api.Item) (*api.Item, error) {
			return out, nil
		}
	case DailyRPC:
		out := &api.Docket{}
		if err = jsonpb.Unmarshal(data, out); err != nil {
			return fmt.Errorf("could not unmarshal json into %T: %v", out, err)
		}
		s.OnDaily = func(context.Context, *api.Day) (*api.Docket, error) {
			return out, nil
		}
	default:
		return fmt.Errorf("unknown RPC %q", rpc)
	}

	return nil
}

// UseError allows you to specify a gRPC status error to return from the specified RPC.
func (s *AgendaServer) UseError(rpc string, code codes.Code, msg string) error {
	switch rpc {
	case ScheduleRPC:
		s.OnSchedule = func(context.Context, *api.Item) (*api.Item, error) {
			return nil, status.Error(code, msg)
		}
	case DailyRPC:
		s.OnDaily = func(context.Context, *api.Day) (*api.Docket, error) {
			return nil, status.Error(code, msg)
		}
	default:
		return fmt.Errorf("unknown RPC %q", rpc)
	}
	return nil
}

func (s *AgendaServer) Daily(ctx context.Context, in *api.Day) (out *api.Docket, err error) {
	s.Calls[DailyRPC]++
	return s.OnDaily(ctx, in)
}

func (s *AgendaServer) Schedule(ctx context.Context, in *api.Item) (out *api.Item, err error) {
	s.Calls[ScheduleRPC]++
	return s.OnSchedule(ctx, in)
}
