// Este arquivo lida com a lógica do jogo e cartas
// (eventos enviados e recebidos do servidor)
let socket = io();

const codigoSala = window.location.href.split("/")[window.location.href.split("/").length - 1];
socket.emit("entrarSala", codigoSala);