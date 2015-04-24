#pragma once
#include "Puckman.h"
#include "MazeSection.h"
#include <string>
#include <memory>

namespace my {

	class Maze;
	typedef Maze* MazePtr;

	//Representa a organiza��o l�gica do labirinto e suas se��es
	class Maze
	{

	private:
		MazeSectionMatrix	m_sections;
		int					m_width;
		int					m_height;
		int					m_size;
						
	public:
		Maze(MazeSectionMatrix sections, int width, int height);
		virtual ~Maze();

		MazeSectionPtr getStartSection() const;
		int getSectionsCount();
		const MazeSectionMatrix getSections();
		MazeSectionPtr getSection(int line, int col) const;
	};

}
