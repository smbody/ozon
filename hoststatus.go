package main

import (
	"fmt"
	"net/http"
)

type HostStatus struct {
	host string
	err  error
	resp *http.Response
}

func (h *HostStatus) ToString() (s string) {
	if h.err != nil {
		s = fmt.Sprintf("[Error] %s unreachable. Error:  %s", h.host, h.err.Error())
	} else {
		s = fmt.Sprintf("[Debug] %s accesible. Response status: %s", h.host, h.resp.Status)
	}
	return
}
