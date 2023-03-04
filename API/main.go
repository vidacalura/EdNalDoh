package main

import(
	//"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type sala struct {
	Jogador1   jogador `json:"jogador"`
	Jogador2   jogador `json:"jogador"`
	CodigoSala string  `json:"codigoSala"`
}

type jogador struct {
	Id 		string     `json:"id"`
	Mao     []carta    `json:"mao"`
	Baralho [60]carta  `json:"baralho"`
	Campo   [6]carta   `json:"campo"`
	Vida    int 	   `json:"vida"`
	//Tempo   int        `json:"tempo"`
}

var salas []sala

func main() {
	j1 := jogador{ Id: "sla", Vida: 2000 }
	j2 := jogador{}
	sala := sala{ Jogador1: j1, Jogador2: j2, CodigoSala: "1234" }
	salas = append(salas, sala)

	r := gin.Default()

	r.GET("/api/jogo/:codigoSala", retornarDadosSala)
	r.GET("/api/cartas", retornarTodasCartas)

	r.Run("localhost:4000")

}

func retornarDadosSala(c *gin.Context) {
	codigoSala := c.Param("codigoSala")

	for i, s := range salas { 
		if codigoSala == s.CodigoSala {
			c.IndentedJSON(http.StatusOK, salas[i])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{ "message": "Erro, sala não encontrada" })
}

func retornarTodasCartas(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cartas);
} 