#pragma once
#include "Player.h"
#include "Maze.h"


namespace my {
	
	class Puckman;
	typedef Puckman* PuckmanPtr;
	typedef Player<sf::CircleShape> PuckmanPlayerType;

	//Jogador principal, nosso amado PACMAN
	class Puckman : public PuckmanPlayerType
	{

	public:
		Puckman(MazePtr m);
		virtual ~Puckman();

	};
}
