#include "Puckman.h"
#include "Constants.h"

namespace my {
	
	Puckman::Puckman() : PuckmanPlayerType::Player(new sf::CircleShape())
	{
		sf::CircleShape* me = dynamic_cast<sf::CircleShape*>(m_el);

		me->setFillColor(sf::Color::Yellow);
		me->setRadius(MAZE_SECTION_WIDTH / 2);
	}

	Puckman::~Puckman()
	{
		delete m_el;
	}
}