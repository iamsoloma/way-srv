package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/TinajXD/way-srv/types"

	"github.com/TinajXD/way"
)

type ApiServer struct {
	Addr        string
	StoragePath string
	TimeOut     time.Duration
	IdleTimeOut time.Duration
}

func (s *ApiServer) Start() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/create", s.handleCreateBlockChain)
	mux.HandleFunc("/getlastblock", s.handleGetLastBlock)
	mux.HandleFunc("/getblockbyid", s.handleGetBlockByID)
	mux.HandleFunc("/addblock", s.handleAddBlock)

	srv := http.Server{
		Addr:         s.Addr,
		Handler:      mux,
		ReadTimeout:  s.TimeOut,
		WriteTimeout: s.TimeOut,
		IdleTimeout:  s.IdleTimeOut,
	}

	err := srv.ListenAndServe()
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
		Path: s.StoragePath,
		Name: req.ChainName,
	}

	if err := Exp.CreateBlockChain(req.Genesis, time.Now().UTC()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *ApiServer) handleGetLastBlock(w http.ResponseWriter, r *http.Request) {
	req := types.GetLastBlock{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	Exp := way.Explorer{
		Path: s.StoragePath,
		Name: req.ChainName,
	}

	block, err := Exp.GetLastBlock()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	jsonResp, err := json.Marshal(block)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func (s *ApiServer) handleGetBlockByID(w http.ResponseWriter, r *http.Request) {
	req := types.GetBlockByID{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	Exp := way.Explorer{
		Path: s.StoragePath,
		Name: req.ChainName,
	}

	block, err := Exp.GetBlockByID(req.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	jsonResp, err := json.Marshal(block)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func (s *ApiServer) handleAddBlock(w http.ResponseWriter, r *http.Request) {
	req := types.AddBlock{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	Exp := way.Explorer{
		Path: s.StoragePath,
		Name: req.ChainName,
	}

	id, err := Exp.AddBlock(req.Data, time.Now().UTC())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	jsonResp, err := json.Marshal(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
