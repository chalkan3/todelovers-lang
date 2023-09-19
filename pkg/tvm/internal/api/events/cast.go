package events

import "mary_guica/pkg/nando"

func toNotifyRequest(r *nando.Request) *NotifyRequest { return r.Data.(*NotifyRequest) }
func toCreateHandlerRequest(r *nando.Request) *CreateHandlerRequest {
	return r.Data.(*CreateHandlerRequest)
}
