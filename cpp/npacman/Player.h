#pragma once
#include "Constants.h"
#include "GameElement.h"
#include "MazeSection.h"
#include "Controllable.h"

namespace my {

	template<typename T>
	class Player : public GameElement<T>, public Controllable
	{
	
	protected:

		const float				m_velocity = DEFAULT_GAME_SPEED;
		
		MazeSectionPtr			m_current_section;
		MazeSectionPtr			m_last_section;
				
	public:

		Player(T* s) {
			m_el = s;
			m_current_section = nullptr;
			m_last_section = nullptr;
		}

		virtual ~Player() {

		}
					
		virtual void update(sf::Time t) {
			auto id = m_current_section->getID();
			auto newPos = sf::Vector2f(id.x * MAZE_SECTION_WIDTH, id.y * MAZE_SECTION_WIDTH);

			m_el->setPosition(newPos);
		}

		//Recupera zona do labirinto onde o jogador est� posicionado
		const MazeSectionPtr getLocation() const {
			return m_current_section;
		}

		//Recupera �ltima zona do labirinto visitada pelo jogador
		const MazeSectionPtr getPreviousLocation() const {
			return m_last_section;
		}

		//Redefine a posi��o do jogador dentro do labirinto (ATEN��O: Refatorar, esse m�todo n�o deve existir, todo movimento deve ser realizado pelo 'goTo')
		void setLocation(MazeSectionPtr s) {
			m_current_section = s;
		}
		
		//Move jogador para a se��o da esquerda, caso poss�vel
		void goLeft() {
			goTo(m_current_section->W);
		}

		//Move jogador para a se��o acima, caso poss�vel
		void goUp() {
			goTo(m_current_section->N);
		}

		//Move jogador para a se��o abaixo, caso poss�vel
		void goDown() {
			goTo(m_current_section->S);
		}

		//Move jogador para a se��o da direita, caso poss�vel
		void goRight() {
			goTo(m_current_section->E);
		}


	protected:

		//Move jogador para a zona do labirinto informada
		void goTo(MazeSectionPtr s) {

			if (s != nullptr && s->allowed) {
				m_last_section = m_current_section;
				m_current_section = s;
			}
		}
	};

}

