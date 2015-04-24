#pragma once
#include "GameScene.h"
#include "Maze.h"
#include "Puckman.h"

namespace my {

	class MazeScene :
		public GameScene
	{
	
	private:
		GamePtr		m_game;
		MazePtr		m_maze;

	public:
		MazeScene(GamePtr g);
		virtual ~MazeScene();

		virtual void update(sf::Time t);
		void prepare();

	private:
		void destroyMaze();
	};
}

