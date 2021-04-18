package types

// Password is the primary data model.
type Password interface{}

// Meta represents additional data for passwords.
type Meta interface{}

// Reader processes input that determines which utiltiy to invoke.
type Reader interface {
	// Option finds an option within the input that corresponds with the given name.
	// If an option is not found false is returned. Options are values that
	// don't have an explicit value, they are a state that is either true or
	// false.
	Option(name string) bool
	// Argument finds an argument within the input that corresponds with the given name.
	// If an argument is not found false will be returned. Arguments are named
	// values.
	Argument(name string) (string, bool)
	// Arguments fetches all arguments.
	Arguments() map[string]string
	// Value finds a value at the given position. If a value does not exist at the
	// given position then false is returned.
	Value(index int) (string, bool)
	// Values fetches all values.
	Values() []string
}

// Writer handles command output.
type Writer interface {
	// Write outputs given value to the defined output source.
	Write(value string)
}

// Storage manages persisted password data.
type Storage interface {
	// Store persists a new password record if a password with the given name
	// does not already exist.
	Store(name string, password string, meta map[string]string)
	// Show finds a password record with the given name and provides the
	// associated meta data.
	Show(name string) (*Password, []*Meta)
	// Drop delets a password record with the given name.
	Drop(name string)
	// List fetches all passwords.
	List() []Password
}

// Cipher encrypts & decrypts messages.
type Cipher interface {
	// Encrypt encrypts a message.
	Encrypt(msg string) (string, error)
	// Decrypt decrypts a message.
	Decrypt(msg string) (string, error)
	// WriteKeys persists the cryptographic keys.
	WriteKeys(path string)
	// GetKeys fetches the keys used to encrypt and decrypt messages.
	GetKeys() (string, string)
}

// A Command segregates the domain logic.
type Command interface {
	// SetCipher sets the cipher adapter used to encrypt data.
	SetCipher(c Cipher)
	// SetStorage sets the storage adapter used to persist passwords.
	SetStorage(s Storage)
	// Action invokes the command action.
	Action(r Reader, w Writer)
}
