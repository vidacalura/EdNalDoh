// Contém todos os dados das cartas do
// jogo Ed-Nal-Doh!
package main

type carta struct {
	Nome          string `json:"nome"`
	ImgPath 	  string `json:"imgPath"`
	Estrelas      int    `json:"estrelas"`
	Ataque        int    `json:"ataque"`
	Defesa        int    `json:"defesa"`
	NumeroDeSerie string `json:"numeroDeSerie"`
	Estado		  string `json:"estado"` // defesa / ataque / none
	Virada		  bool   `json:"virada"`
}

func constructorCarta(nome string, imgPath string, estrelas int, ataque int, defesa int, 
	numeroDeSerie string, estado string, virada bool) carta {
	return carta{ Nome: nome, ImgPath: imgPath, Estrelas: estrelas, Ataque: ataque,
		Defesa: defesa, NumeroDeSerie: numeroDeSerie, Estado: estado, Virada: virada }
}

var cartas = [30]carta{
	constructorCarta("Ednaldo Pereira Futebolístico", "/ednaldo0001.png", 3, 1500, 800, "0001", "none", false),
	constructorCarta("Ednaldo Esticadin", "/ednaldo0002.png", 4, 1200, 1200, "0002", "none", false),
	constructorCarta("Ednaldo Minecraft", "/ednaldo0003.png", 0, 0, 0, "0003", "none", false),
	constructorCarta("EdnaldoZap", "/ednaldo0004.png", 5, 1800, 300, "0004", "none", false),
	constructorCarta("Mini Ednaldo Assassino", "/ednaldo0005.png", 3, 800, 100, "0005", "none", false),
	constructorCarta("Casas Paraíba", "/ednaldo0006.png", 0, 0, 0, "0006", "none", false),
	constructorCarta("eDIOnaldo", "/ednaldo0007.png", 4, 1600, 700, "0007", "none", false),
	constructorCarta("EdnaandE", "/ednaldo0008.png", 0, 0, 0, "0008", "none", false),
	constructorCarta("Ednaldo Flor", "/ednaldo0009.png", 3, 0, 1200, "0009", "none", false),
	constructorCarta("Ednaldo Mengão", "/ednaldo0010.png", 2, 900, 500, "0010", "none", false),
	constructorCarta("Ednaldo Comunista", "/ednaldo0011.png", 0, 0, 0, "0011", "none", false),
	constructorCarta("Ednaldo Dorime", "/ednaldo0012.png", 3, 500, 900, "0012", "none", false),
	constructorCarta("Ednaldo Roblos", "/ednaldo0013.png", 4, 1300, 400, "0013", "none", false),
	constructorCarta("Ednaldo Enjaulado", "/ednaldo0014.png", 6, 0, 1000, "0014", "none", false),
	constructorCarta("Ednaldo Pereira", "/ednaldo0015.png", 10, 2000, 2000, "0015", "none", false),
	constructorCarta("Ednaldo Pistoleiro", "/ednaldo0016.png", 5, 1600, 600, "0016", "none", false),
	constructorCarta("Ednaldo Anjo", "/ednaldo0017.png", 5, 600, 1600, "0017", "none", false),
	constructorCarta("Ednaldo Banido", "/ednaldo0018.png", 0, 0, 0, "0018", "none", false),
	constructorCarta("Ednaldo Tecladista", "/ednaldo0019.png", 2, 0, 800, "0019", "none", false),
	constructorCarta("Caneta Azul (Blue Pen)", "/ednaldo0020.png", 4, 1200, 1200, "0020", "none", false),
	constructorCarta("Bolsonaro Ednaldista", "/ednaldo0021.png", 4, 1400, 900, "0021", "none", false),
	constructorCarta("Caneta Bic azul", "/ednaldo0022.png", 0, 0, 0, "0022", "none", false),
	constructorCarta("Ednaldo Festeiro", "/ednaldo0023.png", 4, 200, 800, "0023", "none", false),
	constructorCarta("Filhote de Ednaldo", "/ednaldo0024.png", 1, 300, 300, "0024", "none", false),
	constructorCarta("The Brother", "/ednaldo0025.png", 2, 600, 600, "0025", "none", false),
	constructorCarta("Ednaldo Drip", "/ednaldo0026.png", 3, 1400, 200, "0026", "none", false),
	constructorCarta("Ednaldo Coreano", "/ednaldo0027.png", 3, 900, 500, "0027", "none", false),
	constructorCarta("Substituição", "/ednaldo0028.png", 0, 0, 0, "0028", "none", false),
	constructorCarta("Ednaldificação", "/ednaldo0029.png", 0, 0, 0, "0029", "none", false),
	constructorCarta("Nego Ney", "/ednaldo0030.png", 8, 1800, 300, "0030", "none", false),
}

var baralho = []carta{
	cartas[0],
	cartas[1], cartas[1], 
	cartas[2], cartas[2], 
	cartas[3], 
	cartas[4], cartas[4], cartas[4], 
	cartas[5], cartas[5], cartas[5], 
	cartas[6],
	cartas[7], cartas[7], cartas[7], 
	cartas[8], cartas[8], 
	cartas[9], cartas[9], cartas[9], 
	cartas[10], cartas[10], 
	cartas[11], cartas[11], cartas[11], 
	cartas[12],
	cartas[13], 
	cartas[15],
	cartas[16], cartas[16],
	cartas[17], cartas[17],
	cartas[18], cartas[18], cartas[18], 
	cartas[19], cartas[19], 
	cartas[20], cartas[20], 
	cartas[21], 
	cartas[22], cartas[22], cartas[22],  
	cartas[23], cartas[23], cartas[23],
	cartas[24], cartas[24], cartas[24], 
	cartas[25], cartas[25],
	cartas[26], cartas[26],
	cartas[27], cartas[27], cartas[27], 
	cartas[28], cartas[28], 
	cartas[29], 
}