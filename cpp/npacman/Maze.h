#pragma once
#include "MazeSection.h"

namespace my {

	class Maze;
	typedef Maze* MazePtr;

	//Representa a organiza��o l�gica do labirinto e suas se��es
	class Maze
	{

	private:
		MazeSectionPtr	*m_sections;
		int				m_width;
		int				m_height;
		
	public:
		Maze(int width, int height);
		~Maze();

		MazeSectionPtr get(int line, int col);

	private:
		void init();
	};

}
