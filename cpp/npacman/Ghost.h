#pragma once
#include "Player.h"

namespace my {
	
	class Ghost;
	typedef Ghost* GhostPtr;
	typedef Player<sf::RectangleShape> GhostPlayerType;

	class Ghost : public GhostPlayerType
	{

	private:
		float m_wait;

	public:
		Ghost();
		virtual ~Ghost();

		virtual void update(sf::Time t);
	};

}
