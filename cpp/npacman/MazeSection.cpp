#include "MazeSection.h"
#include <iostream>

namespace my {

	MazeSection::MazeSection(std::pair<int, int> id) : 
		id(id), W(nullptr), S(nullptr), N(nullptr), E(nullptr)
	{
	}
	
	MazeSection::~MazeSection()
	{
		#if _DEBUG
		std::cout << "MazeSection::~MazeSection" << std::endl;
		#endif
	}

}