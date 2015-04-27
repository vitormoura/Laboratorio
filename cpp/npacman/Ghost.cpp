#include "Ghost.h"

namespace my {

	Ghost::Ghost() : GhostPlayerType(new sf::RectangleShape())
	{
		m_el->setFillColor(sf::Color::Red);
		m_el->setSize(sf::Vector2f(MAZE_SECTION_WIDTH, MAZE_SECTION_WIDTH));
	}

	Ghost::~Ghost()
	{
		delete m_el;
	}

	void Ghost::update(sf::Time t) {
		
		///*
		if (m_wait > 1) {

			if (m_current_section->N->allowed && m_last_section != m_current_section->N) {
				goUp();
			}
			else if (m_current_section->W->allowed && m_last_section != m_current_section->W) {
				goRight();
			}
			else if (m_current_section->S->allowed && m_last_section != m_current_section->S) {
				goDown();
			}
			else {
				goLeft();
			}

			//Path finding
			//
			// 
			// E1-2
			//
			//

			m_wait = 0.0f;
		}
		else {
			m_wait += t.asSeconds();
		}
		//*/

		GhostPlayerType::update(t);
	}
}