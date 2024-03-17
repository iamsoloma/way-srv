package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/TinajXD/way-srv/types"

	"github.com/TinajXD/way"
)

type ApiServer struct {
	Addr string
	TimeOut time.Duration
	IdleTimeOut time.Duration
}

func (s *ApiServer) Start() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/create", s.handleCreateBlockChain)

	err := http.ListenAndServe(s.Addr, mux)
	if err != nil {
		return err
	}
	return nil
}

func (s *ApiServer) handleCreateBlockChain(w http.ResponseWriter, r *http.Request) {
	req := types.CreateBlockChainRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	Exp := way.Explorer{
		Path: "blockchains" + req.Chain,
	}

	if err := Exp.CreateBlockChain(req.Genesis, req.Time_UTC); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}
