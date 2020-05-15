package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// TODO: Describe your actions, these will implment the interface of `sdk.Msg`
// msg offered a template for the message create/commit etc programs <action> is the create scavenge

// verify interface at compile time
var _ sdk.Msg = &MsgCreateScavange{}

// Msg<Action> - struct for unjailing jailed validator
type MsgCreateScavange struct {
	Creator       sdk.AccAddress  `json:"creator" yaml:"creator"` // address of the validator operator
	Description   string      	 `json:"description" yaml:"description"`
	SolutionHash  string 				 `json:"solutionHash" yaml:"solutionHash"`
	Reward				sdk.Coins 		 `json:"reward" yaml:"reward"`
}

// NewMsg<Action> creates a new Msg<Action> instance
func NewMsgCreateScavange(creator sdk.AccAddress,description,solutionHash string, reward sdk.Coins) MsgCreateScavange {
	return MsgCreateScavange {
		Creator:       creator,
		Description:   description,
		SolutionHash:  solutionHash,
		Reward:        reward,
	}
}

const CreateScavangeConst = "CreateScavange"

// nolint
func (msg MsgCreateScavange) Route() string { return RouterKey }
func (msg MsgCreateScavange) Type() string  { return CreateScavangeConst }
func (msg MsgCreateScavange) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgCreateScavange) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgCreateScavange) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	if msg.SolutionHash = "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "solutionScavengerHash can't be empty")
	}
	return nil
}
