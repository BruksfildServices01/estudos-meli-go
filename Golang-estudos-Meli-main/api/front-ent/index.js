import { TimeService, TorneioService, JogadorService } from './service.js';

const $ = (selector) => document.querySelector(selector);
const $$ = (selector) => document.querySelectorAll(selector);
const formTime = $('#form-time');
const btnCarregarTimes = $('#btn-carregar-times');
const listaTimes = $('#lista-times');

formTime.addEventListener('submit', async (e) => {
  e.preventDefault();
  const nome = $('#time-nome').value.trim();
  const cidade = $('#time-cidade').value.trim();
  if (!nome || !cidade) return;

  try {
    const time = await TimeService.create(nome, cidade);
    alert(`Time criado com ID ${time.id}`);
    formTime.reset();
    await carregarTimes();
  } catch (err) {
    alert('Erro ao criar time: ' + err.message);
  }
});

btnCarregarTimes.addEventListener('click', carregarTimes);

async function carregarTimes() {
  listaTimes.innerHTML = 'Carregando...';
  try {
    const times = await TimeService.list();
    listaTimes.innerHTML = '';
    times.forEach((t) => {
      const li = document.createElement('li');
      li.textContent = `ID: ${t.id} | ${t.nome} (${t.cidade})`;
      listaTimes.appendChild(li);
    });
  } catch (err) {
    listaTimes.innerHTML = 'Erro ao carregar times: ' + err.message;
  }
}

const formTorneio = $('#form-torneio');
const btnCarregarTorneios = $('#btn-carregar-torneios');
const listaTorneios = $('#lista-torneios');

formTorneio.addEventListener('submit', async (e) => {
  e.preventDefault();
  const nome = $('#torneio-nome').value.trim();
  const ano = $('#torneio-ano').value;
  if (!nome || !ano) return;

  try {
    const t = await TorneioService.create(nome, ano);
    alert(`Torneio criado com ID ${t.id}`);
    formTorneio.reset();
    await carregarTorneios();
  } catch (err) {
    alert('Erro ao criar torneio: ' + err.message);
  }
});

btnCarregarTorneios.addEventListener('click', carregarTorneios);

async function carregarTorneios() {
  listaTorneios.innerHTML = 'Carregando...';
  try {
    const torneios = await TorneioService.list();
    listaTorneios.innerHTML = '';
    torneios.forEach((t) => {
      const li = document.createElement('li');
      li.textContent = `ID: ${t.id} | ${t.nome} (${t.ano})`;
      listaTorneios.appendChild(li);
    });
  } catch (err) {
    listaTorneios.innerHTML = 'Erro ao carregar torneios: ' + err.message;
  }
}

const formJogador = $('#form-jogador');
const btnCarregarJogadores = $('#btn-carregar-jogadores');
const listaJogadores = $('#lista-jogadores');

formJogador.addEventListener('submit', async (e) => {
  e.preventDefault();
  const nome = $('#jogador-nome').value.trim();
  const idade = $('#jogador-idade').value;
  const timeId = $('#jogador-time-id').value;
  if (!nome || !idade || !timeId) return;

  try {
    const j = await JogadorService.create(nome, idade, timeId);
    alert(`Jogador criado com ID ${j.id}`);
    formJogador.reset();
    await carregarJogadores();
  } catch (err) {
    alert('Erro ao criar jogador: ' + err.message);
  }
});

btnCarregarJogadores.addEventListener('click', carregarJogadores);

async function carregarJogadores() {
  listaJogadores.innerHTML = 'Carregando...';
  try {
    const jogadores = await JogadorService.list();
    listaJogadores.innerHTML = '';
    jogadores.forEach((j) => {
      const li = document.createElement('li');
      li.textContent = `ID: ${j.id} | ${j.nome}, ${j.idade} anos (Time ID: ${j.time_id})`;
      listaJogadores.appendChild(li);
    });
  } catch (err) {
    listaJogadores.innerHTML = 'Erro ao carregar jogadores: ' + err.message;
  }
}

const formTorneioTime = $('#form-torneio-time');
const btnListarTimesTorneio = $('#btn-listar-times-torneio');
const listaTimesTorneio = $('#lista-times-torneio');

formTorneioTime.addEventListener('submit', async (e) => {
  e.preventDefault();
  const torneioId = $('#tt-torneio-id').value;
  const timeId = $('#tt-time-id').value;
  if (!torneioId || !timeId) return;

  try {
    await TorneioService.addTime(torneioId, timeId);
    alert(`Time ${timeId} adicionado ao torneio ${torneioId}`);
    formTorneioTime.reset();
  } catch (err) {
    alert('Erro ao adicionar time ao torneio: ' + err.message);
  }
});

btnListarTimesTorneio.addEventListener('click', async () => {
  const torneioId = $('#tt-list-torneio-id').value;
  if (!torneioId) {
    alert('Informe o ID do torneio');
    return;
  }

  listaTimesTorneio.innerHTML = 'Carregando...';
  try {
    const timeIds = await TorneioService.listTimes(torneioId);
    listaTimesTorneio.innerHTML = '';

    if (!timeIds.length) {
      listaTimesTorneio.textContent = 'Nenhum time nesse torneio.';
      return;
    }

    timeIds.forEach((id) => {
      const li = document.createElement('li');
      li.textContent = `Time ID: ${id}`;
      listaTimesTorneio.appendChild(li);
    });
  } catch (err) {
    listaTimesTorneio.innerHTML = 'Erro ao listar times: ' + err.message;
  }
});