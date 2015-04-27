#include <iostream>
#include "Ghost.h"
#include "SimpleGhostController.h"

namespace my {

	Ghost::Ghost() : GhostPlayerType(new sf::RectangleShape())
	{
		m_el->setFillColor(sf::Color::Red);
		m_el->setSize(sf::Vector2f(MAZE_SECTION_WIDTH, MAZE_SECTION_WIDTH));
		m_controller = new SimpleGhostController(this);
	}

	Ghost::~Ghost()
	{
		#if _DEBUG
		std::cout << "Ghost::~Ghost" << std::endl;
		#endif

		delete m_el;
		delete m_controller;
	}

	void Ghost::update(sf::Time t) {
				
		m_controller->update(t);

		GhostPlayerType::update(t);
	}
}