package endless_test

import (
	"net/http"
	"syscall"
	"testing"

	"github.com/fvbock/endless"
)

func TestHandleBeforeFunc(t *testing.T) {
	srv := endless.NewServer("", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	srv.HandleBeforeFunc(syscall.SIGHUP, func() {})
	if l := len(srv.SignalHooks[endless.PRE_SIGNAL][syscall.SIGHUP]); l != 1 {
		t.Fatalf("Expected 1 SIGHUB hook, got %v", l)
	}
	if l := len(srv.SignalHooks[endless.PRE_SIGNAL][syscall.SIGUSR1]); l != 0 {
		t.Fatalf("Expected 0 SIGUSR1 hook, got %v", l)
	}

	srv.HandleBeforeFunc(syscall.SIGHUP, func() {})
	if l := len(srv.SignalHooks[endless.PRE_SIGNAL][syscall.SIGHUP]); l != 2 {
		t.Fatalf("Expected 2 SIGHUP hook, got %v", l)
	}
}

func TestHandleAfterFunc(t *testing.T) {
	srv := endless.NewServer("", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	srv.HandleAfterFunc(syscall.SIGUSR1, func() {})
	if l := len(srv.SignalHooks[endless.POST_SIGNAL][syscall.SIGUSR1]); l != 1 {
		t.Fatalf("Expected 1 SIGUSR1 hook, got %v", l)
	}

	srv.HandleAfterFunc(syscall.SIGUSR1, func() {})
	if l := len(srv.SignalHooks[endless.POST_SIGNAL][syscall.SIGUSR1]); l != 2 {
		t.Fatalf("Expected 2 SIGUSR1 hook, got %v", l)
	}
}
