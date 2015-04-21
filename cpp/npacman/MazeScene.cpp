#include "MazeScene.h"
#include "Wall.h"
#include "Game.h"

namespace my {

	MazeScene::MazeScene(GamePtr g) : m_game(g)
	{
		m_children.push_back(m_game->getPlayer());
		createMaze();
	}
	
	MazeScene::~MazeScene()
	{
	}

	void MazeScene::update(sf::Time t) {
		GameScene::update(t);
	}

	void MazeScene::createMaze() {
		auto size = m_game->getSize();
		auto defaultHeight = 5;
				
		m_children.push_back(new Wall(0, 0, size.x, defaultHeight));
	}

	void MazeScene::destroyMaze() {

	}
}
