#include "Maze.h"
#include "Constants.h"
#include <iostream>

namespace my {

	Maze::Maze(MazeSectionMatrix sections, int width, int height) : 
		m_sections(sections), m_width(width), m_height(height), m_size(width * height) {
	}
			
	Maze::~Maze()
	{
		#if _DEBUG
		std::cout << "Maze::~Maze (" << m_size << " secoes)" << std::endl;
		#endif

		for (int x = 0; x < m_size; x++)
			delete m_sections[x];

		delete m_sections;
	}

	int Maze::getSectionsCount() {
		return m_size;
	}

	const MazeSectionMatrix Maze::getSections() {
		return (m_sections);
	}

	MazeSectionPtr Maze::getSection(int line, int col) {
		return m_sections[line * col];
	}
}
