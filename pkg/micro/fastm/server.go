package fastm

import (
	"context"
	"reflect"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/server"
)

type Server struct {
	opts server.Options
	r    *routing.Router
	srv  fasthttp.Server
}

func NewServer(opts ...server.Option) server.Server {
	this := &Server{
		r:    routing.New(),
		opts: server.NewOptions(opts...),
	}

	return this
}

func (s *Server) Init(option ...server.Option) error {
	for _, opt := range option {
		opt(&s.opts)
	}

	return nil
}

func (s *Server) Options() server.Options {
	return s.opts
}

func (s *Server) Handle(handler server.Handler) error {
	// Register metrics route
	s.r.To(
		"GET",
		"/metrics",
		routingHandlerWrapper(fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler())),
	)

	s.r.To(
		"PUT",
		"/netrics",
		routingHandlerWrapper(fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler())),
	)

	// Register any route which was added
	for _, endpoint := range handler.Endpoints() {
		method := endpoint.Metadata[MetaKeyMethod.String()]
		url := endpoint.Metadata[MetaKeyURL.String()]
		getDecoder := s.opts.Codecs["application/json"]
		impl, implFound := reflect.TypeOf(handler.Handler()).MethodByName(endpoint.Name)

		if len(method) == 0 || len(url) == 0 || !implFound {
			s.opts.Logger.Logf(
				logger.ErrorLevel,
				"unable to detect method or URL or implementation for endpoint [%s]",
				endpoint.Name,
			)
			continue
		}

		paramCnt := impl.Type.NumIn()
		reqType := impl.Type.In(paramCnt - requestParamExtractCount)
		respType := impl.Type.In(paramCnt - responseParamExtractCount)

		if reqType.Kind() == reflect.Ptr {
			reqType = reqType.Elem()
		}

		if respType.Kind() == reflect.Ptr {
			respType = respType.Elem()
		}

		var h server.HandlerFunc = func(ctx context.Context, req server.Request, rsp interface{}) error {
			if err := impl.Func.Call([]reflect.Value{
				reflect.ValueOf(handler.Handler()),
				reflect.ValueOf(ctx),
				reflect.ValueOf(req.Body()),
				reflect.ValueOf(rsp),
			})[0].Interface(); err != nil {
				return err.(error)
			}

			return nil
		}

		// Added wrapper for each handler
		for _, wrapper := range s.opts.HdlrWrappers {
			h = wrapper(h)
		}

		s.r.To(method, url, func(context *routing.Context) error {
			resp := reflect.New(respType)

			adapter := newContextAdapter(context)
			decoder := getDecoder(adapter)
			request := NewRequest(context, decoder, reqType)

			if err := h(context, request, resp.Interface()); err != nil {
				return err
			}

			context.SetContentType("application/json")
			return decoder.Write(nil, resp.Interface())
		})
	}

	return nil
}

func (s *Server) NewHandler(i any, option ...server.HandlerOption) server.Handler {
	return NewHandler(i, option...)
}

func (s *Server) NewSubscriber(_ string, _ interface{}, _ ...server.SubscriberOption) server.Subscriber {
	panic("implement me")
}

func (s *Server) Subscribe(_ server.Subscriber) error {
	panic("implement me")
}

func (s *Server) Start() error {
	s.srv.Handler = fasthttp.CompressHandler(s.r.HandleRequest)

	return s.srv.ListenAndServe(s.opts.Address)
}

func (s *Server) Stop() error {
	return s.srv.Shutdown()
}

func (s *Server) String() string {
	if s.opts.TLSConfig != nil {
		return "https"
	}
	return "http"
}
