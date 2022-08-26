// Copyright 2022 Ben Hale
// SPDX-License-Identifier: Apache-2.0

package mqtt_test

import (
	"fmt"
	"time"
)

var unexpectedOperationToken = &stubToken{
	err: fmt.Errorf("unexpected operation"),
}

type stubToken struct {
	err error
}

func (s *stubToken) Wait() bool {
	return true
}

func (s *stubToken) WaitTimeout(t time.Duration) bool {
	return true
}

func (s *stubToken) Done() <-chan struct{} {
	ch := make(chan struct{})
	close(ch)
	return ch
}

func (s *stubToken) Error() error {
	return s.err
}
