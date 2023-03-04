package main

type requestCriarSala struct {
	IdJogador1 string `json:"idJogador1"`
	IdJogador2 string `json:"idJogador2"`
	CodigoSala string `json:"codigoSala"`
}