package recurly

import ()

type AccountAcquisitionUpdatable struct {
	Params `json:"-"`

	// Account balance
	Cost *AccountAcquisitionCostCreate `json:"cost,omitempty"`

	// The channel through which the account was acquired.
	Channel *string `json:"channel,omitempty"`

	// An arbitrary subchannel string representing a distinction/subcategory within a broader channel.
	Subchannel *string `json:"subchannel,omitempty"`

	// An arbitrary identifier for the marketing campaign that led to the acquisition of this account.
	Campaign *string `json:"campaign,omitempty"`
}

func (attr *AccountAcquisitionUpdatable) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
