#include "MazeScene.h"
#include "Wall.h"
#include "Game.h"
#include "MazeUtils.h"
#include "AutoPlayerController.h"
#include "SimplePlayerController.h"
#include "Background.h"
#include <iostream>

namespace my {

	MazeScene::MazeScene(GamePtr g) : m_game(g)
	{
		m_maze = buildDefaultMaze(g);
		
		prepareWalls();
		prepareCharacters();
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

		auto rm = m_game->getResourceManager();
		auto bg = new Background(rm.getDefaultMazeTemplate());
		
		m_children.push_back(bg);
		
		/*
		auto sections = m_maze->getSections();
		auto size = m_maze->getSectionsCount();
		
		for (int i = 0; i < size; i++) {

			if (!sections[i]->allowed) {

				auto id = sections[i]->getID();
				auto w = new Wall(id.x * MAZE_SECTION_WIDTH, id.y * MAZE_SECTION_HEIGHT, MAZE_SECTION_WIDTH, MAZE_SECTION_HEIGHT);

				m_children.push_back(w);
			}
		}
		//*/
	}	
}
