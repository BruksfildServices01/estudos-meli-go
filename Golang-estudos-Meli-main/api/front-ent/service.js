import { Time, Torneio, Jogador } from './model.js';

const BASE_URL = 'http://localhost:8081';

async function request(url, options = {}) {
  const res = await fetch(url, options);
  if (!res.ok) {
    const text = await res.text();
    throw new Error(text || `Erro HTTP ${res.status}`);
  }
  
  const contentType = res.headers.get('Content-Type') || '';
  if (contentType.includes('application/json')) {
    return res.json();
  }
  return null;
}

export const TimeService = {
  async create(nome, cidade) {
    const data = await request(`${BASE_URL}/times`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ nome, cidade }),
    });
    return new Time(data.id, data.nome, data.cidade);
  },

  async list() {
    const data = await request(`${BASE_URL}/times`);
    return data.map(t => new Time(t.id, t.nome, t.cidade));
  },
};

export const TorneioService = {
  async create(nome, ano) {
    const data = await request(`${BASE_URL}/torneios`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ nome, ano: Number(ano) }),
    });
    return new Torneio(data.id, data.nome, data.ano);
  },

  async list() {
    const data = await request(`${BASE_URL}/torneios`);
    return data.map(t => new Torneio(t.id, t.nome, t.ano));
  },

  async addTime(torneioId, timeId) {
    await request(`${BASE_URL}/torneios/${torneioId}/times`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ time_id: Number(timeId) }),
    });
  },

  async listTimes(torneioId) {
    const ids = await request(`${BASE_URL}/torneios/${torneioId}/times`);
    return ids; 
  },
};

export const JogadorService = {
  async create(nome, idade, timeId) {
    const data = await request(`${BASE_URL}/jogadores`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        nome,
        idade: Number(idade),
        time_id: Number(timeId),
      }),
    });
    return new Jogador(data.id, data.nome, data.idade, data.time_id);
  },

  async list() {
    const data = await request(`${BASE_URL}/jogadores`);
    return data.map(j => new Jogador(j.id, j.nome, j.idade, j.time_id));
  },
};