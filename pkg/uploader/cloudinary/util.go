package cloudinary

import (
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
)

func connectCloudinary() *cloudinary.Cloudinary {
	cld, err := cloudinary.NewFromURL("cloudinary://952375858838726:vjPMa9qbGgN7CVX4p3ZFcfMDakI@goldilan")
	if err != nil {
		log.Fatalln(err)
	}
	return cld
}
