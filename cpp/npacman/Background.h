#pragma once
#include <SFML\Graphics.hpp>
#include "GameElement.h"

namespace my {
	
	class Background;
	typedef Background* BackgroundPtr;
	typedef GameElement<sf::Sprite> BackgroundGEType;

	class Background : public GameElement<sf::Sprite>
	{
	
	public:
		Background(sf::Texture* texture);
		virtual ~Background();
	};
}

