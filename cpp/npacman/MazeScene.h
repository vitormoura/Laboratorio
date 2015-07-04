#pragma once
#include "GameScene.h"
#include "Maze.h"
#include "Ghost.h"


namespace my {

	class MazeScene :
		public GameScene
	{
	
	private:
		enum characters {
			PuckmanT,
			Blinky,
			Inky,
			//Pinky,
			//Clyde,
			size
		};

		GamePtr				m_game;
		MazePtr				m_maze;
				
		PlayerControllerPtr m_controllers[characters::size];

	public:
		MazeScene(GamePtr g);
		virtual ~MazeScene();

		virtual void			update(sf::Time t);

	private:
		void					prepareCharacters();
		void					prepareWalls();
	
	};
}

