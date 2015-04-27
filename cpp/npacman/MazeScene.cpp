#include "MazeScene.h"
#include "Wall.h"
#include "Game.h"
#include "MazeUtils.h"
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
	}

	void MazeScene::prepare() {
		
		auto sections = m_maze->getSections();
		auto size = m_maze->getSectionsCount();
		auto defaultSize = MAZE_SECTION_WIDTH;

		m_ghosts[0] = new Ghost();
		m_ghosts[0]->setLocation(m_maze->getSection(13, 14));

		m_children.push_back(m_ghosts[0]);

		for (int i = 0; i < size; i++) {

			if (!sections[i]->allowed) {

				auto id = sections[i]->getID();
				auto w = new Wall(id.second * defaultSize, id.first * defaultSize, defaultSize, defaultSize);

				m_children.push_back(w);
			}
		}
	}
		
	void MazeScene::destroyMaze() {

	}
}
