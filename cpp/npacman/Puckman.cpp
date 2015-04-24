#include "Puckman.h"
#include "Constants.h"

namespace my {
	
	Puckman::Puckman()
	{
		auto b = new sf::CircleShape();
		b->setFillColor(sf::Color::Yellow);
		b->setRadius(5);
		
		m_el = b;
		m_current_section = nullptr;
	}

	Puckman::~Puckman()
	{
		delete m_el;
	}

	void Puckman::init() {
		
	}
		
	void Puckman::update(sf::Time t) {
				
		auto id = m_current_section->getID();
		auto newPos = sf::Vector2f( (id.first * MAZE_SECTION_WIDTH), id.second * MAZE_SECTION_WIDTH);

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

	void Puckman::goLeft() {
		if (m_current_section->E != nullptr) {
			m_current_section = m_current_section->E;
			m_facing_dir = sf::Vector2f(-1, 0);
			m_next_move = std::bind(&Puckman::goLeft,this);
		}
		else {
			m_next_move = nullptr;
		}
	}

	void Puckman::goUp() {
		if (m_current_section->N != nullptr) {
			m_current_section = m_current_section->N;
			m_facing_dir = sf::Vector2f(0, -1);
			m_next_move = std::bind(&Puckman::goUp,this);
		}
		else {
			m_next_move = nullptr;
		}
	}

	void Puckman::goDown() {
		if (m_current_section->S != nullptr) {
			m_current_section = m_current_section->S;
			m_facing_dir = sf::Vector2f(0, 1);
			m_next_move = std::bind(&Puckman::goDown,this);
		}
		else {
			m_next_move = nullptr;
		}
	}

	void Puckman::goRight() {
		if (m_current_section->W != nullptr) {
			m_current_section = m_current_section->W;
			m_facing_dir = sf::Vector2f(1, 0);
			m_next_move = std::bind(&Puckman::goRight,this);
		}
		else {
			m_next_move = nullptr;
		}
	}
}