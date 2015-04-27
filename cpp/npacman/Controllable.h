#pragma once

namespace my {
	
	class Controllable;
	typedef Controllable* ControllablePtr;

	//Interface que define comportamento de elementos controláveis
	class Controllable
	{

	public:
		virtual void goLeft() = 0;
		virtual void goUp() = 0;
		virtual void goDown() = 0;
		virtual void goRight() = 0;
	};
}