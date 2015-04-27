#include "Puckman.h"
#include "Constants.h"
#include <iostream>

namespace my {
	
	Puckman::Puckman() : PuckmanPlayerType(new sf::CircleShape())
	{
		sf::CircleShape* me = dynamic_cast<sf::CircleShape*>(m_el);

		me->setFillColor(sf::Color::Yellow);
		me->setRadius(MAZE_SECTION_WIDTH / 2);
	}

	Puckman::~Puckman()
	{
		#if _DEBUG
		std::cout << "Puckman::~Puckman" << std::endl;
		#endif

		delete m_el;
	}
}