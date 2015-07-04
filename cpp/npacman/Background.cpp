#include "Background.h"

namespace my {

	Background::Background(sf::Texture* texture)
	{
		m_el = new sf::Sprite();
		m_el->setPosition(sf::Vector2f(0,0));
		m_el->setTexture(*texture);
	}
	
	Background::~Background()
	{
		delete m_el;
	}
}