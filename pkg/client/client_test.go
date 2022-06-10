package client_test

import (
	"testing"
	"time"

	"github.com/rotationalio/agenda/pkg/client"
	"github.com/rotationalio/agenda/pkg/client/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

func TestSchedule(t *testing.T) {
	// Setup RemoteAgenda mock server
	srv := mock.New(nil)
	defer srv.Shutdown()

	// Create a client to test that is connected to the mock server.
	dialer := grpc.WithContextDialer(srv.Channel().Dialer)
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())

	agenda, err := client.New("bufconn", dialer, creds)
	require.NoError(t, err, "could not connect to remote agenda via bufconn")

	// Test the case where the server returns an error
	srv.UseError(mock.ScheduleRPC, codes.DataLoss, "something bad happened")

	err = agenda.Schedule("hello", "world", time.Now(), time.Now().Add(5*time.Minute))
	require.Error(t, err, "expected an error returned from the server")
	require.Equal(t, 1, srv.Calls[mock.ScheduleRPC])

	// Test the case where server returns a response
	srv.UseFixture(mock.ScheduleRPC, "testdata/item.json")
	err = agenda.Schedule("hello", "world", time.Now(), time.Now().Add(5*time.Minute))
	require.NoError(t, err, "expected no error in happy path")
	require.Equal(t, 2, srv.Calls[mock.ScheduleRPC])
}
