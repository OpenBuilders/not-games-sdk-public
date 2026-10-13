package contracts_test

import (
	"bytes"
	"testing"

	squadspb "github.com/OpenBuilders/not-games-sdk-public/gen/grpc/go/squads"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestGetAccountSquadContract(t *testing.T) {
	service := squadspb.File_squads_squads_proto.Services().ByName("SquadsService")
	method := service.Methods().ByName("GetAccountSquad")
	if method == nil || method.IsStreamingClient() || method.IsStreamingServer() {
		t.Fatal("GetAccountSquad must be a unary RPC")
	}
	if squadspb.SquadsService_GetAccountSquad_FullMethodName != "/squad.v1.SquadsService/GetAccountSquad" {
		t.Fatal("unexpected GetAccountSquad method path")
	}

	accountID := method.Input().Fields().ByName("account_id")
	if accountID == nil || accountID.Number() != 1 || accountID.Kind() != protoreflect.Int64Kind {
		t.Fatal("account_id must remain int64 field 1")
	}
	squad := method.Output().Fields().ByName("squad")
	if squad == nil || squad.Number() != 1 || squad.Message().FullName() != "squad.v1.Squad" || !squad.HasPresence() {
		t.Fatal("squad must be a presence-aware Squad message at field 1")
	}
}

func TestGetAccountSquadRequestWireRoundTrip(t *testing.T) {
	want := &squadspb.GetAccountSquadRequest{AccountId: 123456789}
	wire, err := proto.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(wire, []byte{0x08, 0x95, 0x9a, 0xef, 0x3a}) {
		t.Fatalf("unexpected account_id encoding: %x", wire)
	}
	got := new(squadspb.GetAccountSquadRequest)
	if err := proto.Unmarshal(wire, got); err != nil {
		t.Fatal(err)
	}
	if !proto.Equal(got, want) {
		t.Fatalf("account ID did not survive round trip: got %v", got)
	}
}

func TestGetAccountSquadResponsePresence(t *testing.T) {
	for _, squad := range []*squadspb.Squad{nil, {Id: 42, Name: "Squad"}} {
		want := &squadspb.GetAccountSquadResponse{Squad: squad}
		wire, err := proto.Marshal(want)
		if err != nil {
			t.Fatal(err)
		}
		got := new(squadspb.GetAccountSquadResponse)
		if err := proto.Unmarshal(wire, got); err != nil {
			t.Fatal(err)
		}
		if !proto.Equal(got, want) {
			t.Fatalf("squad presence did not survive round trip: got %v, want %v", got, want)
		}
	}
	json, err := protojson.Marshal(&squadspb.GetAccountSquadResponse{})
	if err != nil {
		t.Fatal(err)
	}
	if string(json) != "{}" {
		t.Fatalf("no squad must encode as an empty response, got %s", json)
	}
}
