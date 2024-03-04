package config

import "github.com/go-chi/chi/v5"

type MuxInterface interface {
	GetMux() *chi.Mux
}

type Mux struct {
	mux *chi.Mux
}

var M MuxInterface

func NewMux() {
	m := &Mux{
		mux: chi.NewMux(),
	}
	M = m
}

func (m *Mux) GetMux() *chi.Mux {
	return m.mux
}
