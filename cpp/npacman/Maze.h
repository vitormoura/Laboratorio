#pragma once
#include "MazeSection.h"
#include <string>

namespace my {

	class Maze;
	typedef Maze* MazePtr;

	//Representa a organização lógica do labirinto e suas seções
	class Maze
	{

	private:
		MazeSectionPtr	*m_sections;
		int				m_width;
		int				m_height;
		
	public:
		Maze(int width, int height);
		Maze(const std::string& referenceMap);
		~Maze();

		MazeSectionPtr get(int line, int col);

	private:
		void init();
	};

}
