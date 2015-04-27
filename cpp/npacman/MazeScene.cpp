#include "MazeScene.h"
#include "Wall.h"
#include "Game.h"
#include "MazeUtils.h"
#include "SimplePlayerController.h"
#include <iostream>

namespace my {

	MazeScene::MazeScene(GamePtr g) : m_game(g)
	{
		m_maze = buildDefaultMaze(g);

		auto player = m_game->getPlayer();
		player->setLocation(m_maze->getStartSection());
				
		m_children.push_back(player);
				
		prepare();
	}
	
	MazeScene::~MazeScene()
	{
		delete m_maze;
	}

	void MazeScene::update(sf::Time t) {
						
		GameScene::update(t);
		
		for (size_t i = 0; i < ghosts::size; i++)
			m_ghost_ctrls[i]->update(t);
	}

	void MazeScene::prepare() {
		
		auto sections = m_maze->getSections();
		auto size = m_maze->getSectionsCount();
		auto defaultSize = MAZE_SECTION_WIDTH;

		m_ghosts[ghosts::Blinky] = new Ghost();
		m_ghosts[ghosts::Blinky]->setLocation(m_maze->getSection(13, 14));

		m_ghost_ctrls[ghosts::Blinky] = new SimplePlayerController(m_ghosts[ghosts::Blinky]);
		

		m_children.push_back(m_ghosts[ghosts::Blinky]);

		for (int i = 0; i < size; i++) {

			if (!sections[i]->allowed) {

				auto id = sections[i]->getID();
				auto w = new Wall(id.x * defaultSize, id.y * defaultSize, defaultSize, defaultSize);

				m_children.push_back(w);
			}
		}
	}
		
	void MazeScene::destroyMaze() {

	}
}
