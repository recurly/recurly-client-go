package recurly

import ()

type SubscriptionPause struct {
	Params `json:"-"`

	// Number of billing cycles to pause the subscriptions.
	RemainingPauseCycles *int `json:"remaining_pause_cycles,omitempty"`
}

func (attr *SubscriptionPause) toParams() *Params {
	return &Params{
		IdempotencyKey: attr.IdempotencyKey,
		Header:         attr.Header,
		Context:        attr.Context,
		Data:           attr,
	}
}
