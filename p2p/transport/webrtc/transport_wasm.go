//go:build js && wasm

package libp2pwebrtc

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/libp2p/go-libp2p/core/connmgr"
	ic "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/pnet"
	tpt "github.com/libp2p/go-libp2p/core/transport"

	ma "github.com/multiformats/go-multiaddr"
)

var errWasmNotSupported = errors.New("WebRTC transport is not supported in WASM builds")

// ListenUDPFn is a function that creates a UDP listener.
type ListenUDPFn func(network string, laddr *net.UDPAddr) (net.PacketConn, error)

// Option is a function that configures the WebRTC transport.
type Option func(*WebRTCTransport) error

// WebRTCTransport is a stub for WASM builds.
type WebRTCTransport struct{}

var _ tpt.Transport = &WebRTCTransport{}

// New creates a new WebRTCTransport stub that returns an error on WASM.
func New(privKey ic.PrivKey, psk pnet.PSK, gater connmgr.ConnectionGater, rcmgr network.ResourceManager, listenUDP ListenUDPFn, opts ...Option) (*WebRTCTransport, error) {
	return nil, fmt.Errorf("WebRTC transport: %w", errWasmNotSupported)
}

func (t *WebRTCTransport) Dial(ctx context.Context, raddr ma.Multiaddr, p peer.ID) (tpt.CapableConn, error) {
	return nil, errWasmNotSupported
}

func (t *WebRTCTransport) CanDial(addr ma.Multiaddr) bool {
	return false
}

func (t *WebRTCTransport) Listen(laddr ma.Multiaddr) (tpt.Listener, error) {
	return nil, errWasmNotSupported
}

func (t *WebRTCTransport) Protocols() []int {
	return nil
}

func (t *WebRTCTransport) Proxy() bool {
	return false
}
