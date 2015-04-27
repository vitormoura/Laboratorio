#pragma once
#include "Player.h"


namespace my {
	
	class Puckman;
	typedef Puckman* PuckmanPtr;
	typedef Player<sf::CircleShape> PuckmanPlayerType;

	//Jogador principal, nosso amado PACMAN
	class Puckman : public PuckmanPlayerType
	{

	public:
		Puckman();
		virtual ~Puckman();

	};
}
