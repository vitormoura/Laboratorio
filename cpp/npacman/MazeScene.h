#pragma once
#include "GameScene.h"
#include "Maze.h"
#include "Puckman.h"
#include "Ghost.h"

namespace my {

	class MazeScene :
		public GameScene
	{
	
	private:
		GamePtr		m_game;
		MazePtr		m_maze;

		GhostPtr	m_ghosts[1];

	public:
		MazeScene(GamePtr g);
		virtual ~MazeScene();

		virtual void update(sf::Time t);
		void prepare();

	private:
		void destroyMaze();
	};
}

