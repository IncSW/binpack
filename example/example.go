package example

type DeviceType uint8

// binpack:gen
type Device struct {
	Type     DeviceType       `binpack:"index=1,type=uint8"`
	Vendor   *KeyValue        `binpack:"index=2"`
	Model    *KeyValue        `binpack:"index=3"`
	OS       *KeyValueVersion `binpack:"index=4"`
	Browser  *KeyValueVersion `binpack:"index=5"`
	Language *KeyValue        `binpack:"index=6"`
}

// binpack:gen
type KeyValue struct {
	ID   uint64 `binpack:"index=1"`
	Name string `binpack:"index=2"`
}

// binpack:gen
type KeyValueVersion struct {
	ID      uint64   `binpack:"index=1"`
	Name    string   `binpack:"index=2"`
	Version *Version `binpack:"index=3"`
}

// binpack:gen
type Version struct {
	Major string `binpack:"index=1"`
	Minor string `binpack:"index=2"`
}
