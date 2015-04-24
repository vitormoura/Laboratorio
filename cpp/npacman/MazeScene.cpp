#include "MazeScene.h"
#include "Wall.h"
#include "Game.h"
#include "MazeUtils.h"
#include <iostream>

namespace my {

	MazeScene::MazeScene(GamePtr g) : m_game(g)
	{
		m_maze = buildDefaultMaze(g);
		m_children.push_back(m_game->getPlayer());
	}
	
	MazeScene::~MazeScene()
	{
		delete m_maze;
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

	void MazeScene::prepare(const std::string& map) {
		
		int line = 0, col = 0;
		int defaultSize = 15;

		for (const auto& c : map) {
			
			if (c == '\n') {
				line++;
				col = 0;
				continue;
			}
						
			if (c == MAZE_BP_FILLED_BLOCK) {
				auto w = new Wall((col * defaultSize), (line * defaultSize), defaultSize, defaultSize);
				w->setFillColor(sf::Color(0xCC,0xCC,0xCC, 0xFF));
				
				m_children.push_back(w);
			}
			
			col++;
		}
	}
		
	void MazeScene::destroyMaze() {

	}
}
