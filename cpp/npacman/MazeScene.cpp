#include "MazeScene.h"
#include "Wall.h"
#include "Game.h"
#include "MazeUtils.h"
#include "AutoPlayerController.h"
#include <iostream>

namespace my {

	MazeScene::MazeScene(GamePtr g) : m_game(g), m_puckman(nullptr)
	{
		m_maze = buildDefaultMaze(g);
		prepare();
	}
	
	MazeScene::~MazeScene()
	{
		delete m_maze;
	}

	PuckmanPtr MazeScene::getPlayer() const {
		return m_puckman;
	}
		
	void MazeScene::update(sf::Time t) {
		
		for (size_t i = 0; i < characters::size; i++)
			m_controllers[i]->update(t);

		GameScene::update(t);
	}

	void MazeScene::prepare() {
				
		m_puckman = new Puckman();
		m_puckman->setLocation(m_maze->getStartSection());

		m_children.push_back(m_puckman);


		auto blinky = new Ghost();
		blinky->setLocation(m_maze->getSection(13, 13));
		
		m_controllers[characters::PuckmanT] = new InputPlayerController(m_game->getCanvas(), m_puckman);
		m_controllers[characters::Blinky] = new AutoPlayerController(blinky, m_puckman);
		
		m_children.push_back(blinky);


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
		
	void MazeScene::destroyMaze() {

	}
}
