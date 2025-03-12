package fastm

type ctxKey string

const (
	ctxRequestID ctxKey = "request_id"
)

type metaKeys string

func (mk metaKeys) String() string {
	return string(mk)
}

const (
	MetaKeyMethod metaKeys = "method"
	MetaKeyURL    metaKeys = "URL"
)

const (
	requestParamExtractCount  = 2
	responseParamExtractCount = 1
)
