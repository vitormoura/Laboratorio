#pragma once

#include "Maze.tests.h"
#include "MazeSection.tests.h"

void runTests() {

	my::tests::maze::test_Maze();
	my::tests::mzsection::test_mzSection();
	
	std::cin.get();
}