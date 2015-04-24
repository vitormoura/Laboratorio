#include <iostream>
#include "ResourceManager.h"
#include "Maze.h"
#include "MazeUtils.h"
#include <exception>


void testCase(const char* name) {
	std::cout << name << std::endl;
}

void assertTrue(bool v, const char* error_msg) {
	if (!v) {
		std::cout << "[ERROR] " << error_msg << std::endl;
		return;
	}

	std::cout << "[OK   ] " << error_msg << std::endl;
}

void test_Maze() {

	my::ResourceManager rm;
	int width, height;

	my::MazeSectionMatrix sections = my::buildMazeSections(*rm.getDefaultMazeBlueprints(), width, height);
	my::Maze maze(sections, width, height);

	assertTrue(maze.getSectionsCount() == 840, "Qtde secoes 840");

	///*

	testCase("ELEMENTO INICIAL (0,0)");
	auto s1 = maze.getSection(0, 0);
	auto id1 = s1->getID();
	assertTrue(id1.first == 0 && id1.second == 0, "ID corretos (0,0)");
	assertTrue(s1->N == nullptr, "N nullptr");
	assertTrue(s1->S != nullptr, "S !nullptr");
	assertTrue(s1->E == nullptr, "E nullptr");
	assertTrue(s1->W != nullptr, "W !nullptr");
	//*/

	///*
	testCase("POSICAO INICIAL NO LABIRINTO");
	auto s = maze.getStartSection();
	auto id = s->getID();

	assertTrue(id.first == 28 && id.second == 14, "Secao localizada (28,14)");
	assertTrue(s->N != nullptr, "N !nullptr");
	assertTrue(s->S != nullptr, "S !nullptr");
	assertTrue(s->E != nullptr, "E !nullptr");
	assertTrue(s->W != nullptr, "W !nullptr");

	assertTrue(!s->N->allowed && !s->S->allowed, "S e N !allowed");
	assertTrue(s->E->allowed && s->W->allowed, "S e N allowed");

	//assertTrue(s->W->allowed && s->E->allowed, "W e E livres");
	//*/
}

void testAll() {
	test_Maze();
}