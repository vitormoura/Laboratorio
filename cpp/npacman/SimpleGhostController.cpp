#include "SimpleGhostController.h"

namespace my {

	SimpleGhostController::SimpleGhostController(GhostPtr ghost) : m_ghost(ghost)
	{
	}
	
	SimpleGhostController::~SimpleGhostController()
	{
	}

	void SimpleGhostController::update(sf::Time t) {

		auto current	= m_ghost->getLocation();
		auto last		= m_ghost->getPreviousLocation();
				
		if (m_wait > 0.5f) {

			if (current->N->allowed && last != current->N) {
				m_ghost->goUp();
			}
			else if (current->W->allowed && last != current->W) {
				m_ghost->goRight();
			}
			else if (current->S->allowed && last != current->S) {
				m_ghost->goDown();
			}
			else {
				m_ghost->goLeft();
			}

			m_wait = 0.0f;
		}
		else 
		{
			m_wait += t.asSeconds();
		}
	}
}
