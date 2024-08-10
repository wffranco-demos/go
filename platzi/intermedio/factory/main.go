package main

import (
	"fmt"
)

// SMS and Email Notification

type NotificationFactory interface {
	SendNotification()
	GetSender() Sender
}

type Sender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type SMSNotification struct {
}

type SMSNotificationSender struct {
}

func (SMSNotification) SendNotification() {
	fmt.Println("Sending Notification via SMS")
}

func (SMSNotification) GetSender() Sender {
	return SMSNotificationSender{}
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

type EmailNotification struct {
}

type EmailNotificationSender struct {
}

func (EmailNotification) SendNotification() {
	fmt.Println("Sending Notification via Email")
}

func (EmailNotification) GetSender() Sender {
	return EmailNotificationSender{}
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

func getNotificationFactory(notificationType string) (NotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}

	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}

	return nil, fmt.Errorf("no notification type with name %s", notificationType)
}

func sendNotification(f NotificationFactory) {
	f.SendNotification()
}

func getMethod(f NotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	smsNotificationFactory, _ := getNotificationFactory("SMS")
	sendNotification(smsNotificationFactory)
	getMethod(smsNotificationFactory)

	emailNotificationFactory, _ := getNotificationFactory("Email")
	sendNotification(emailNotificationFactory)
	getMethod(emailNotificationFactory)
}
