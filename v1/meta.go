package v1

const VERSION = "v1"

type Plugin interface {
	Doctor() error
}

type DoctorInformation map[string]any

type DoctorFunction func() (DoctorInformation, error)
