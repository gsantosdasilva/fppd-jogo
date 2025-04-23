package main

import (
	"math/rand"
	"time"
)

// Elementos concorrentes do jogo
var (
	PortalMagico = Elemento{'⭕', CorAzul, CorPadrao, false}
	ArmadilhaElem = Elemento{'💣', CorVermelho, CorPadrao, false}
	PatrollerElem = Elemento{'👾', CorRoxo, CorPadrao, true}
)

// Canais de comunicação
var (
	canalAvistamento = make(chan struct{})
	canalPortal      = make(chan struct{})
	canalArmadilha   = make(chan struct{})
)

// Patroller representa um inimigo que patrulha o mapa
type Patroller struct {
	x, y          int
	modoPerseguir bool
}

// Portal representa um portal mágico que aparece e desaparece
type Portal struct {
	x, y     int
	ativo    bool
	timeout  time.Duration
}

// ArmadilhaStruct representa uma armadilha explosiva
type ArmadilhaStruct struct {
	x, y     int
	ativa    bool
	timeout  time.Duration
}

// Inicia todas as goroutines dos elementos concorrentes
func iniciarElementosConcorrentes(jogo *Jogo) {
	// Inicia o patroller
	go func() {
		p := &Patroller{x: 5, y: 5}
		for {
			select {
			case <-canalAvistamento:
				p.modoPerseguir = true
				time.Sleep(2 * time.Second)
				p.modoPerseguir = false
			default:
				if p.modoPerseguir {
					p.perseguirJogador(jogo)
				} else {
					p.patrulhar(jogo)
				}
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// Inicia o portal mágico
	go func() {
		p := &Portal{timeout: 5 * time.Second}
		for {
			select {
			case <-canalPortal:
				p.ativar(jogo)
				time.Sleep(p.timeout)
				p.desativar(jogo)
			default:
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// Inicia a armadilha
	go func() {
		a := &ArmadilhaStruct{timeout: 3 * time.Second}
		for {
			select {
			case <-canalArmadilha:
				a.ativar(jogo)
				time.Sleep(a.timeout)
				a.desativar(jogo)
			default:
				time.Sleep(1 * time.Second)
			}
		}
	}()
}

// Métodos do Patroller
func (p *Patroller) patrulhar(jogo *Jogo) {
	// Move aleatoriamente
	dx := rand.Intn(3) - 1
	dy := rand.Intn(3) - 1
	
	if jogoPodeMoverPara(jogo, p.x+dx, p.y+dy) {
		jogoMoverElemento(jogo, p.x, p.y, dx, dy)
		p.x += dx
		p.y += dy
	}
}

func (p *Patroller) perseguirJogador(jogo *Jogo) {
	// Calcula direção para o jogador
	dx := 0
	if jogo.PosX > p.x {
		dx = 1
	} else if jogo.PosX < p.x {
		dx = -1
	}
	
	dy := 0
	if jogo.PosY > p.y {
		dy = 1
	} else if jogo.PosY < p.y {
		dy = -1
	}
	
	if jogoPodeMoverPara(jogo, p.x+dx, p.y+dy) {
		jogoMoverElemento(jogo, p.x, p.y, dx, dy)
		p.x += dx
		p.y += dy
	}
}

// Métodos do Portal
func (p *Portal) ativar(jogo *Jogo) {
	p.x = rand.Intn(len(jogo.Mapa[0]))
	p.y = rand.Intn(len(jogo.Mapa))
	p.ativo = true
	jogo.Mapa[p.y][p.x] = PortalMagico
}

func (p *Portal) desativar(jogo *Jogo) {
	p.ativo = false
	jogo.Mapa[p.y][p.x] = Vazio
}

// Métodos da Armadilha
func (a *ArmadilhaStruct) ativar(jogo *Jogo) {
	a.x = rand.Intn(len(jogo.Mapa[0]))
	a.y = rand.Intn(len(jogo.Mapa))
	a.ativa = true
	jogo.Mapa[a.y][a.x] = ArmadilhaElem
}

func (a *ArmadilhaStruct) desativar(jogo *Jogo) {
	a.ativa = false
	jogo.Mapa[a.y][a.x] = Vazio
} 