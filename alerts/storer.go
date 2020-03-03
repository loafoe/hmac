package alerts

// Storer interface for payloads
type Storer interface {
	Store(payload Payload) error
	Remove(payoad Payload) error
}
