#include "Game.h"
#include "MazeUtils.h"
#include "Constants.h"
#include <iostream>
#include <utility>
#include <sstream>
#include <memory>
#include "MazeUtils.h"

namespace my {

	MazePtr buildDefaultMaze(my::GamePtr g) {

		auto rm = g->getResourceManager();
		auto defaultMazeMap = rm.getDefaultMazeBlueprints();

		int width = 0, height = 0;
		auto sections = buildMazeSections(*defaultMazeMap, width, height);

		return new Maze(sections, width, height);
	}

	MazeSectionMatrix buildMazeSections(const std::string& referenceMap, int& width, int& height) {

		auto strReferenceMap = std::stringstream(referenceMap);
		
		//Os primeiros dados dizem respeito a largura e altura do mapa
		strReferenceMap >> width;
		strReferenceMap >> height;
		
		MazeSectionPtr* sections = new MazeSectionPtr[width * height];
		MazeSectionPtr lastE = nullptr, lastN = nullptr;
		char c;
		int qtde = 0;

		for (int line = 0; line < height; line++) {

			//�ndice da primeira coluna da linha
			int firstCol = line * width;

			for (int col = 0; col < width; col++) {
				
				strReferenceMap >> c;
				//std::cout << c;
								
				MazeSectionPtr s = new MazeSection(std::make_pair(line, col));
				s->E = lastE;
				s->allowed = (c == MAZE_BP_PATH_BLOCK);

				//Redefinindo o 'oeste' o �ltimo leste
				if (lastE != nullptr) {
					lastE->W = s;
				}

				//Redefinindo o 'sul' do �ltimo norte e avan�ando o norte para o elemento ao lado
				if (line > 0) {
					s->N = sections[((line - 1) * width) + col];
					s->N->S = s;
				}
				
				lastE = s;
				sections[firstCol + col] = s;
				qtde++;

			}

			//strReferenceMap.get();

			//O �ltimo elemento N ser� o primeiro da linha que acaba de ser processada
			//lastN = sections[firstCol];
			lastE = nullptr;
		}

		return MazeSectionMatrix(sections);
	}

	MazeSectionMatrix buildMazeSections(int width, int height) {

		MazeSectionPtr* sections = new MazeSectionPtr[width * height];
		MazeSectionPtr lastE = nullptr, lastN = nullptr;

		for (int line = 0; line < height; line++) {

			//�ndice da primeira coluna da linha
			int firstCol = line * width;

			for (int col = 0; col < width; col++) {

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
				sections[firstCol + col] = s;
			}

			//O �ltimo elemento N ser� o primeiro da linha que acaba de ser processada
			lastN = sections[firstCol];
			lastE = nullptr;
		}

		return MazeSectionMatrix(sections);
	}
}