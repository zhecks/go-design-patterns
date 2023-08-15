package chainofresponsibility

import "testing"

func TestChainofresonsibility(t *testing.T) {
	req1 := NewGeneralManagerRequest()
	req2 := NewProjectManagerRequest()
	req3 := NewDepManagerRequest()
	req1.SetSuccessor(req2)
	req2.SetSuccessor(req3)

	handle := Handler(req1)
	handle.HandleFeeRequest(100)
	handle.HandleFeeRequest(600)
	handle.HandleFeeRequest(1100)
}
