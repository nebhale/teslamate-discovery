// Copyright 2022 Ben Hale
// SPDX-License-Identifier: Apache-2.0

package mqtt_test

type stubMessage struct {
	duplicate bool
	qos       byte
	retained  bool
	topic     string
	messageId uint16
	payload   []byte
	ack       bool
}

func (s *stubMessage) Duplicate() bool {
	return s.duplicate
}

func (s *stubMessage) Qos() byte {
	return s.qos
}

func (s *stubMessage) Retained() bool {
	return s.retained
}

func (s *stubMessage) Topic() string {
	return s.topic
}

func (s *stubMessage) MessageID() uint16 {
	return s.messageId
}

func (s *stubMessage) Payload() []byte {
	return s.payload
}

func (s *stubMessage) Ack() {
	s.ack = true
}
