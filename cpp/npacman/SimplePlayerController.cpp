#include "SimplePlayerController.h"

namespace my {

	SimplePlayerController::SimplePlayerController(ControllablePtr target) : m_target(target)
	{
	}
	
	SimplePlayerController::~SimplePlayerController()
	{
	}

	void SimplePlayerController::update(sf::Time t) {

		auto current	= m_target->getLocation();
		auto last		= m_target->getPreviousLocation();
				
		if (m_wait > 0.5f) {

			if (current->N->allowed && last != current->N) {
				m_target->goUp();
			}
			else if (current->E->allowed && last != current->E) {
				m_target->goRight();
			}
			else if (current->S->allowed && last != current->S) {
				m_target->goDown();
			}
			else {
				m_target->goLeft();
			}

			m_wait = 0.0f;
		}
		else 
		{
			m_wait += t.asSeconds();
		}
	}
}
