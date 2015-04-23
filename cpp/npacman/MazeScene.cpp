#include "MazeScene.h"
#include "Wall.h"
#include "Game.h"
#include <iostream>

namespace my {

	MazeScene::MazeScene(GamePtr g) : m_game(g)
	{
		m_maze = new Maze(10, 10);
		m_children.push_back(m_game->getPlayer());
		createMaze();
	}
	
	MazeScene::~MazeScene()
	{
	}

	void MazeScene::update(sf::Time t) {
		
		auto player		= m_game->getPlayer();
		
		player->update(t);

		auto playerBounds = player->getGlobalBounds();

		///*
		//Varrendo todos os elementos a partir do segundo, justamente as paredes do labirinto
		for (auto p = ++m_children.begin(); p != m_children.end(); p++) {
			auto wall = dynamic_cast<Wall*>(*p);
			auto wallBounds = wall->getGlobalBounds();
						
			//Verifica se o jogador está tocando na parede
			if (playerBounds.intersects(wallBounds)) {
		
				auto distance = t.asSeconds() * DEFAULT_GAME_SPEED;
				auto newPosition = player->getPosition();
								
				if (player->isInVertical())
				{
					//Jogador tocou o topo de uma parede:
					if (playerBounds.top < wallBounds.top) {
						newPosition.y = playerBounds.top - ((playerBounds.top + playerBounds.height) - wallBounds.top) - distance;
					}
					//...ou tocou a base da parede
					else {
						newPosition.y = (wallBounds.top + wallBounds.height) + distance;
					}

				} else {

					//Jogador tocou o lado esquerdo de uma parede:
					if (playerBounds.left < wallBounds.left) {
						newPosition.x = playerBounds.left - ((playerBounds.left + playerBounds.width) - wallBounds.left)  - distance;
					}
					//...ou tocou o lado direito da parede
					else {
						newPosition.x = (wallBounds.left + wallBounds.width) + distance;
					}
				}
				
				player->setPosition(newPosition);
				player->stop();

				break;
			}
		}
		//*/
	}

	void MazeScene::createMaze() {
		auto size = m_game->getSize();
		auto defaultHeight = 5;
				
		m_children.push_back(new Wall(0, 0, size.x, defaultHeight));
		m_children.push_back(new Wall(0, 150, size.x, defaultHeight));
		m_children.push_back(new Wall(150, 50, defaultHeight, 50));
		m_children.push_back(new Wall(50, 50, defaultHeight, 50));
		
	}

	void MazeScene::destroyMaze() {

	}
}
