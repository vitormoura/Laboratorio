#include "InputPlayerController.h"

namespace my {

	InputPlayerController::InputPlayerController(sf::Window* win, ControllablePtr target) : m_window(win), m_target(target)
	{
	}

	void InputPlayerController::update(sf::Time t) {

		sf::Event event;

		while (m_window->pollEvent(event))
		{
			if (event.type == sf::Event::Closed)
				m_window->close();


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
