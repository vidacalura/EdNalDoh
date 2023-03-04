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
	constructorCarta("Ednaldo Pereira Futebolístico", "/assets/", 3, 1500, 800, "0001", "none", false),
	constructorCarta("Ednaldo Esticadin", "/assets/", 4, 1200, 1200, "0002", "none", false),
	constructorCarta("Ednaldo Minecraft", "/assets/", 0, 0, 0, "0003", "none", false),
	constructorCarta("EdnaldoZap", "/assets/", 5, 1800, 300, "0004", "none", false),
	constructorCarta("Mini Ednaldo Assassino", "/assets/", 3, 800, 100, "0005", "none", false),
	constructorCarta("Casas Paraíba", "/assets/", 0, 0, 0, "0006", "none", false),
	constructorCarta("eDIOnaldo", "/assets/", 4, 1600, 700, "0007", "none", false),
	constructorCarta("EdnaandE", "/assets/", 0, 0, 0, "0008", "none", false),
	constructorCarta("Ednaldo Flor", "/assets/", 3, 0, 1200, "0009", "none", false),
	constructorCarta("Ednaldo Mengão", "/assets/", 2, 900, 500, "0010", "none", false),
	constructorCarta("Ednaldo Comunista", "/assets/", 0, 0, 0, "0011", "none", false),
	constructorCarta("Ednaldo Dorime", "/assets/", 3, 500, 900, "0012", "none", false),
	constructorCarta("Ednaldo Roblos", "/assets/", 4, 1300, 400, "0013", "none", false),
	constructorCarta("Ednaldo Enjaulado", "/assets/", 6, 0, 1000, "0014", "none", false),
	constructorCarta("Ednaldo Pereira", "/assets/", 10, 2000, 2000, "0015", "none", false),
	constructorCarta("Ednaldo Pistoleiro", "/assets/", 5, 1600, 600, "0016", "none", false),
	constructorCarta("Ednaldo Anjo", "/assets/", 5, 600, 1600, "0017", "none", false),
	constructorCarta("Ednaldo Banido", "/assets/", 0, 0, 0, "0018", "none", false),
	constructorCarta("Ednaldo Tecladista", "/assets/", 2, 0, 800, "0019", "none", false),
	constructorCarta("Caneta Azul (Blue Pen)", "/assets/", 4, 1200, 1200, "0020", "none", false),
	constructorCarta("Bolsonaro Ednaldista", "/assets/", 4, 1400, 900, "0021", "none", false),
	constructorCarta("Caneta Bic azul", "/assets/", 0, 0, 0, "0022", "none", false),
	constructorCarta("Ednaldo Festeiro", "/assets/", 4, 200, 800, "0023", "none", false),
	constructorCarta("Filhote de Ednaldo", "/assets/", 1, 300, 300, "0024", "none", false),
	constructorCarta("The Brother", "/assets/", 2, 600, 600, "0025", "none", false),
	constructorCarta("Ednaldo Drip", "/assets/", 3, 1400, 200, "0026", "none", false),
	constructorCarta("Ednaldo Coreano", "/assets/", 3, 900, 500, "0027", "none", false),
	constructorCarta("Substituição", "/assets/", 0, 0, 0, "0028", "none", false),
	constructorCarta("Ednaldificação", "/assets/", 0, 0, 0, "0029", "none", false),
	constructorCarta("Nego Ney", "/assets/", 8, 1800, 300, "0030", "none", false),
}