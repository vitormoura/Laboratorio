#include "Game.h"
#include <iostream>
#include "All.tests.h"

int main()
{
	#ifdef TESTING
	
	my::Game* g = new my::Game();
	g->run();
		
	delete g;
	
	#else
	runTests();
	#endif
}

