#pragma once
#include "Constants.h"
#include "Enums.h"
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
		MazeSectionPtr			m_next_section;
		Directions				m_facing_direction;

		float					m_moving_duration;
										
	public:

		Player(T* s) {
			m_el = s;
			m_current_section	= nullptr;
			m_last_section		= nullptr;
			m_next_section		= nullptr;
			m_facing_direction	= Directions::E;
			m_moving_duration = 0.25;
		}

		virtual ~Player() {

		}
					
		virtual void update(sf::Time t) {
			auto id = m_current_section->getID();
			m_el->setPosition(sf::Vector2f(id.x * MAZE_SECTION_WIDTH, id.y * MAZE_SECTION_HEIGHT));
		}

		//Recupera zona do labirinto onde o jogador está posicionado
		const MazeSectionPtr getLocation() const {
			return m_current_section;
		}

		//Recupera última zona do labirinto visitada pelo jogador
		const MazeSectionPtr getPreviousLocation() const {
			return m_last_section;
		}

		//Redefine a posição do jogador dentro do labirinto (ATENÇÃO: Refatorar, esse método não deve existir, todo movimento deve ser realizado pelo 'goTo')
		void setLocation(MazeSectionPtr s) {

			if (s->allowed) {
				m_last_section = m_current_section;
				m_current_section = s;
				m_next_section = nullptr;
				m_facing_direction = Directions::E;
			}
		}
		
		//Move jogador para a seção da esquerda, caso possível
		void goLeft() {
			goTo(m_current_section->W, Directions::W);
		}

		//Move jogador para a seção acima, caso possível
		void goUp() {
			goTo(m_current_section->N, Directions::N);
		}

		//Move jogador para a seção abaixo, caso possível
		void goDown() {
			goTo(m_current_section->S, Directions::S);
		}

		//Move jogador para a seção da direita, caso possível
		void goRight() {
			goTo(m_current_section->E, Directions::E);
		}

		//Move o jogador para a direção indicada
		void go(Directions d) {
			
			switch (d) {

			case Directions::N:
				goUp();
				break;
			case Directions::E:
				goRight();
				break;
			case Directions::S:
				goDown();
				break;
			case Directions::W:
				goLeft();
				break;
			}
		}


	protected:

		//Move jogador para a zona do labirinto informada
		void goTo(MazeSectionPtr s, Directions d) {

			if (s != nullptr && s->allowed) {
				m_current_section = s;
				m_next_section = s;
				m_facing_direction = d;
			}
		}
	};

}

