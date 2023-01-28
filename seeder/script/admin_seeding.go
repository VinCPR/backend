package script

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/rs/zerolog/log"

	"github.com/VinCPR/backend/api"
)

func SeedAdminData(basePath string, server *api.Server) {
	type userData struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8,max=64"`
		RoleName string `json:"role_name" binding:"required,alpha"`
	}
	var admins = []userData{
		{
			Email:    "tam.nt@vinuni.com",
			Password: "nguyenthitam",
			RoleName: "admin",
		},
		{
			Email:    "thuy.hm@vinuni.com",
			Password: "haminhthuy",
			RoleName: "admin",
		},
	}
	for _, admin := range admins {
		url := basePath + "/users"
		recorder := httptest.NewRecorder()
		// Marshal body data to JSON
		data, err := json.Marshal(admin)
		if err != nil {
			log.Fatal().Err(err).Msg("cannot parse create_user request body")
		}
		request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
		if err != nil {
			log.Fatal().Err(err).Msg("cannot init create_user request")
		}
		server.ServeHTTP(recorder, request)
		if recorder.Code != 200 {
			log.Fatal().Msg("failed to call create_user request")
		}
	}
}
