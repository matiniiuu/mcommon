package cloudinary

import (
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
)

func connectCloudinary(cloudinaryURL string) *cloudinary.Cloudinary {
	cld, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		log.Fatalln(err)
	}
	return cld
}
