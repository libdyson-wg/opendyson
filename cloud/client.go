package cloud

import (
	"context"
	"fmt"
	"github.com/libdyson-wg/libdyson-go/internal/generated/oapi"
	"net/http"
)

type ServerRegion int

const (
	RegionGlobal ServerRegion = iota
	RegionChina
)

var (
	region ServerRegion = RegionGlobal
	client oapi.ClientWithResponsesInterface

	provisioned bool

	token string

	servers = []string{
		"https://appapi.cp.dyson.com",
		"https://appapi.cp.dyson.cn",
	}
)

func init() {
	setServer()
}

func addAuthToken(_ context.Context, r *http.Request) error {
	if token != "" {
		r.Header.Set("Authorization", "Bearer "+token)
	}
	return nil
}

func addUserAgent(_ context.Context, r *http.Request) error {
	r.Header.Set("User-Agent", "android client")
	return nil
}

func setServer() {
	var err error
	client, err = oapi.NewClientWithResponses(
		servers[region],
		oapi.WithRequestEditorFn(addAuthToken),
		oapi.WithRequestEditorFn(addUserAgent),
	)
	if err != nil {
		panic(fmt.Errorf("unable to initialize client: %w", err))
	}

	_, err = client.ProvisionWithResponse(context.Background())
	if err != nil {
		panic(fmt.Errorf("unable to provision api client: %w", err))
	}
}

func SetToken(t string) {
	token = t
}

func SetServerRegion(r ServerRegion) {
	region = r
	setServer()
}
