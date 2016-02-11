package nstdptys2subgrp

//NoNested2PartySubIDs is a repeating group in NstdPtys2SubGrp
type NoNested2PartySubIDs struct {
	//Nested2PartySubID is a non-required field for NoNested2PartySubIDs.
	Nested2PartySubID *string `fix:"760"`
	//Nested2PartySubIDType is a non-required field for NoNested2PartySubIDs.
	Nested2PartySubIDType *int `fix:"807"`
}

//Component is a fix50 NstdPtys2SubGrp Component
type Component struct {
	//NoNested2PartySubIDs is a non-required field for NstdPtys2SubGrp.
	NoNested2PartySubIDs []NoNested2PartySubIDs `fix:"806,omitempty"`
}

func New() *Component { return new(Component) }
