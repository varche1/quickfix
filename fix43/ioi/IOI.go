//Package ioi msg type = 6.
package ioi

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix43/instrument"
	"github.com/quickfixgo/quickfix/fix43/spreadorbenchmarkcurvedata"
	"time"
)

//NoIOIQualifiers is a repeating group in IOI
type NoIOIQualifiers struct {
	//IOIQualifier is a non-required field for NoIOIQualifiers.
	IOIQualifier *string `fix:"104"`
}

//NoRoutingIDs is a repeating group in IOI
type NoRoutingIDs struct {
	//RoutingType is a non-required field for NoRoutingIDs.
	RoutingType *int `fix:"216"`
	//RoutingID is a non-required field for NoRoutingIDs.
	RoutingID *string `fix:"217"`
}

//Message is a IOI FIX Message
type Message struct {
	FIXMsgType string `fix:"6"`
	Header     fix43.Header
	//IOIid is a required field for IOI.
	IOIid string `fix:"23"`
	//IOITransType is a required field for IOI.
	IOITransType string `fix:"28"`
	//IOIRefID is a non-required field for IOI.
	IOIRefID *string `fix:"26"`
	//Instrument Component
	Instrument instrument.Component
	//Side is a required field for IOI.
	Side string `fix:"54"`
	//QuantityType is a non-required field for IOI.
	QuantityType *int `fix:"465"`
	//IOIQty is a required field for IOI.
	IOIQty string `fix:"27"`
	//PriceType is a non-required field for IOI.
	PriceType *int `fix:"423"`
	//Price is a non-required field for IOI.
	Price *float64 `fix:"44"`
	//Currency is a non-required field for IOI.
	Currency *string `fix:"15"`
	//ValidUntilTime is a non-required field for IOI.
	ValidUntilTime *time.Time `fix:"62"`
	//IOIQltyInd is a non-required field for IOI.
	IOIQltyInd *string `fix:"25"`
	//IOINaturalFlag is a non-required field for IOI.
	IOINaturalFlag *bool `fix:"130"`
	//NoIOIQualifiers is a non-required field for IOI.
	NoIOIQualifiers []NoIOIQualifiers `fix:"199,omitempty"`
	//Text is a non-required field for IOI.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for IOI.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for IOI.
	EncodedText *string `fix:"355"`
	//TransactTime is a non-required field for IOI.
	TransactTime *time.Time `fix:"60"`
	//URLLink is a non-required field for IOI.
	URLLink *string `fix:"149"`
	//NoRoutingIDs is a non-required field for IOI.
	NoRoutingIDs []NoRoutingIDs `fix:"215,omitempty"`
	//SpreadOrBenchmarkCurveData Component
	SpreadOrBenchmarkCurveData spreadorbenchmarkcurvedata.Component
	//Benchmark is a non-required field for IOI.
	Benchmark *string `fix:"219"`
	Trailer   fix43.Trailer
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
	return enum.BeginStringFIX43, "6", r
}
