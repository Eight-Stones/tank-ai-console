package fastm

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"

	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"
)

const (
	countHandlerParam = 4
)

type Handler struct {
	handler   any
	name      string
	opts      server.HandlerOptions
	endpoints []*registry.Endpoint
}

func NewHandler(object any, opts ...server.HandlerOption) server.Handler {
	this := &Handler{
		handler:   object,
		endpoints: make([]*registry.Endpoint, 0),
	}

	for _, opt := range opts {
		opt(&this.opts)
	}

	handlerType := reflect.TypeOf(object)
	handler := reflect.ValueOf(object)

	this.handler = object
	this.name = reflect.Indirect(handler).Type().Name()

	for i := 0; i < handlerType.NumMethod(); i++ {
		m := handlerType.Method(i)
		paramCnt := m.Type.NumIn()

		if paramCnt != countHandlerParam {
			logger.Errorf("method with name [%s] has [%d] parameters but 3 are required", m.Name, m.Type.NumIn())
			continue
		}

		e := &registry.Endpoint{
			Request:  this.extractSchema(m.Type.In(paramCnt - requestParamExtractCount)),
			Response: this.extractSchema(m.Type.In(paramCnt - responseParamExtractCount)),
			Metadata: this.extractMeta(m.Name),
			Name:     m.Name,
		}

		this.endpoints = append(this.endpoints, e)
	}

	return this
}

func (h Handler) Name() string {
	return h.name
}

func (h Handler) Handler() interface{} {
	return h.handler
}

func (h Handler) Endpoints() []*registry.Endpoint {
	return h.endpoints
}

func (h Handler) Options() server.HandlerOptions {
	return h.opts
}

func (h Handler) extractMeta(name string) map[string]string {
	const (
		prefixPost   = "Post"
		prefixPatch  = "Patch"
		prefixPut    = "Put"
		prefixDelete = "Delete"
		prefixGet    = "Get"
	)

	result := make(map[string]string)

	// Extract method from method name
	switch {
	case strings.HasPrefix(name, prefixPost):
		result[MetaKeyMethod.String()] = "POST"
		name = strings.TrimPrefix(name, prefixPost)
	case strings.HasPrefix(name, prefixPatch):
		result[MetaKeyMethod.String()] = "PATCH"
		name = strings.TrimPrefix(name, prefixPatch)
	case strings.HasPrefix(name, prefixPut):
		result[MetaKeyMethod.String()] = "PUT"
		name = strings.TrimPrefix(name, prefixPut)
	case strings.HasPrefix(name, prefixDelete):
		result[MetaKeyMethod.String()] = "DELETE"
		name = strings.TrimPrefix(name, prefixDelete)
	default:
		result[MetaKeyMethod.String()] = "GET"
		name = strings.TrimPrefix(name, prefixGet)
	}

	// Extract URL
	path := make([]string, 0)
	prevIdx := 0

	for idx, sym := range name {
		if unicode.IsUpper(sym) && unicode.IsLetter(sym) {
			path = append(path, strings.ToLower(name[prevIdx:idx]))
			prevIdx = idx
		}
	}
	path = append(path, strings.ToLower(name[prevIdx:]))

	pattern := make([]string, 0, len(path))
	for i := 0; i < len(path); i++ {
		if path[i] == "param" && (i+1) < len(path) {
			pattern = append(pattern, fmt.Sprintf("<%s>", path[i+1]))
			i++
		} else {
			pattern = append(pattern, path[i])
		}
	}

	result[MetaKeyURL.String()] = strings.Join(pattern, "/")

	return result
}

func (h Handler) extractSchema(t reflect.Type) *registry.Value { //nolint:gocognit // framework needs
	if t == nil {
		return nil
	}

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	arg := &registry.Value{
		Name: t.Name(),
		Type: t.Name(),
	}

	switch t.Kind() { //nolint:exhaustive // does not need at this moment
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			if f.PkgPath != "" {
				continue
			}

			val := h.extractSchema(f.Type)
			if val == nil {
				continue
			}

			// If we can find a JSON tag use it
			if tags := f.Tag.Get("json"); len(tags) > 0 {
				parts := strings.Split(tags, ",")
				if parts[0] == "-" || parts[0] == "omitempty" {
					continue
				}
				val.Name = parts[0]
			}

			// If there's no name default it
			if len(val.Name) == 0 {
				val.Name = t.Field(i).Name
			}

			// Still no name then continue
			if len(val.Name) == 0 {
				continue
			}

			arg.Values = append(arg.Values, val)
		}
	case reflect.Slice:
		p := t.Elem()
		if p.Kind() == reflect.Ptr {
			p = p.Elem()
		}
		arg.Type = "[]" + p.Name()
	}

	return arg
}
