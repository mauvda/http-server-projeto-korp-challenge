# Desafio Técnico - http-server-projeto-korp

Este repositório contém a solução do desafio técnico para a **Korp**. O projeto consiste no desenvolvimento de um microsserviço escalável em Golang, containerização com Docker, proxy reverso com NGINX, observabilidade via Prometheus/Grafana e automação ponta a ponta (do zero) utilizando Ansible.

---

## 🛠️ Stack Tecnológica e Arquitetura

* **Linguagem Principal:** Go (Golang 1.23) utilizando o cliente oficial do Prometheus para instrumentação.
* **Proxy Reverso:** NGINX atuando na porta 80 do host.
* **Orquestração local:** Docker & Docker Compose (Rede isolada em modo `bridge`).
* **Observabilidade:** Prometheus (coleta de métricas) e Grafana (visualização).
* **Automação de Infraestrutura (IaC):** Ansible Playbook.

### Fluxo de Tráfego e Redes:
1. O usuário/teste faz a requisição na porta `80` do Host.
2. O container **NGINX** recebe o tráfego e faz o proxy reverso para o container **Go** na porta `8080` utilizando o nome do serviço na rede interna do Docker.
3. O container da aplicação em Go responde dinamicamente com o JSON e o horário em UTC.
4. O **Prometheus** realiza o *scrape* dos dados diretamente da aplicação via rede `bridge`.

---

## 🚀 Como Executar (Provisionamento 100% Automatizado)

Seguindo os critérios de aceitação do desafio, todo o ambiente — desde a instalação do Docker no sistema operacional até o deploy e validação dos containers — é executado com **um único comando**.

### Pré-requisitos
* Um ambiente Linux baseado em Debian/Ubuntu (pode ser executado localmente via WSL).
* **Ansible** instalado na máquina de controle.

### Execução:

1. Clone o repositório para o diretório de sua preferência:
```bash
   git clone [https://github.com/mauvda/http-server-projeto-korp-challenge.git](https://github.com/mauvda/http-server-projeto-korp-challenge.git)
```   
2. Navegue até o diretório:
```bash
   cd http-server-projeto-korp-challenge
```
3. Execute o comando ansible para rodar todo o projeto:
```bash
   ansible-playbook -i ansible/inventory.ini ansible/playbook.yml --ask-become-pass
```