// Package device contains the logic for devices
// including persistence, service and handler functions
// inside it
package device

import "time"

// Device is the "user" of this repo.
// A device can be any microcontroller with wifi
// access that sends some kind of telemetry here
type Device struct {
	ID        uint64    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Password  []byte    `json:"-" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// AsTest clones the device with some modified fields
// so they can be checked on comparison functions without giving errors
func (d Device) AsTest() Device {
	newD := d
	newD.CreatedAt = d.CreatedAt.Round(time.Microsecond).UTC()
	newD.UpdatedAt = d.UpdatedAt.Round(time.Microsecond).UTC()
	return newD
}
