#include "Maze.h"
#include <utility>
#include <iostream>

namespace my {

	Maze::Maze(const std::string& referenceMap) {
		//TODO		
	}

	Maze::Maze(int width, int height) :
		m_width(width), m_height(height)
	{
		m_sections = new MazeSectionPtr[m_width * m_height];

		init();
	}
	
	Maze::~Maze()
	{
		delete[] m_sections;

		#if _DEBUG
		std::cout << "Maze::~Maze" << std::endl;
		#endif
	}

	MazeSectionPtr Maze::get(int line, int col) {
		return m_sections[line * col];
	}

	void Maze::init() {

		MazeSectionPtr lastE = nullptr, lastN = nullptr;
		
		for (int line = 0; line < m_height; line++) {
			
			//�ndice da primeira coluna da linha
			int firstCol = line * m_width;
									
			for (int col = 0; col < m_width; col++) {
				
				MazeSectionPtr s = new MazeSection(std::make_pair(line, col));
				s->E = lastE;
				s->N = lastN;

				//Redefinindo o 'oeste' o �ltimo leste
				if (lastE != nullptr) {
					lastE->W = s;
				}

				//Redefinindo o 'sul' do �ltimo norte e avan�ando o norte para o elemento ao lado
				if (lastN != nullptr) {
					lastN->S = s;
					lastN++;
				}

				lastE = s;
				m_sections[firstCol + col] = s;
			}

			//O �ltimo elemento N ser� o primeiro da linha que acaba de ser processada
			lastN = m_sections[firstCol];
			lastE = nullptr;
		}
	}
}
