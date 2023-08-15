package bridge

import "testing"

func TestBridge(t *testing.T) {
	sms := MessageSMS{}
	email := MessageEmail{}

	commonMessage := NewCommonMessage(&sms)
	commonMessage.SendMessage("hello Alice", "Bob")
	commonMessage = NewCommonMessage(&email)
	commonMessage.SendMessage("hello Alice", "Bob")
	urgencyMessage := NewUrgencyMessage(&sms)
	urgencyMessage.SendMessage("hello Alice", "Bob")
	urgencyMessage = NewUrgencyMessage(&email)
	urgencyMessage.SendMessage("hello Alice", "Bob")
}
