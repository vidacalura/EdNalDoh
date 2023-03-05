// Este arquivo lida com a lógica do jogo e cartas
// (eventos enviados e recebidos do servidor)
let socket = io();

const codigoSala = window.location.href.split("/")[window.location.href.split("/").length - 1];
socket.emit("entrarSala", codigoSala);

const codigoSalaTextbox = document.getElementById("input-convite");
codigoSalaTextbox.value = window.location.href;


socket.on("iniciarPartida", (dadosPartida) => {
    if (codigoSala == dadosPartida.codigoSala) {
        const telaConvite = document.getElementById("convite-container");
        telaConvite.style.display = "none";
    }
});

socket.on("desconexaoDeAdversario", (codigoSalaDesconexao) => {
    if (codigoSala == codigoSalaDesconexao) {
        alert("Oh não! Seu oponente saiu da partida :(");
    }
});