#include "MazeSection.h"
#include <iostream>

namespace my {

	MazeSection::MazeSection(std::pair<int, int> id) : 
		m_id(id), W(nullptr), S(nullptr), N(nullptr), E(nullptr)
	{
	}
	
	MazeSection::~MazeSection()
	{
		#if _DEBUG
		std::cout << "MazeSection::~MazeSection" << std::endl;
		#endif
	}

	std::pair<int, int> MazeSection::getID() {
		return m_id;
	}

}