# Impacto CLI

**CLI local-first para geração de manifestos Crossplane Managed Resources.**

O Impacto CLI ajuda desenvolvedores a provisionar infraestrutura em nuvem através do Crossplane, gerando templates YAML prontos para produção com comentários claros e links para documentação — sem necessidade de expertise profunda em cloud.

No momento, estamos focados em recursos da Oracle Cloud Infrastructure (OCI), mas a arquitetura é extensível para suportar outros provedores no futuro.

---

## 🚀 Funcionalidades

- **Geração de manifestos**: Crie Managed Resources do Crossplane com um único comando.
- **Templates prontos para produção**: YAML com placeholders `# CHANGE:` que guiam a personalização.
- **Multi-recursos OCI**: Suporte a Compartment, VCN, Subnet e Bucket (Oracle Cloud Infrastructure).
- **Listagem interativa**: Consulte os provedores e recursos disponíveis.
- **Extensível**: Adicione novos provedores e recursos facilmente via templates embutidos.

---

## 📦 Instalação

### Pré-requisitos
- Go 1.21+ (apenas para compilar; o binário final é independente)
- Cluster Kubernetes com Crossplane instalado e Providers OCI configurados

### Onde encontrar provedores e recursos suportados

A Oracle disponibiliza os pacotes do Crossplane Provider OCI no GitHub Packages. Você pode verificar os recursos disponíveis e suas versões acessando:
[https://github.com/orgs/oracle/packages?repo_name=crossplane-provider-oci](https://github.com/orgs/oracle/packages?repo_name=crossplane-provider-oci)

### Instalação via `go install`
```bash
go install github.com/ShlomoChanoch/impacto-cli@latest
```

### Compilação manual
```bash
git clone https://github.com/ShlomoChanoch/impacto-cli.git
cd impacto-cli
go build -ldflags "-X github.com/ShlomoChanoch/impacto-cli/cmd.version=$(git describe --tags --always)" -o impacto .
sudo install -m 755 impacto /usr/local/bin/
```

---

## 📖 Uso

### Listar recursos disponíveis
```bash
impacto list
```
**Saída:**
```
PROVIDER   RESOURCE      DESCRIPTION
--------   --------      -----------
oci        bucket        Object Storage Bucket
oci        compartment   OCI Compartment (pré-requisito para todos os recursos)
oci        subnet        Subnet dentro de uma VCN
oci        vcn           Virtual Cloud Network (pré-requisito de rede)
```

### Gerar um Managed Resource
```bash
impacto mr <provedor> <recurso>
```

**Exemplos:**
```bash
# Compartimento OCI
impacto mr oci compartment

# Bucket no Object Storage
impacto mr oci bucket

# Rede Virtual (VCN)
impacto mr oci vcn

# Subnet
impacto mr oci subnet

# Salvar diretamente em arquivo
impacto mr oci bucket > meu-bucket.yaml

# Edite o arquivo meu-bucket.yaml para preencher os campos necessários
```

### Aplicar o manifesto gerado depois da edição
```
kubectl apply -f meu-bucket.yaml
```


### Verificar versão e provedores suportados
```bash
impacto version
```


---

## 🏗️ Estrutura do Projeto

```
.
├── cmd/                    # Comandos da CLI (Cobra)
│   ├── root.go             # Comando raiz
│   ├── list.go             # Listagem de provedores/recursos
│   ├── mr.go               # Geração de Managed Resources
│   └── version.go          # Informações de versão
├── internal/
│   └── generator/          # Lógica de validação e renderização
├── templates/              # Templates YAML embutidos (embed.FS)
│   └── oci/                # Recursos Oracle Cloud Infrastructure
│       ├── bucket.yaml
│       ├── compartment.yaml
│       ├── subnet.yaml
│       └── vcn.yaml
├── main.go                 # Ponto de entrada
└── go.mod
```

---

## 🔧 Personalização

Cada template gerado contém comentários `# CHANGE:` indicando onde você deve ajustar os valores. Exemplo:

```yaml
# CHANGE: your compartment OCID
compartmentId: ocid1.compartment.oc1..aaaaaaaaxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

Basta substituir os placeholders pelos valores do seu ambiente OCI e aplicar com `kubectl`.

---

## 📋 Roadmap (MVP estendido)

- [x] Comando `list` — listar provedores e recursos
- [x] Comando `mr` — gerar templates OCI (Compartment, VCN, Subnet, Bucket)
- [x] Comando `version` — exibir versão e provedores suportados
- [x] Comentários `# CHANGE:` em todos os templates
- [x] Testes unitários para comandos, generator e templates
- [ ] Suporte a mais recursos OCI (Instance, Cluster, Database)
- [ ] Suporte a outros provedores (AWS, GCP, Azure)
- [ ] Validação interativa de parâmetros obrigatórios
- [ ] Geração de pacotes (vários recursos relacionados de uma vez)
- [ ] Documentação automática de fields via OpenAPI

---

## 🧪 Testes

Execute todos os testes unitários:

```bash
go test ./... -v
```

Cobertura atual inclui:
- Validação de argumentos dos comandos
- Geração e integridade dos templates YAML
- Idempotência das chamadas
- Mensagens de erro amigáveis
- Embed dos arquivos de template

---

## 🚧 Status do Projeto

O Impacto CLI é atualmente um experimento/MVP focado em uso local. No momento, **não estamos aceitando Pull Requests ou contribuições externas**, pois a arquitetura do projeto passará por grandes mudanças estruturais em breve. Fique à vontade para abrir Issues com feedbacks e sugestões!

---

**Impacto CLI** — Provisione nuvem sem ser especialista em cloud. ☁️✨
