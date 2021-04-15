package device

import "time"

type Device struct {
	ID        uint64    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Password  []byte    `json:"-" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (d Device) AsTest() Device {
	newD := d
	newD.CreatedAt = d.CreatedAt.Round(time.Microsecond).UTC()
	newD.UpdatedAt = d.UpdatedAt.Round(time.Microsecond).UTC()
	return newD
}
