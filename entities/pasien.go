package entities

type Pasien struct {
	Id           int64
	NamaLengkap  string `validate:"required" label:"nama lengkap"`
	NIK          string `validate:"required" label:"nik"`
	JenisKelamin string `validate:"required" label:"jenis kelamin"`
	TempatLahir  string `validate:"required" label:"tempat lahir"`
	TanggalLahir string `validate:"required" label:" tanggal lahir"`
	Alamat       string `validate:"required" label:"alamat"`
	NoHP         string `validate:"required" label:"no hp"`
}
