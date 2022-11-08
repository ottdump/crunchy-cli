package utils

import (
	"github.com/ottdump/crunchyroll-go/v3"
	"net/http"
)

var Version = "master"

var (
	Crunchy *crunchyroll.Crunchyroll
	Client  *http.Client
	Log     Logger
)
