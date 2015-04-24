#include "Puckman.h"
#include "Constants.h"

namespace my {
	
	Puckman::Puckman()
	{
		auto b = new sf::CircleShape();
		b->setFillColor(sf::Color::Yellow);
		b->setRadius(MAZE_SECTION_WIDTH / 2);
		
		m_el = b;
		m_current_section = nullptr;
		m_last_section = nullptr;
	}

	Puckman::~Puckman()
	{
		delete m_el;
	}

	void Puckman::init() {
		
	}
		
	void Puckman::update(sf::Time t) {
		
		auto id = m_current_section->getID();
		auto newPos = sf::Vector2f(id.second * MAZE_SECTION_WIDTH, id.first * MAZE_SECTION_WIDTH);
				
		m_el->setPosition(newPos);
	}

	const MazeSectionPtr Puckman::getLocation() const {
		return m_current_section;
	}

	void Puckman::setLocation(MazeSectionPtr s) {
		m_current_section = s;
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

	void Puckman::goTo(MazeSectionPtr s) {

		if (s != nullptr && s->allowed) {
			m_last_section = m_current_section;
			m_current_section = s;
		}
	}

	void Puckman::goLeft() {
		goTo(m_current_section->E);
	}

	void Puckman::goUp() {
		goTo(m_current_section->N);
	}

	void Puckman::goDown() {
		goTo(m_current_section->S);
	}

	void Puckman::goRight() {
		goTo(m_current_section->W);
	}
}