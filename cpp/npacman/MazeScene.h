#pragma once
#include "GameScene.h"
#include "Puckman.h"

namespace my {

	class MazeScene :
		public GameScene
	{
	
	private:
		GamePtr m_game;

	public:
		MazeScene(GamePtr g);
		~MazeScene();

		void update(sf::Time t);

	private:
		void createMaze();
		void destroyMaze();
	};
}

