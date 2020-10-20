package recurly

import ()

type MeasuredUnitUpdate struct {
	Params `json:"-"`

	// Unique internal name of the measured unit on your site.
	Name *string `json:"name,omitempty"`

	// Display name for the measured unit.
	DisplayName *string `json:"display_name,omitempty"`

	// Optional internal description.
	Description *string `json:"description,omitempty"`
}

func (attr *MeasuredUnitUpdate) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
