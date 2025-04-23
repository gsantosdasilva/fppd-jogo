# Jogo com Elementos Concorrentes em Go

Este projeto implementa um jogo de terminal em Go com elementos concorrentes que interagem entre si através de goroutines e canais de comunicação.

## Compilação e Execução

### Pré-requisitos
- Go 1.16 ou superior
- Biblioteca termbox-go (`github.com/nsf/termbox-go`)

### Instalação
1. Clone o repositório
2. Instale a dependência necessária:
```bash
go get github.com/nsf/termbox-go
```

### Compilação
Para compilar o jogo, execute no terminal:
```bash
go build -o jogo.exe
```

### Execução
Para executar o jogo:
```bash
./jogo.exe
```
ou especifique um mapa personalizado:
```bash
./jogo.exe mapa.txt
```

## Controles
- WASD: Movimentação do personagem
- E: Ativar portal mágico
- ESC: Sair do jogo

## Elementos Implementados

### 1. Inimigo Patrulheiro (👾)
- **Comportamento**: Patrulha o mapa de forma autônoma e persegue o jogador quando detectado
- **Concorrência**: Executa em uma goroutine independente
- **Comunicação**: Recebe sinais através do `canalAvistamento` quando o jogador está próximo
- **Estados**:
  - Modo Patrulha: Move-se aleatoriamente pelo mapa
  - Modo Perseguição: Persegue o jogador por 2 segundos quando avistado
- **Timeout**: Retorna ao modo patrulha após 2 segundos de perseguição
- **Atualização**: A cada 500ms

### 2. Portal Mágico (⭕)
- **Comportamento**: Aparece em posições aleatórias quando ativado
- **Concorrência**: Executa em uma goroutine independente
- **Comunicação**: Ativado através do `canalPortal` quando o jogador pressiona 'E'
- **Funcionalidade**: Teletransporta o jogador para uma posição aleatória do mapa
- **Timeout**: Desaparece após 5 segundos de ativação
- **Atualização**: Verifica ativação a cada 1 segundo

### 3. Armadilha Explosiva (💣)
- **Comportamento**: Aparece em posições aleatórias e é ativada quando o jogador pisa nela
- **Concorrência**: Executa em uma goroutine independente
- **Comunicação**: Ativada através do `canalArmadilha` quando o jogador colide
- **Estados**: 
  - Inativa: Invisível no mapa
  - Ativa: Visível e causa dano ao jogador
- **Timeout**: Desaparece após 3 segundos de ativação
- **Atualização**: Verifica ativação a cada 1 segundo

## Aspectos Técnicos

### Concorrência
- Cada elemento é executado em sua própria goroutine
- Os elementos funcionam de forma independente da thread principal
- Utilização de `select` para escuta concorrente de múltiplos canais

### Comunicação
- Canais sem buffer para comunicação síncrona
- Uso de `struct{}` vazio para sinais entre elementos
- Timeouts implementados com `time.Sleep`

### Sincronização
- Acesso ao mapa protegido pela função `jogoPodeMoverPara`
- Comunicação entre elementos através de canais
- Atualização da interface sincronizada com o loop principal

## Estrutura do Código

- `main.go`: Loop principal e inicialização do jogo
- `elementos_concorrentes.go`: Implementação dos elementos concorrentes
- `interface.go`: Interface gráfica usando termbox
- `personagem.go`: Lógica de movimentação e interação do jogador

## Melhorias Futuras

1. Implementar sistema de pontuação
2. Adicionar mais tipos de inimigos
3. Criar power-ups e itens coletáveis
4. Implementar níveis de dificuldade
5. Adicionar efeitos sonoros


