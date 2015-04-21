#pragma once
#include "GameElement.h"

namespace my {
	
	class Puckman;
	typedef Puckman* PuckmanPtr;


	class Puckman : public GameElement<sf::Shape> 
	{
	
	private:
		sf::Vector2f m_facing_dir;

	public:
		Puckman();
		~Puckman();

	public:
		virtual void init();
		virtual void update(sf::Time t);

		void faceLeft();
		void faceUp();
		void faceDown();
		void faceRight();

	};
}
