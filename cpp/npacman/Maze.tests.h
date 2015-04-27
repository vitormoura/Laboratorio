#include <iostream>
#include "Funcoes.h"
#include "ResourceManager.h"
#include "Maze.h"
#include "MazeUtils.h"
#include <exception>

namespace my {

	namespace tests {

		namespace maze {
						
			void test_Maze() {

				my::ResourceManager rm;
				int width, height;

				my::MazeSectionMatrix sections = my::buildMazeSections(*rm.getDefaultMazeBlueprints(), width, height);
				my::Maze maze(sections, width, height);

				testCase("QUANTIDADE TOTAL DE SEÇÕES");
				assertTrue(maze.getSectionsCount() == 840, "Qtde secoes 840");

				///*

				testCase("ELEMENTO INICIAL (0,0)");
				auto s1 = maze.getSection(0, 0);
				auto id1 = s1->getID();
				assertTrue(id1.x == 0 && id1.y == 0, "ID corretos (0,0)");
				assertTrue(s1->N == nullptr, "N nullptr");
				assertTrue(s1->S != nullptr, "S !nullptr");
				assertTrue(s1->W == nullptr, "E nullptr");
				assertTrue(s1->E != nullptr, "W !nullptr");
				//*/

				///*
				testCase("POSICAO INICIAL NO LABIRINTO");
				auto s = maze.getStartSection();
				auto id = s->getID();

				assertTrue(id.y == 28 && id.x == 14, "Secao localizada (28,14)");
				assertTrue(s->N != nullptr, "N !nullptr");
				assertTrue(s->S != nullptr, "S !nullptr");
				assertTrue(s->W != nullptr, "E !nullptr");
				assertTrue(s->E != nullptr, "W !nullptr");

				assertTrue(!s->N->allowed && !s->S->allowed, "S e N !allowed");
				assertTrue(s->W->allowed && s->E->allowed, "S e N allowed");

				//assertTrue(s->W->allowed && s->E->allowed, "W e E livres");
				//*/
			}
		}
	}
}