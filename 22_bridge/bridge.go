package bridge

import "fmt"

type Message struct {
	sender MessageSender
}

func (msg *Message) SendMessage(message, to string) {
	msg.sender.Send(message, to)
}

type MessageSender interface {
	Send(message, to string)
}

type MessageSMS struct {
}

func (sender *MessageSMS) Send(message, to string) {
	fmt.Printf("send %s to %s via SMS\n", message, to)
}

type MessageEmail struct {
}

func (sender *MessageEmail) Send(message, to string) {
	fmt.Printf("send %s to %s via Email\n", message, to)
}

type CommonMessage struct {
	Message
}

func NewCommonMessage(sender MessageSender) *CommonMessage {
	return &CommonMessage{
		Message{sender: sender},
	}
}

type UrgencyMessage struct {
	Message
}

func NewUrgencyMessage(sender MessageSender) *UrgencyMessage {
	return &UrgencyMessage{
		Message{sender: sender},
	}
}

func (msg *UrgencyMessage) SendMessage(message, to string) {
	message = "[Urgency] " + message
	msg.sender.Send(message, to)
}
