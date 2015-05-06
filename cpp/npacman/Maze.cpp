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

	MazeSectionPtr Maze::getStartSection() const {
		return getSection(29, 14);
	}

	MazeSectionPtr Maze::getGhostLairSection() const {
		return getSection(13, 13);
	}

	int Maze::getSectionsCount() {
		return m_size;
	}

	const MazeSectionMatrix Maze::getSections() {
		return (m_sections);
	}

	MazeSectionPtr Maze::getSection(int line, int col) const {
		return m_sections[(line * m_width) + col];
	}

	MazeSectionPtr Maze::findSection(float x, float y) {
		int posX = x / MAZE_SECTION_WIDTH;
		int posY = y / MAZE_SECTION_HEIGHT;

		return this->getSection(posY, posX);
	}
}
