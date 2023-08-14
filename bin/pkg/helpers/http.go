package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/Padliwinata/iam-mfa/bin/pkg/errors"

	"go.elastic.co/apm/module/apmhttp"
	"golang.org/x/net/context/ctxhttp"
)

type HttpPostFormRequestPayload struct {
	Url      string
	FormData url.Values
	Result   interface{}
}

func HttpPostFormRequest(payload HttpPostFormRequestPayload, ctx context.Context) error {
	req, err := http.NewRequest("POST", payload.Url, strings.NewReader(payload.FormData.Encode()))

	if err != nil {
		return errors.InternalServerError(err.Error())
	}

	req.Close = true
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// --- do request

	newClient := http.Client{
		Timeout: 15 * time.Second,
	}

	var wrapClient = apmhttp.WrapClient(&newClient)

	resp, err := ctxhttp.Do(ctx, wrapClient, req)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		return errors.InternalServerError("request timeout 10s.")
	}

	if err != nil {
		return errors.InternalServerError(err.Error())
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.InternalServerError("request error.")
	}

	readResp, err := io.ReadAll(resp.Body)

	if err != nil {
		return errors.InternalServerError(err.Error())
	}

	if err := json.Unmarshal(readResp, &payload.Result); err != nil {
		return errors.InternalServerError("cannot marshal response payload")
	}

	return nil
}

type HttpGetFormRequestPayload struct {
	Url      string
	FormData url.Values
	Token    string
	Result   interface{}
}

func HttpGetFormRequest(payload HttpGetFormRequestPayload, ctx context.Context) error {
	req, err := http.NewRequest("GET", payload.Url, strings.NewReader(payload.FormData.Encode()))

	if err != nil {
		return errors.InternalServerError(err.Error())
	}

	req.Close = true
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", payload.Token))
	// --- do request

	newClient := http.Client{
		Timeout: 5 * time.Second,
	}

	var wrapClient = apmhttp.WrapClient(&newClient)

	resp, err := ctxhttp.Do(ctx, wrapClient, req)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		return errors.InternalServerError("request timeout 5s.")
	}

	if err != nil {
		return errors.InternalServerError(err.Error())
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.InternalServerError("request error.")
	}

	readResp, err := io.ReadAll(resp.Body)

	if err != nil {
		return errors.InternalServerError(err.Error())
	}

	if err := json.Unmarshal(readResp, &payload.Result); err != nil {
		return errors.InternalServerError("cannot marshal response payload")
	}

	return nil
}
