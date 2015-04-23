#include "Puckman.h"
#include "Constants.h"

namespace my {
	
	Puckman::Puckman()
	{
		auto b = new sf::CircleShape();
		b->setFillColor(sf::Color::Yellow);
		b->setRadius(5);
		
		m_el = b;
	}

	Puckman::~Puckman()
	{
		delete m_el;
	}

	void Puckman::init() {
		
	}
		
	void Puckman::update(sf::Time t) {
		m_el->move(m_facing_dir * t.asSeconds() * DEFAULT_GAME_SPEED);
	}

	bool Puckman::isInHorizontal() {
		return !isInVertical();
	}

	bool Puckman::isInVertical() {
		return m_facing_dir.x == 0;
	}

	void Puckman::stop() {
		m_facing_dir.x = 0;
		m_facing_dir.y = 0;
	}

	void Puckman::faceLeft() {
		m_facing_dir = sf::Vector2f( -1, 0);
	}

	void Puckman::faceUp() {
		m_facing_dir = sf::Vector2f(0, -1);
	}

	void Puckman::faceDown() {
		m_facing_dir = sf::Vector2f(0, 1);
	}

	void Puckman::faceRight() {
		m_facing_dir = sf::Vector2f(1, 0);
	}
}