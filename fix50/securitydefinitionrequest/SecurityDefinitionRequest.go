//Package securitydefinitionrequest msg type = c.
package securitydefinitionrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/instrumentextension"
	"github.com/quickfixgo/quickfix/fix50/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//SubscriptionRequestType is a repeating group in SecurityDefinitionRequest
type SubscriptionRequestType struct {
	//LegOptionRatio is a non-required field for SubscriptionRequestType.
	LegOptionRatio *float64 `fix:"1017"`
	//LegPrice is a non-required field for SubscriptionRequestType.
	LegPrice *float64 `fix:"566"`
}

//Message is a SecurityDefinitionRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"c"`
	Header     fixt11.Header
	//SecurityReqID is a required field for SecurityDefinitionRequest.
	SecurityReqID string `fix:"320"`
	//SecurityRequestType is a required field for SecurityDefinitionRequest.
	SecurityRequestType int `fix:"321"`
	//Instrument Component
	Instrument instrument.Component
	//InstrumentExtension Component
	InstrumentExtension instrumentextension.Component
	//UndInstrmtGrp Component
	UndInstrmtGrp undinstrmtgrp.Component
	//Currency is a non-required field for SecurityDefinitionRequest.
	Currency *string `fix:"15"`
	//Text is a non-required field for SecurityDefinitionRequest.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for SecurityDefinitionRequest.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for SecurityDefinitionRequest.
	EncodedText *string `fix:"355"`
	//TradingSessionID is a non-required field for SecurityDefinitionRequest.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for SecurityDefinitionRequest.
	TradingSessionSubID *string `fix:"625"`
	//InstrmtLegGrp Component
	InstrmtLegGrp instrmtleggrp.Component
	//ExpirationCycle is a non-required field for SecurityDefinitionRequest.
	ExpirationCycle *int `fix:"827"`
	//SubscriptionRequestType is a non-required field for SecurityDefinitionRequest.
	SubscriptionRequestType []SubscriptionRequestType `fix:"263,omitempty"`
	Trailer                 fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.BeginStringFIX50, "c", r
}
