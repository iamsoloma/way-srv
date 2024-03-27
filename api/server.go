package api

import (
	"encoding/json"
	"net/http"
	"strconv"
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

	mux.HandleFunc("POST /create", s.handleCreateBlockChain)
	mux.HandleFunc("DELETE /delete", s.handleDeleteBlockChain)
	mux.HandleFunc("GET /getlastblock/{chainName}", s.handleGetLastBlock)
	mux.HandleFunc("GET /getblockbyid/{chainName}/{id}", s.handleGetBlockByID)
	mux.HandleFunc("PUT /addblock", s.handleAddBlock)

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

func (s *ApiServer) handleDeleteBlockChain(w http.ResponseWriter, r *http.Request) {
	req := types.DeleteBlockChainRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	Exp := way.Explorer{
		Path: s.StoragePath,
		Name: req.ChainName,
	}

	if found, err := Exp.DeleteBlockChain(); err != nil && found {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if !found && err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	
}

func (s *ApiServer) handleGetLastBlock(w http.ResponseWriter, r *http.Request) {
	/*req := types.GetLastBlockRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}*/

	Exp := way.Explorer{
		Path: s.StoragePath,
		Name: r.PathValue("chainName"),//req.ChainName
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
	/*req := types.GetBlockByIDRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}*/

	Exp := way.Explorer{
		Path: s.StoragePath,
		Name: r.PathValue("chainName"),//req.ChainName
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err!= nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	
	block, err := Exp.GetBlockByID(id/*req.ID*/)
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
	req := types.AddBlockRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	Exp := way.Explorer{
		Path: s.StoragePath,
		Name: req.ChainName,
	}


	var err error
	var id int
	if time.Time.IsZero(req.Time_UTC) {
		id, err = Exp.AddBlock(req.Data, time.Now().UTC())
	} else {
		id, err = Exp.AddBlock(req.Data, req.Time_UTC)
	}
	
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
