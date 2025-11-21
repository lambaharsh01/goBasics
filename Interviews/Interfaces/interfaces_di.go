package main

import "fmt"

// ----------------------------
// Step 1: Define the Interface
// ----------------------------
type Notifier interface {
	Send(message string) error   // send a message
	Status() string             // check service status
}

// ----------------------------
// Step 2: Implement EmailNotifier
// ----------------------------
type EmailNotifier struct {
	EmailAddress string
}

func (e *EmailNotifier) Send(message string) error {
	fmt.Println("[Email] Sending to", e.EmailAddress, ":", message)
	return nil
}

func (e *EmailNotifier) Status() string {
	return "Email service OK"
}

// ----------------------------
// Step 2: Implement SMSNotifier
// ----------------------------
type SMSNotifier struct {
	PhoneNumber string
}

func (s *SMSNotifier) Send(message string) error {
	fmt.Println("[SMS] Sending to", s.PhoneNumber, ":", message)
	return nil
}

func (s *SMSNotifier) Status() string {
	return "SMS service OK"
}

// ----------------------------
// Step 3: Create a Service that Depends on Notifier
// ----------------------------
type UserService struct {
	notifier Notifier  // injected dependency
}
// Your struct doesn’t directly depend on a concrete type // If I had defined Send Notification Methods directly to UserService it would not have been called a DI

// It depends on behavior, which is defined by an interface.// Here the User Service Depends On Notifier which further calls Notification Services implemented using Interfaces

// UserService does NOT implement Notifier. Instead, another struct (like EmailNotifier or SMSNotifier) implements Notifier // UserService implements Notifier
// That implementation is injected into your struct from the outside. Notifier is a Dependency

// Constructor-style DI
func NewUserService(n Notifier) *UserService {
	return &UserService{
		notifier: n,
	}
}

// Business method using the injected dependency
func (u *UserService) SendWelcomeMessage(user string) {
	message := fmt.Sprintf("Welcome %s!", user)
	err := u.notifier.Send(message)
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}
	fmt.Println("Notifier Status:", u.notifier.Status())
}

// ----------------------------
// Step 4: Use the Service
// ----------------------------
func InterfaceDIImplementation() {
	// Inject EmailNotifier
	emailNotifier := &EmailNotifier{EmailAddress: "harsh@example.com"}
	userService1 := NewUserService(emailNotifier)
	userService1.SendWelcomeMessage("Harsh")

	fmt.Println("------")

	// Inject SMSNotifier
	smsNotifier := &SMSNotifier{PhoneNumber: "1234567890"}
	userService2 := NewUserService(smsNotifier)
	userService2.SendWelcomeMessage("Raju")
}
