package main

import (
	svc "Testprj03/hometest.globalgroup.com/backend/service"
	"testing"
)

func TestGetAccounts(t *testing.T) {
	p := svc.PrettyServiceType{}
	if _, err := p.GetAccounts(); err != nil {
		t.Errorf("fetch accounts error:%#v", err)
	}
}

func TestGetPayments(t *testing.T) {
	p := svc.PrettyServiceType{}
	if _, err := p.GetAllPayments(); err != nil {
		t.Errorf("fetch payments error:%#v", err)
	}
}
