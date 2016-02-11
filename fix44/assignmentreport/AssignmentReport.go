//Package assignmentreport msg type = AW.
package assignmentreport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/instrumentleg"
	"github.com/quickfixgo/quickfix/fix44/parties"
	"github.com/quickfixgo/quickfix/fix44/positionamountdata"
	"github.com/quickfixgo/quickfix/fix44/positionqty"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
)

//NoLegs is a repeating group in AssignmentReport
type NoLegs struct {
	//InstrumentLeg Component
	InstrumentLeg instrumentleg.Component
}

//NoUnderlyings is a repeating group in AssignmentReport
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
}

//Message is a AssignmentReport FIX Message
type Message struct {
	FIXMsgType string `fix:"AW"`
	Header     fix44.Header
	//AsgnRptID is a required field for AssignmentReport.
	AsgnRptID string `fix:"833"`
	//TotNumAssignmentReports is a non-required field for AssignmentReport.
	TotNumAssignmentReports *int `fix:"832"`
	//LastRptRequested is a non-required field for AssignmentReport.
	LastRptRequested *bool `fix:"912"`
	//Parties Component
	Parties parties.Component
	//Account is a non-required field for AssignmentReport.
	Account *string `fix:"1"`
	//AccountType is a required field for AssignmentReport.
	AccountType int `fix:"581"`
	//Instrument Component
	Instrument instrument.Component
	//Currency is a non-required field for AssignmentReport.
	Currency *string `fix:"15"`
	//NoLegs is a non-required field for AssignmentReport.
	NoLegs []NoLegs `fix:"555,omitempty"`
	//NoUnderlyings is a non-required field for AssignmentReport.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//PositionQty Component
	PositionQty positionqty.Component
	//PositionAmountData Component
	PositionAmountData positionamountdata.Component
	//ThresholdAmount is a non-required field for AssignmentReport.
	ThresholdAmount *float64 `fix:"834"`
	//SettlPrice is a required field for AssignmentReport.
	SettlPrice float64 `fix:"730"`
	//SettlPriceType is a required field for AssignmentReport.
	SettlPriceType int `fix:"731"`
	//UnderlyingSettlPrice is a required field for AssignmentReport.
	UnderlyingSettlPrice float64 `fix:"732"`
	//ExpireDate is a non-required field for AssignmentReport.
	ExpireDate *string `fix:"432"`
	//AssignmentMethod is a required field for AssignmentReport.
	AssignmentMethod string `fix:"744"`
	//AssignmentUnit is a non-required field for AssignmentReport.
	AssignmentUnit *float64 `fix:"745"`
	//OpenInterest is a required field for AssignmentReport.
	OpenInterest float64 `fix:"746"`
	//ExerciseMethod is a required field for AssignmentReport.
	ExerciseMethod string `fix:"747"`
	//SettlSessID is a required field for AssignmentReport.
	SettlSessID string `fix:"716"`
	//SettlSessSubID is a required field for AssignmentReport.
	SettlSessSubID string `fix:"717"`
	//ClearingBusinessDate is a required field for AssignmentReport.
	ClearingBusinessDate string `fix:"715"`
	//Text is a non-required field for AssignmentReport.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for AssignmentReport.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for AssignmentReport.
	EncodedText *string `fix:"355"`
	Trailer     fix44.Trailer
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
	return enum.BeginStringFIX44, "AW", r
}
