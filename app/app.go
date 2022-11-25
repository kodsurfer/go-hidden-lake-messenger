package app

import (
  "net/http"
)

type App struct {
  webSvc *http.Server
  incSvc *http.Server
}
