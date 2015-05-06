#include <iostream>
#include "Ghost.h"
#include "SimplePlayerController.h"

namespace my {

	Ghost::Ghost(MazePtr m) : GhostPlayerType(new sf::RectangleShape(), m)
	{
		m_el->setFillColor(sf::Color::Red);
		m_el->setSize(sf::Vector2f(MAZE_SECTION_WIDTH, MAZE_SECTION_WIDTH));
	}

	Ghost::~Ghost()
	{
		#if _DEBUG
		std::cout << "Ghost::~Ghost" << std::endl;
		#endif

		delete m_el;
	}
		
	void Ghost::update(sf::Time t) {
		GhostPlayerType::update(t);
	}
}