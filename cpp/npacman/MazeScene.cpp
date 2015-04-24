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
		
		auto player		= m_game->getPlayer();
		
		player->update(t);

		auto playerBounds = player->getGlobalBounds();

		/*
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

	void MazeScene::prepare() {
		
		auto sections = m_maze->getSections();
		auto size = m_maze->getSectionsCount();
		auto defaultSize = MAZE_SECTION_WIDTH;

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
