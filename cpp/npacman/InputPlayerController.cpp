#include "InputPlayerController.h"
#include "Game.h"

namespace my {

	InputPlayerController::InputPlayerController(GamePtr game, ControllablePtr target) : m_game(game), m_window(nullptr), m_target(target)
	{
		m_window = game->getCanvas();
	}

	void InputPlayerController::update(sf::Time t) {

		sf::Event event;

		while (m_window->pollEvent(event))
		{
			if (event.type == sf::Event::Closed)
				m_game->end();


			if (sf::Keyboard::isKeyPressed(sf::Keyboard::Left)) {
				m_target->goLeft();
			}
			else if (sf::Keyboard::isKeyPressed(sf::Keyboard::Right)) {
				m_target->goRight();
			}
			else if (sf::Keyboard::isKeyPressed(sf::Keyboard::Up)) {
				m_target->goUp();
			}
			else if (sf::Keyboard::isKeyPressed(sf::Keyboard::Down)) {
				m_target->goDown();
			}
		}
	}
	
	InputPlayerController::~InputPlayerController()
	{
	}
}
