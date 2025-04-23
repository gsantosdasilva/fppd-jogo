// personagem.go - Funções para movimentação e ações do personagem
package main

import (
	"math/rand"
)

// Atualiza a posição do personagem com base na tecla pressionada (WASD)
func personagemMover(tecla rune, jogo *Jogo) {
	dx, dy := 0, 0
	switch tecla {
	case 'w', 'W':
		dy = -1
	case 's', 'S':
		dy = 1
	case 'a', 'A':
		dx = -1
	case 'd', 'D':
		dx = 1
	}

	// Verifica se pode mover
	if jogoPodeMoverPara(jogo, jogo.PosX+dx, jogo.PosY+dy) {
		// Verifica se há elementos especiais na nova posição
		novaPos := jogo.Mapa[jogo.PosY+dy][jogo.PosX+dx]
		if novaPos == PortalMagico {
			// Teletransporta para posição aleatória
			jogo.PosX = rand.Intn(len(jogo.Mapa[0]))
			jogo.PosY = rand.Intn(len(jogo.Mapa))
			jogo.StatusMsg = "Você foi teletransportado!"
		} else if novaPos == ArmadilhaElem {
			jogo.StatusMsg = "Você ativou uma armadilha!"
			// Envia sinal para ativar a armadilha
			canalArmadilha <- struct{}{}
		} else {
			// Move normalmente
			jogoMoverElemento(jogo, jogo.PosX, jogo.PosY, dx, dy)
			jogo.PosX += dx
			jogo.PosY += dy
		}
	}

	// Verifica se o patroller está próximo
	if abs(jogo.PosX-5) <= 3 && abs(jogo.PosY-5) <= 3 {
		canalAvistamento <- struct{}{}
	}
}

// Define o que ocorre quando o jogador pressiona a tecla de interação
// Neste exemplo, apenas exibe uma mensagem de status
// Você pode expandir essa função para incluir lógica de interação com objetos
func personagemInteragir(jogo *Jogo) {
	// Ativa o portal mágico
	canalPortal <- struct{}{}
	jogo.StatusMsg = "Portal mágico ativado!"
}

// Processa o evento do teclado e executa a ação correspondente
func personagemExecutarAcao(evento EventoTeclado, jogo *Jogo) bool {
	switch evento.Tipo {
	case "sair":
		return false
	case "interagir":
		personagemInteragir(jogo)
	case "mover":
		personagemMover(evento.Tecla, jogo)
	}
	return true
}

// Função auxiliar para valor absoluto
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
