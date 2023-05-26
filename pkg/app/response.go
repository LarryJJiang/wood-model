package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"reflect"
	"runtime"
	bizcode "wood/pkg/bizerror"
	"wood/pkg/e"
	"wood/pkg/setting"
)

type GinInterface interface {
	// Response setting gin.JSON
	Response(httpCode int, errCode int, data interface{})
	ResponseWithResp(resp *Response)
	ResponseWithErr(err error)
	//Json(data proto.Message)
}

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	CodeKey   = "code"
	DetailKey = "detail"
)

func (r *Response) Error() (msg string) {
	_, msg = r.GetMessage()
	return
}

func (g *Gin) Response(httpCode int, data interface{}) {
	g.C.JSON(http.StatusOK, NewResponse(httpCode, data))
	return
}

func (g *Gin) ResponseWithResp(resp *Response) {
	g.C.JSON(resp.Code, resp)
	return
}

func (g *Gin) ResponseWithErr(err error) {
	if err != nil {
		if r, ok := FromResponseErr(err); ok {
			g.ResponseWithResp(r)
		} else {
			g.Response(bizcode.ERROR, fmt.Sprintf("err: %v", err))
		}
		return
	}
}

//func (g *Gin) JSON(data proto.Message) {
//	// any data
//	anyData, err := ptypes.MarshalAny(data)
//	if err != nil {
//		logging.Error(err)
//		g.ResponseWithErr(NewResponse(http.StatusOK, bizcode.ERROR, nil))
//		return
//	}
//	resp := &pb.Response{
//		Code: http.StatusOK,
//		Msg:  &pb.Msg{Code: e.SUCCESS, Detail: e.GetMsg(e.SUCCESS)},
//		Data: anyData,
//	}
//
//	// marshal
//	var buf = new(bytes.Buffer)
//	if err := JSONHandler.Marshal(buf, resp); err != nil {
//		logging.Error(err)
//		g.ResponseErr(NewResponseErr(http.StatusOK, e.ERROR, nil))
//		return
//	}
//
//	// resp
//	g.C.Data(http.StatusOK, JSONContentType, buf.Bytes())
//	return
//}

// GetMessage code and Msg
func (r *Response) GetMessage() (code int, msg string) {
	if r.Msg == nil {
		return
	}

	// map
	if msgMap, ok := r.Msg.(map[string]interface{}); ok {
		if c, cOk := msgMap[CodeKey]; cOk && c != nil {
			code, _ = c.(int)
		}
		if m, mOK := msgMap[DetailKey]; mOK && m != nil {
			msg, _ = m.(string)
		}
		return
	}

	// string
	if m, ok := r.Msg.(string); ok {
		code = bizcode.SUCCESS
		msg = m
		return
	}

	code = bizcode.SUCCESS
	return
}

// NewResponseErr response to error
func NewResponseErr(errCode int, data interface{}) error {
	// debug
	if setting.IsOutputDebug() {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("file : %s, line(%d), code(%d) \n", file, line, errCode)
	}
	return NewResponse(errCode, data)
}

// NewResponse response to error
func NewResponse(httpCode int, data interface{}) *Response {
	msg := bizcode.GetMsg(httpCode)
	if httpCode != e.SUCCESS {
		dataType := reflect.TypeOf(data)
		dataValue := reflect.ValueOf(data)
		if dataType.Kind().String() == "string" {
			msg = dataValue.String()
			data = nil
		}
	}
	return &Response{
		Code: httpCode,
		Msg:  msg,
		Data: data,
	}
}

// FromResponseErr 解析错误
func FromResponseErr(err error) (*Response, bool) {
	if r, ok := errors.Cause(err).(*Response); ok {
		return r, true
	}
	return nil, false
}
