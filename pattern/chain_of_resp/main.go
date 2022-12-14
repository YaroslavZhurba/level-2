package main

import (
	"fmt"
)

type Priority int64

const (
	ROUTINE Priority = iota
	IMPORTANT
	ASAP
)

type Notifier interface {
	SetNextNotifier(Notifier)
	NotifyManager(string, Priority)
}

type BaseNotifier struct {
	priority     Priority
	nextNotifier Notifier
}

func (notifier *BaseNotifier) SetNextNotifier(nextNotifier Notifier) {
	notifier.nextNotifier = nextNotifier
}

func (notifier *BaseNotifier) NotifyManager(message string, priority Priority) {
	if notifier.priority <= priority {
		notifier.write(message)
	}

	if notifier.nextNotifier != nil {
		notifier.nextNotifier.NotifyManager(message, priority)
	}
}

func (notifier *BaseNotifier) write(message string) {
	fmt.Println(message)
}

type SimpleReportNotifier struct {
	BaseNotifier
}

func (notifier *SimpleReportNotifier) NotifyManager(message string, priority Priority) {
	notifier.BaseNotifier.NotifyManager("Sending simple report: "+message, priority)
}

type EmailNotifier struct {
	BaseNotifier
}

func (notifier *EmailNotifier) NotifyManager(message string, priority Priority) {
	notifier.BaseNotifier.NotifyManager("Sending email: "+message, priority)
}

type SmsNotifier struct {
	BaseNotifier
}

func (notifier *SmsNotifier) NotifyManager(message string, priority Priority) {
	notifier.BaseNotifier.NotifyManager("Sending sms: "+message, priority)
}

func main() {
	reportNotifier := &SimpleReportNotifier{BaseNotifier: BaseNotifier{priority: ROUTINE}}
	emailNotifier := &EmailNotifier{BaseNotifier: BaseNotifier{priority: IMPORTANT}}
	smsNotifier := &SmsNotifier{BaseNotifier: BaseNotifier{priority: ASAP}}

	reportNotifier.SetNextNotifier(emailNotifier)
	emailNotifier.SetNextNotifier(smsNotifier)

	reportNotifier.NotifyManager("Task num 5 is done", ROUTINE)
	fmt.Println()
	reportNotifier.NotifyManager("Bug has founded..", IMPORTANT)
	fmt.Println()
	reportNotifier.NotifyManager("Server has fallen!!!", ASAP)
}
