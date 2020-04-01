package alerts

// Storer interface for payloads
type Storer interface {
	Init() error
	Store(payload Payload) error
	Remove(payload Payload) error
}
