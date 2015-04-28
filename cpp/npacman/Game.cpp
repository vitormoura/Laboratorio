#include <iostream>
#include "Game.h"
#include "Puckman.h"
#include "MazeScene.h"


namespace my {
		
	Game::Game()
	{
		m_canvas = new sf::RenderWindow(sf::VideoMode(460, 460), "NPACMAN");
		m_canvas->setFramerateLimit(DEFAULT_GAME_SPEED);

		m_puckman = new Puckman();
		m_puckman_ctrl = new InputPlayerController(m_canvas, m_puckman);
				
		m_current_scene = new MazeScene(this);
	}

	Game::~Game()
	{
		#if _DEBUG
		std::cout << "Game::~Game" << std::endl;
		#endif

		delete m_canvas;
		delete m_puckman_ctrl;
		delete m_puckman;
		delete m_current_scene;
	}

	const ResourceManager& Game::getResourceManager() {
		return m_rm;
	}

	const sf::Vector2u&	Game::getSize() const {
		return m_canvas->getSize();
	}

	PuckmanPtr Game::getPlayer() const {
		return m_puckman;
	}

	PlayerControllerPtr	Game::getPlayerController() const {
		return m_puckman_ctrl;
	}

	void Game::run() {

		sf::Clock clock;
		sf::Time timeSinceLastUpdate = sf::Time::Zero;
		sf::Time timePerFrame = sf::seconds(1.f / Game::DEFAULT_SPEED);

		while (m_canvas->isOpen())
		{
			handleUpdates(timePerFrame);

			timeSinceLastUpdate += clock.restart();
						
			while (timeSinceLastUpdate > timePerFrame) {
				timeSinceLastUpdate -= timePerFrame;

				handleUpdates(timePerFrame);
			}
			
			handleRender();
		}
	}

	void Game::handleRender() {

		m_canvas->clear();
		m_current_scene->render(m_canvas);
		m_canvas->display();
	}
		
	void Game::handleUpdates(sf::Time t) {
		m_puckman_ctrl->update(t);
		m_current_scene->update(t);
	}

	
}