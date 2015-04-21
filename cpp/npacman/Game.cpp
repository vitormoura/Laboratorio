#include "Game.h"
#include "Puckman.h"
#include "MazeScene.h"

namespace my {
		
	Game::Game()
	{
		m_canvas = new sf::RenderWindow(sf::VideoMode(200, 200), "NPACMAN");
		m_player = new Puckman();
		m_player->setPosition(sf::Vector2f(50, 50));
		m_current_scene = new MazeScene(this);
	}

	Game::~Game()
	{
		delete m_canvas;
		delete m_player;
	}

	const sf::Vector2u&	Game::getSize() const {
		return m_canvas->getSize();
	}

	PuckmanPtr Game::getPlayer() const {
		return m_player;
	}

	void Game::run() {

		sf::Clock clock;
		sf::Time timeSinceLastUpdate = sf::Time::Zero;
		sf::Time timePerFrame = sf::seconds(1.f / Game::DEFAULT_SPEED);

		while (m_canvas->isOpen())
		{
			handleEvents();

			timeSinceLastUpdate += clock.restart();

			while (timeSinceLastUpdate > timePerFrame) {
				timeSinceLastUpdate -= timePerFrame;

				handleEvents();
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

	void Game::handleEvents() {

		sf::Event event;

		while (m_canvas->pollEvent(event))
		{
			if (event.type == sf::Event::Closed)
				m_canvas->close();

			
			if (sf::Keyboard::isKeyPressed(sf::Keyboard::Left)) {
				m_player->faceLeft();
			}
			else if (sf::Keyboard::isKeyPressed(sf::Keyboard::Right)) {
				m_player->faceRight();
			}
			else if (sf::Keyboard::isKeyPressed(sf::Keyboard::Up)) {
				m_player->faceUp();
			}
			else if (sf::Keyboard::isKeyPressed(sf::Keyboard::Down)) {
				m_player->faceDown();
			}
		}
	}

	void Game::handleUpdates(sf::Time t) {
		m_current_scene->update(t);
	}

	
}