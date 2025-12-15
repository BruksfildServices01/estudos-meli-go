export class Time {
  constructor(id, nome, cidade) {
    this.id = id;
    this.nome = nome;
    this.cidade = cidade;
  }
}

export class Torneio {
  constructor(id, nome, ano) {
    this.id = id;
    this.nome = nome;
    this.ano = ano;
  }
}

export class Jogador {
  constructor(id, nome, idade, timeId) {
    this.id = id;
    this.nome = nome;
    this.idade = idade;
    this.time_id = timeId;
  }
}

// faz relacao de time com campeoanto
export class TorneioTime {
  constructor(torneioId, timeId) {
    this.torneio_id = torneioId;
    this.time_id = timeId;
  }
}