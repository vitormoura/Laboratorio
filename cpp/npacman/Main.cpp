#include "Game.h"
#include <iostream>
#include "All.tests.h"

int main()
{
	#if TESTING == 0
	
	my::Game* g = new my::Game();
	g->run();
		
	delete g;
	
	#else
	runTests();
	#endif
}

