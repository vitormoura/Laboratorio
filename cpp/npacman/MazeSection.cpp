#include "MazeSection.h"
#include <iostream>
#include <utility>

namespace my {

	MazeSection::MazeSection(sf::Vector2i id) : 
		m_id(id), W(nullptr), S(nullptr), N(nullptr), E(nullptr)
	{
	}

	MazeSection::MazeSection(int id_x, int id_y) : 
		m_id(sf::Vector2i(id_x, id_y)), W(nullptr), S(nullptr), N(nullptr), E(nullptr) 
	{
	}
	
	MazeSection::~MazeSection()
	{
		#if _DEBUG
		//std::cout << "MazeSection::~MazeSection" << std::endl;
		#endif
	}

	sf::Vector2i MazeSection::getID() {
		return m_id;
	}

}