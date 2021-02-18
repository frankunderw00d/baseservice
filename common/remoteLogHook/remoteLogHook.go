package remoteLogHook

import (
	"io"
	"jarvis/base/network"
)

type (
	remoteHook struct {
		client network.Client
	}
)

const ()

var ()

func NewSocketRemoteHook(address string) (io.WriteCloser, error) {
	client := network.NewSocketClient(address, network.DefaultPackager(), network.DefaultEncrypter())

	if err := client.Initialize(); err != nil {
		return nil, err
	}

	return newRemoteHook(client), nil
}

func NewWebSocketRemoteHook(address string) (io.WriteCloser, error) {
	client := network.NewWebSocketClient(address, network.DefaultPackager(), network.DefaultEncrypter())

	if err := client.Initialize(); err != nil {
		return nil, err
	}

	return newRemoteHook(client), nil
}

func NewGRPCRemoteHook(address string) (io.WriteCloser, error) {
	client := network.NewGRPCClient(address, network.DefaultPackager(), network.DefaultEncrypter())

	if err := client.Initialize(); err != nil {
		return nil, err
	}

	return newRemoteHook(client), nil
}

func newRemoteHook(client network.Client) io.WriteCloser {
	return &remoteHook{
		client: client,
	}
}

func (rh *remoteHook) Write(b []byte) (int, error) {
	if err := rh.client.Send(network.Message{
		Module: "Log",
		Route:  "print",
		Data:   b,
	}); err != nil {
		return 0, err
	}

	return len(b), nil
}

func (rh *remoteHook) Close() error {
	return rh.client.Close()
}
