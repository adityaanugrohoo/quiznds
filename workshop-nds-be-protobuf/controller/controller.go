package controller

import (
	"proto-workshop/client"

	"bitbucket.bri.co.id/nds/nds-go-module-logger/interfaces"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ControllerInterface interface {
	HitungLuasPersegi(panjang int32, lebar int32) (int32, error)
	HitungVolumeKubus(sisi int32) (int32, error)
	HitungUmur(sisi int32) (int32, error)
	HitungKelilingPersegi(panjang int32, lebar int32) (int32, error)
	HitungLingkaran(radius float64) (float64, float64, error)
	HitungSegitiga(alas int32, tinggi int32, sisi1 int32, sisi2 int32, sisi3 int32) (int32, int32, error)
}

type ControllerStruct struct {
	logger interfaces.Logger
	cl     client.ClientInterFace
}

type ControlleStruct struct {
	logger interfaces.Logger
}

func NewController(logger interfaces.Logger, cl client.ClientInterFace) ControllerInterface {
	return &ControllerStruct{
		logger: logger,
		cl:     cl,
	}
}

func (c *ControllerStruct) HitungLuasPersegi(panjang int32, lebar int32) (int32, error) {
	if panjang < 1 {
		c.logger.Error("Error", "Panjang Harus Lebih Dri 0")
		return 0, status.Error(codes.FailedPrecondition, "Panjang harus Lebih Dari 0")
	}
	if lebar < 1 {
		c.logger.Error("Error", "Lebar Harus Lebih Dri 0")
		return 0, status.Error(codes.FailedPrecondition, "Lebar harus Lebih Dari 0")
	}
	c.logger.Info("PANJANG SISI", panjang)
	c.logger.Info("LEBAR SISI", lebar)
	luas := c.cl.HitungLuasPersegi(panjang, lebar)
	c.logger.Info("LUAS", luas)
	return luas, nil
}

func (c *ControllerStruct) HitungVolumeKubus(sisi int32) (int32, error) {
	if sisi < 1 {
		c.logger.Error("Error", "Panjang Harus Lebih Dri 0")
		return 0, status.Error(codes.FailedPrecondition, "Panjang harus Lebih Dari 0")
	}
	c.logger.Info("PANJANG SISI", sisi)
	volume := c.cl.HitungVolumeKubus(sisi)
	c.logger.Info("LUAS", volume)
	return volume, nil
}

func (c *ControllerStruct) HitungKelilingPersegi(panjang int32, lebar int32) (int32, error) {
	if panjang < 1 {
		c.logger.Error("Error", "Panjang Harus Lebih Dri 0")
		return 0, status.Error(codes.FailedPrecondition, "Panjang harus Lebih Dari 0")
	}
	if lebar < 1 {
		c.logger.Error("Error", "Lebar Harus Lebih Dri 0")
		return 0, status.Error(codes.FailedPrecondition, "Lebar harus Lebih Dari 0")
	}
	c.logger.Info("PANJANG SISI", panjang)
	c.logger.Info("LEBAR SISI", lebar)
	keliling := c.cl.HitungKelilingPersegi(panjang, lebar)
	c.logger.Info("LUAS", keliling)
	return keliling, nil
}

func (c *ControllerStruct) HitungUmur(tanggalLahir int32) (int32, error) {
	if tanggalLahir == 0 {
		c.logger.Error("Error", "Format Tanggal Lahir tidak sesuai")
		return 0, status.Error(codes.FailedPrecondition, "Format Tanggal Lahir tidak sesuai")
	}
	c.logger.Info("TANGGAL LAHIR", tanggalLahir)
	volume := c.cl.HitungUmur(tanggalLahir)

	c.logger.Info("LUAS", volume)
	return volume, nil
}

func (c *ControllerStruct) HitungLingkaran(radius float64) (float64, float64, error) {
	if radius == 0 {
		c.logger.Error("Error", "radius tidak boleh kosong")
		return 0, 0, status.Error(codes.FailedPrecondition, "radius tidak boleh kosong")
	}

	luas, keliling := c.cl.HitungLingkaran(radius)
	return luas, keliling, nil
}

func (c *ControllerStruct) HitungSegitiga(alas int32, tinggi int32, sisi1 int32, sisi2 int32, sisi3 int32) (int32, int32, error) {
	if alas == 0 {
		c.logger.Error("Error", "alas tidak boleh kosong")
		return 0, 0, status.Error(codes.FailedPrecondition, "alas tidak boleh kosong")
	}
	if tinggi == 0 {
		c.logger.Error("Error", "tinggi tidak boleh kosong")
		return 0, 0, status.Error(codes.FailedPrecondition, "tinggi tidak boleh kosong")
	}
	if sisi1 == 0 {
		c.logger.Error("Error", "sisi tidak boleh kosong")
		return 0, 0, status.Error(codes.FailedPrecondition, "sisi tidak boleh kosong")
	}

	luas, keliling := c.cl.HitungSegitiga(alas, tinggi, sisi1, sisi2, sisi3)
	return luas, keliling, nil
}
