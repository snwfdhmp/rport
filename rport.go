package rport

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/snwfdhmp/errlog"
)

type Reporter interface {
	Report(err error, name string, data ...interface{}) bool
}

type reporter struct {
	remoteURL string
}

func NewReporter(remoteURL string) *reporter {
	if len(remoteURL) > 0 && remoteURL[len(remoteURL)-1] == '/' {
		remoteURL = remoteURL[:len(remoteURL)-1]
	}
	return &reporter{
		remoteURL: remoteURL,
	}
}

func (r *reporter) Report(err error, reportName string, data ...interface{}) (errNotNil bool) {
	errNotNil = err != nil //define return value
	if !errNotNil {        // return early if no err
		return
	}

	var body []byte
	var vErr error
	if len(data) == 1 { // Print object alone if alone
		body, vErr = json.Marshal(data[0])
		if errlog.Debug(vErr) {
			return
		}
	} else { //Otherwise print data as is
		body, vErr = json.Marshal(data)
		if errlog.Debug(vErr) {
			return
		}
	}

	bodyReader := bytes.NewReader(body)
	_, err = http.Post(r.remoteURL+"/"+reportName, "application/json", bodyReader)
	if errlog.Debug(err) {
		return
	}

	return
}
