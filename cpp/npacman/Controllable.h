#pragma once

namespace my {
	
	class Controllable;
	typedef Controllable* ControllablePtr;

	//Interface que define comportamento de elementos controláveis
	class Controllable
	{

	public:
		enum class directions {
			N,
			E,
			S,
			W
		};

	public:
		virtual const MazeSectionPtr getLocation() const = 0;
		virtual const MazeSectionPtr getPreviousLocation() const = 0;
		virtual void goLeft() = 0;
		virtual void goUp() = 0;
		virtual void goDown() = 0;
		virtual void goRight() = 0;
		virtual void go(Controllable::directions d) = 0;
	};
}