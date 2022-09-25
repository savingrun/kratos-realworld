package server

import (
	"kratos-realworld/internal/errors"
	netHttp "net/http"

	"github.com/go-kratos/kratos/v2/transport/http"
)

func errorEncoder(w netHttp.ResponseWriter, r *netHttp.Request, err error) {
	se := errors.FromError(err)
	codec, _ := http.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(se)
	if err != nil {
		w.WriteHeader(netHttp.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/"+codec.Name())
	if se.Code >= netHttp.StatusContinue || se.Code <= netHttp.StatusNetworkAuthenticationRequired {
		w.WriteHeader(se.Code)
	} else {
		w.WriteHeader(netHttp.StatusInternalServerError)
	}
	_, _ = w.Write(body)
}
