package client

import (
	"math"
	"strconv"
	"time"

	"bitbucket.bri.co.id/nds/nds-go-module-logger/interfaces"
)

type ClientInterFace interface {
	HitungLuasPersegi(panjang int32, lebar int32) int32
	HitungVolumeKubus(sisi int32) int32
	HitungUmur(tanggalLahir int32) int32
	HitungKelilingPersegi(panjang int32, lebar int32) int32
	HitungLingkaran(radius float64) (float64, float64)
	HitungSegitiga(alas int32, tinggi int32, sisi1 int32, sisi2 int32, sisi3 int32) (int32, int32)
}

type ClientStruct struct {
	logger interfaces.Logger
}

func NewClientStruct(logger interfaces.Logger) ClientInterFace {
	return &ClientStruct{
		logger: logger,
	}
}

func (c *ClientStruct) HitungVolumeKubus(sisi int32) int32 {
	c.logger.Info("PANJANG SISI", sisi)
	volume := sisi * sisi * sisi
	c.logger.Info("VOLUMENIHHHHH", volume)
	return volume
}

func (c *ClientStruct) HitungKelilingPersegi(panjang int32, lebar int32) int32 {
	c.logger.Info("PANJANG SISI", panjang)
	c.logger.Info("LEBAR SISI", lebar)
	keliling := panjang + lebar + panjang + lebar
	c.logger.Info("KELILING NIH", keliling)
	return keliling
}

func (c *ClientStruct) HitungLuasPersegi(panjang int32, lebar int32) int32 {
	c.logger.Info("PANJANG SISI", panjang)
	c.logger.Info("LEBAR SISI", lebar)
	luas := panjang * lebar
	c.logger.Info("LUAS", luas)
	return luas
}

func (c *ClientStruct) HitungUmur(tanggalLahir int32) int32 {
	c.logger.Info("Tanggal Lahir", tanggalLahir)
	tanggalLahirString := strconv.Itoa(int(tanggalLahir))
	tanggalLahirDate, err := time.Parse("02012006", tanggalLahirString)

	if err != nil {
		return 0
	}

	timenow := time.Now()
	Age := timenow.Year() - tanggalLahirDate.Year()

	c.logger.Info("Umur", Age)
	return int32(Age)
}

func (c *ClientStruct) HitungLingkaran(radius float64) (float64, float64) {
	c.logger.Info("RADIUS LINGKARAN", radius)
	const pi float64 = math.Pi

	luas := pi * radius * radius
	keliling := 2 * pi * radius
	c.logger.Info("LUAS", luas)
	c.logger.Info("keliling", keliling)
	return luas, keliling
}

func (c *ClientStruct) HitungSegitiga(alas int32, tinggi int32, sisi1 int32, sisi2 int32, sisi3 int32) (int32, int32) {
	// c.logger.Info("RADIUS LINGKARAN", radius)
	luas := (alas * tinggi) / 2
	keliling := sisi1 + sisi2 + sisi3

	return luas, keliling

}
