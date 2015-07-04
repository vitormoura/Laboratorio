#pragma once
#include "Player.h"
#include "PlayerController.h"

namespace my {
	
	class Ghost;
	typedef Ghost* GhostPtr;
	typedef Player<sf::RectangleShape> GhostPlayerType;

	class Ghost : public GhostPlayerType
	{
			
	public:
		Ghost(MazePtr m);
		virtual ~Ghost();
		

		virtual void update(sf::Time t);
	};

}
