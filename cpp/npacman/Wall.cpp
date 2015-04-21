#include "Wall.h"

namespace my {

	Wall::Wall(int x, int y, int width, int height)
	{
		m_el = new sf::RectangleShape(sf::Vector2f(width, height));
		
		m_el->setFillColor(sf::Color::Cyan);
		m_el->setPosition(x, y);
	}
	
	Wall::~Wall()
	{
		delete m_el;
	}
}