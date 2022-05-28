package model

import "mime/multipart"

type Request struct {
	file multipart.File `json:"file"`
}
