// Copyright 2022 Ben Hale
// SPDX-License-Identifier: Apache-2.0

package mqtt_test

import paho "github.com/eclipse/paho.mqtt.golang"

type stubPubSub struct {
	publishArgs       []publishArgs
	publishTokens     []paho.Token
	subscribeArgs     []subscribeArgs
	subscribeHandler  func(callback paho.MessageHandler)
	subscribeTokens   []paho.Token
	unsubscribeArgs   []unsubscribeArgs
	unsubscribeTokens []paho.Token
}

type publishArgs struct {
	topic    string
	qos      byte
	retained bool
	payload  interface{}
}

func (s *stubPubSub) Publish(topic string, qos byte, retained bool, payload interface{}) paho.Token {
	s.publishArgs = append(s.publishArgs, publishArgs{
		topic:    topic,
		qos:      qos,
		retained: retained,
		payload:  payload,
	})

	count := len(s.publishArgs) - 1
	if count < len(s.publishTokens) {
		return s.publishTokens[count]
	}
	if len(s.publishTokens) == 1 {
		return s.publishTokens[0]
	}
	return unexpectedOperationToken
}

type subscribeArgs struct {
	topic    string
	qos      byte
	callback paho.MessageHandler
}

func (s *stubPubSub) Subscribe(topic string, qos byte, callback paho.MessageHandler) paho.Token {
	s.subscribeArgs = append(s.subscribeArgs, subscribeArgs{
		topic:    topic,
		qos:      qos,
		callback: callback,
	})

	if s.subscribeHandler != nil {
		go s.subscribeHandler(callback)
	}

	count := len(s.subscribeArgs)
	if count < len(s.subscribeTokens) {
		return s.subscribeTokens[count]
	}
	if len(s.subscribeTokens) == 1 {
		return s.subscribeTokens[0]
	}
	return unexpectedOperationToken
}

type unsubscribeArgs struct {
	topics []string
}

func (s *stubPubSub) Unsubscribe(topics ...string) paho.Token {
	s.unsubscribeArgs = append(s.unsubscribeArgs, unsubscribeArgs{
		topics: topics,
	})

	count := len(s.unsubscribeArgs) - 1
	if count < len(s.unsubscribeTokens) {
		return s.unsubscribeTokens[count]
	}
	if len(s.unsubscribeTokens) == 1 {
		return s.unsubscribeTokens[0]
	}
	return unexpectedOperationToken
}
