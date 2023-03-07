// Backend que controla a lógica de jogo
// do jogo Ed-Nal-Doh!
package main

import(
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type sala struct {
	Jogador1   jogador `json:"jogador"`
	Jogador2   jogador `json:"jogador"`
	CodigoSala string  `json:"codigoSala"`
	Turno 	   string  `json:"turno"`
}

type jogador struct {
	Id 		string   `json:"id"`
	Mao     []carta  `json:"mao"`
	Baralho []carta  `json:"baralho"`
	Campo   [6]carta `json:"campo"`
	Vida    int 	 `json:"vida"`
	//Tempo   int        `json:"tempo"`
}

var salas []sala

func main() {

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"PUT", "GET", "POST", "DELETE"},
        AllowHeaders:     []string{"*"},
        AllowCredentials: true,
	}))

	/* Rotas */
	r.GET("/api/jogo/:codigoSala", retornarDadosSala)
	r.POST("/api/jogo", criarSala)
	// r.DELETE("/api/jogo", removerSala)
	r.GET("/api/cartas", retornarTodasCartas)
	r.GET("/api/jogo/usuario/:id", retornarEstadoDeJogo)

	r.Run("127.0.0.1:4000")

	fmt.Println("API rodando!")
	
}


func retornarDadosSala(c *gin.Context) {
	codigoSala := c.Param("codigoSala")

	for i, s := range salas { 
		if codigoSala == s.CodigoSala {
			c.IndentedJSON(http.StatusOK, gin.H{ "jogador1": salas[i].Jogador1, "jogador2": salas[i].Jogador2 })
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
		Baralho: baralho,
		Vida: 2000,
	}
	
	j2 := jogador{
		Id: reqBody.IdJogador2,
		Baralho: baralho,
		Vida: 2000,
	}

	for i := 0; i < 5; i++ {
		var cartaJ1 carta
		var cartaJ2 carta

		j1.Baralho, cartaJ1 = tirarCartaDoBaralho(j1.Baralho)
		j2.Baralho, cartaJ2 = tirarCartaDoBaralho(j2.Baralho)

		j1.Mao = append(j1.Mao, cartaJ1)
		j2.Mao = append(j2.Mao, cartaJ2)
	}

	// Registrar sala
	sala := sala{ Jogador1: j1, Jogador2: j2, CodigoSala: reqBody.CodigoSala, Turno: j1.Id }
	salas = append(salas, sala)

	c.IndentedJSON(http.StatusOK, gin.H{ "message": "É HORA DO DUELO!", "codigoSala": reqBody.CodigoSala })
}

func retornarEstadoDeJogo(c *gin.Context) {
	userId := c.Param("id")

	for _, s := range salas {
		if (userId == s.Jogador1.Id) {
			j1 := jogador{ Mao: s.Jogador1.Mao, Campo: s.Jogador1.Campo, Vida: s.Jogador1.Vida }
			j2 := jogador{ Campo: s.Jogador2.Campo, Vida: s.Jogador2.Vida }

			// VULNERABILIDADE -> CLIENT RECEBE CARTAS DO ADVERSÁRIO EM CAMPO!
			c.IndentedJSON(http.StatusOK, gin.H{ "voce": j1, "adversario": j2, "maoAdversario": len(s.Jogador2.Mao) })
			return
		} else if (userId == s.Jogador2.Id) {
			j1 := jogador{ Mao: s.Jogador2.Mao, Campo: s.Jogador2.Campo, Vida: s.Jogador2.Vida }
			j2 := jogador{ Campo: s.Jogador1.Campo, Vida: s.Jogador1.Vida }

			// VULNERABILIDADE -> CLIENT RECEBE CARTAS DO ADVERSÁRIO EM CAMPO!
			c.IndentedJSON(http.StatusOK, gin.H{ "voce": j1, "adversario": j2, "maoAdversario": len(s.Jogador1.Mao) })
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{ "error": "Usuário não encontrado em nenhuma sala :(" })
}

func retornarTodasCartas(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cartas);
}

func tirarCartaDoBaralho(baralho []carta) ([]carta, carta) {
	randNum := rand.Intn(len(baralho))
	tempCard := baralho[randNum]

	baralho[randNum] = baralho[len(baralho) - 1]
	baralho = baralho[:len(baralho) - 1]

	return baralho, tempCard
}