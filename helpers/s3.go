package helpers

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"mime/multipart"
	"time"
)

// Fungsi untuk mengunggah gambar ke S3
func UploadImageToS3(photoID string, file multipart.File, fileSize int64, svc *s3.S3) error {
	// Buat nama objek unik di S3
	objectKey := fmt.Sprintf("photos/%s.jpg", photoID)

	// Buat input untuk operasi upload
	params := &s3.PutObjectInput{
		Bucket:      aws.String("mybucketgram"), // Ganti dengan nama bucket Anda
		Key:         aws.String(objectKey),
		Body:        file,
		ContentType: aws.String("image/jpeg"),
		ACL:         aws.String("private"), // Akses publik untuk URL gambar
	}

	// Upload objek ke S3
	_, err := svc.PutObject(params)
	if err != nil {
		return err
	}

	return nil
}

// Fungsi untuk mendapatkan URL gambar dari S3
func GetS3ImageURL(photoID string, svc *s3.S3) string {
	// Dapatkan URL gambar dari S3 berdasarkan bucket dan kunci objek
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String("mybucketgram"), // Ganti dengan nama bucket Anda
		Key:    aws.String(fmt.Sprintf("photos/%s.jpg", photoID)),
	})

	url, err := req.Presign(8760 * time.Hour)
	if err != nil {
		log.Fatal(err)
	}

	return url
}
