// Backend que controla a lógica de jogo
// do jogo Ed-Nal-Doh!
package main

import(
	"fmt"
	"log"
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

	r := gin.Default()

	r.GET("/api/jogo/:codigoSala", retornarDadosSala)
	r.POST("/api/jogo", criarSala)
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

func criarSala(c *gin.Context) {
	var reqBody requestCriarSala
	err := c.BindJSON(&reqBody)
	if err != nil {
		log.Fatal("Erro ao criar sala.")
	}

	j1 := jogador{
		Id: reqBody.IdJogador1,
		// Mao: []carta{},
		Baralho: baralho,
		Vida: 2000,
	}
	
	j2 := jogador{
		Id: reqBody.IdJogador1,
		// Mao: []carta{},
		Baralho: baralho,
		Vida: 2000,
	}

	// Registrar sala
	sala := sala{ Jogador1: j1, Jogador2: j2, CodigoSala: reqBody.CodigoSala }
	fmt.Println(j1)
	salas = append(salas, sala)

	c.IndentedJSON(http.StatusOK, gin.H{ "message": "É HORA DO DUELO!" });
}

func retornarTodasCartas(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cartas);
} 