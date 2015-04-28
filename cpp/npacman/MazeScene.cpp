#include "MazeScene.h"
#include "Wall.h"
#include "Game.h"
#include "MazeUtils.h"
#include "AutoPlayerController.h"
#include "SimplePlayerController.h"
#include <iostream>

namespace my {

	MazeScene::MazeScene(GamePtr g) : m_game(g)
	{
		m_maze = buildDefaultMaze(g);
		
		prepareCharacters();
		prepareWalls();
	}
	
	MazeScene::~MazeScene()
	{
		delete m_maze;
		
		for (size_t i = 0; i < characters::size; i++)
			delete m_controllers[i];
	}
				
	void MazeScene::update(sf::Time t) {
		
		for (size_t i = 0; i < characters::size; i++)
			m_controllers[i]->update(t);

		GameScene::update(t);
	}

	void MazeScene::prepareCharacters() {
				
		auto puckman = new Puckman();
		puckman->setLocation(m_maze->getStartSection());
				
		auto blinky = new Ghost();
		blinky->setLocation(m_maze->getGhostLairSection());
		
		auto inky = new Ghost();
		inky->setLocation(m_maze->getGhostLairSection());

		//Definindo quais os controladores de cada personagem
		m_controllers[characters::PuckmanT] = new InputPlayerController(m_game, puckman);
		m_controllers[characters::Blinky] = new AutoPlayerController(blinky, puckman);
		m_controllers[characters::Inky] = new SimplePlayerController(inky);
		
		m_children.push_back(puckman);
		m_children.push_back(blinky);
		m_children.push_back(inky);
	}

	void MazeScene::prepareWalls() {

		auto sections = m_maze->getSections();
		auto size = m_maze->getSectionsCount();
		auto defaultSize = MAZE_SECTION_WIDTH;

		for (int i = 0; i < size; i++) {

			if (!sections[i]->allowed) {

				auto id = sections[i]->getID();
				auto w = new Wall(id.x * defaultSize, id.y * defaultSize, defaultSize, defaultSize);

				m_children.push_back(w);
			}
		}
	}	
}
