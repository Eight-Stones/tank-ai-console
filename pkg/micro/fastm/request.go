package fastm

import (
	"reflect"
	"strconv"

	routing "github.com/qiangxue/fasthttp-routing"
	"go-micro.dev/v4/codec"
)

type Request struct {
	ctx     *routing.Context
	reader  codec.Reader
	reqType reflect.Type
}

func NewRequest(ctx *routing.Context, reader codec.Reader, reqType reflect.Type) *Request {
	return &Request{ctx: ctx, reader: reader, reqType: reqType}
}

func (r Request) Service() string {
	return string(r.ctx.Request.Header.Peek("Fastmicro-Service"))
}

func (r Request) Method() string {
	return string(r.ctx.Method())
}

func (r Request) Endpoint() string {
	return string(r.ctx.Request.RequestURI())
}

func (r Request) ContentType() string {
	return string(r.ctx.Request.Header.Peek("Content-Type"))
}

func (r Request) Header() map[string]string {
	result := make(map[string]string)

	r.ctx.Request.Header.VisitAll(func(key, value []byte) {
		result[string(key)] = string(value)
	})

	return result
}

func (r Request) Body() any { //nolint:gocognit // framework needs
	req := reflect.New(r.reqType)
	obj := req.Elem()

	_ = r.reader.ReadBody(req.Interface())

	for i := 0; i < r.reqType.NumField(); i++ {
		f := r.reqType.Field(i)
		t, err := newTag(f.Tag.Get("fastmicro"))
		if err != nil {
			continue
		}

		var val string

		switch t.Location() {
		case path:
			val = r.ctx.Param(t.name)
		case query:
			val = string(r.ctx.QueryArgs().Peek(t.name))
		}

		if len(val) == 0 {
			continue
		}

		switch obj.FieldByName(f.Name).Kind() { //nolint:exhaustive // too much for this code
		case reflect.Int64:
			if v, errParse := strconv.ParseInt(val, 10, 64); errParse == nil {
				obj.FieldByName(f.Name).SetInt(v)
			}
		case reflect.Int:
			if v, errParse := strconv.ParseInt(val, 10, 32); errParse == nil {
				obj.FieldByName(f.Name).SetInt(v)
			}
		case reflect.String:
			obj.FieldByName(f.Name).SetString(val)
		case reflect.Bool:
			if v, errParse := strconv.ParseBool(val); errParse == nil {
				obj.FieldByName(f.Name).SetBool(v)
			}
		case reflect.Float64:
			if v, errParse := strconv.ParseFloat(val, 64); errParse == nil {
				obj.FieldByName(f.Name).SetFloat(v)
			}
		case reflect.Float32:
			if v, errParse := strconv.ParseFloat(val, 32); errParse == nil {
				obj.FieldByName(f.Name).SetFloat(v)
			}
		}
	}

	return req.Interface()
}

func (r Request) Read() ([]byte, error) {
	return r.ctx.Request.Body(), nil
}

func (r Request) Codec() codec.Reader {
	return r.reader
}

func (r Request) Stream() bool {
	return r.ctx.IsBodyStream()
}
