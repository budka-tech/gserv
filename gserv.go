package gserv

import (
	"context"
	"fmt"
	"github.com/budka-tech/iport"
	"github.com/budka-tech/logit-go"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Params struct {
	Host    iport.Host
	Port    iport.Port
	Options []grpc.ServerOption
	Logger  logit.Logger
}

type GServ struct {
	listener net.Listener
	host     iport.Host
	port     iport.Port
	options  []grpc.ServerOption
	logger   logit.Logger
}

func NewGServ(params Params) *GServ {
	return &GServ{
		listener: nil,
		host:     params.Host,
		port:     params.Port,
		options:  params.Options,
		logger:   params.Logger,
	}
}

func (g *GServ) Init(ctx context.Context, registerServices func(*grpc.Server)) error {
	const op = "gserv.Init"
	ctx = g.logger.NewOpCtx(ctx, op)

	var err error

	g.listener, err = net.Listen("tcp", iport.FormatLocal(g.port))
	if err != nil {
		log.Fatalf("ошибка прослушивания порта для сервиса %s на порту %d: %v", g.host, g.port, err)
	}
	s := grpc.NewServer(g.options...)

	if registerServices != nil {
		registerServices(s)
	}

	g.logger.Info(ctx, fmt.Sprintf("сервис %s запущен по адресу %d", g.host, g.port))

	go func() {
		if err := s.Serve(g.listener); err != nil {
			err = fmt.Errorf("ошибка при работе сервера %s: %v", g.host, err)
			g.logger.Error(ctx, err)
		}
	}()

	return nil
}

func (g *GServ) Dispose(ctx context.Context) error {
	const op = "gserv.Dispose"
	ctx = g.logger.NewOpCtx(ctx, op)

	if g.listener != nil {
		err := g.listener.Close()
		if err != nil {
			err = fmt.Errorf("ошибка закрытия сервиса %s, слушающего порт %d: %v", g.host, g.port, err)
			return err
		}
		msg := fmt.Sprintf("сервис %s остановлен", g.host)
		g.logger.Info(ctx, msg)
	}

	return nil
}
