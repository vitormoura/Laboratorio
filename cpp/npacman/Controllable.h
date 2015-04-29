#pragma once
#include "Enums.h"

namespace my {
	
	class Controllable;
	typedef Controllable* ControllablePtr;

	//Interface que define comportamento de elementos controláveis
	class Controllable
	{
			
	public:
		virtual const MazeSectionPtr getLocation() const = 0;
		virtual const MazeSectionPtr getPreviousLocation() const = 0;
		virtual void goLeft() = 0;
		virtual void goUp() = 0;
		virtual void goDown() = 0;
		virtual void goRight() = 0;
		virtual void go(Directions d) = 0;
	};
}