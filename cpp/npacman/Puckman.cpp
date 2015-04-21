#include "Puckman.h"

namespace my {
	
	Puckman::Puckman()
	{
		auto b = new sf::CircleShape();
		b->setFillColor(sf::Color::Yellow);
		b->setRadius(10);
		
		m_el = b;
	}

	Puckman::~Puckman()
	{
		delete m_el;
	}

	void Puckman::init() {
		
	}

	void Puckman::update(sf::Time t) {
		m_el->move(m_facing_dir);
	}

	void Puckman::faceLeft() {
		m_facing_dir = sf::Vector2f(-1, 0);
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