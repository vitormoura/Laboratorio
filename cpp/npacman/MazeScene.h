#pragma once
#include "GameScene.h"
#include "Maze.h"
#include "Puckman.h"
#include "Ghost.h"
#include "PlayerController.h"

namespace my {

	class MazeScene :
		public GameScene
	{
	
	private:
		enum ghosts {
			Blinky,
			//Inky,
			//Pinky,
			//Clyde,
			size
		};


		GamePtr				m_game;
		MazePtr				m_maze;
		GhostPtr			m_ghosts[ghosts::size];
		PlayerControllerPtr m_ghost_ctrls[ghosts::size];

	public:
		MazeScene(GamePtr g);
		virtual ~MazeScene();

		virtual void update(sf::Time t);
		void prepare();

	private:
		void destroyMaze();
	};
}

