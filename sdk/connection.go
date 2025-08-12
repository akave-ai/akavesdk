// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package sdk

import (
	"fmt"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/akave-ai/akavesdk/private/pb"
)

type connectionPool struct {
	mu          sync.RWMutex
	connections map[string]*grpc.ClientConn
}

func newConnectionPool() *connectionPool {
	return &connectionPool{
		connections: make(map[string]*grpc.ClientConn),
	}
}

func (p *connectionPool) createIPCClient(addr string, pooled bool) (pb.IPCNodeAPIClient, func() error, error) {
	if pooled {
		conn, err := p.get(addr)
		if err != nil {
			return nil, nil, err
		}
		return pb.NewIPCNodeAPIClient(conn), nil, nil
	}

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}
	return pb.NewIPCNodeAPIClient(conn), conn.Close, nil
}

func (p *connectionPool) createStreamingClient(addr string, pooled bool) (pb.StreamAPIClient, func() error, error) {
	if pooled {
		conn, err := p.get(addr)
		if err != nil {
			return nil, nil, err
		}
		return pb.NewStreamAPIClient(conn), nil, nil
	}

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}
	return pb.NewStreamAPIClient(conn), conn.Close, nil
}

func (p *connectionPool) get(addr string) (*grpc.ClientConn, error) {
	p.mu.RLock()
	if conn, exists := p.connections[addr]; exists {
		p.mu.RUnlock()
		return conn, nil
	}
	p.mu.RUnlock()

	// Lock to prevent race condition
	p.mu.Lock()
	defer p.mu.Unlock()

	// Double-check to see if another goroutine has added the connection
	if conn, exists := p.connections[addr]; exists {
		return conn, nil
	}

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	// Add the new connection to the pool
	p.connections[addr] = conn

	return conn, nil
}

func (p *connectionPool) close() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	var errList []error

	for addr, conn := range p.connections {
		if err := conn.Close(); err != nil {
			errList = append(errList, fmt.Errorf("failed to close connection to %s: %w", addr, err))
		}
		delete(p.connections, addr)
	}

	if len(errList) > 0 {
		return fmt.Errorf("encountered errors while closing connections: %v", errList)
	}

	return nil
}
