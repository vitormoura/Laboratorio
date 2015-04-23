#pragma once
#include "GameElement.h"

namespace my {

	class Wall : public GameElement<sf::RectangleShape>
	{

	public:
		Wall(int x, int y, int width, int height);
		~Wall();

		void setFillColor(sf::Color c);
	
	};

}