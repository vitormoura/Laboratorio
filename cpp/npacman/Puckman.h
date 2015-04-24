#pragma once
#include "Constants.h"
#include "GameElement.h"
#include "MazeSection.h"
#include <functional>

namespace my {
	
	class Puckman;
	typedef Puckman* PuckmanPtr;


	class Puckman : public GameElement<sf::Shape> 
	{
	
	private:
		sf::Vector2f			m_facing_dir;
		const float				m_velocity = DEFAULT_GAME_SPEED;
		std::function<void()>	m_next_move;
		MazeSectionPtr			m_current_section;
		MazeSectionPtr			m_last_section;
		
		void					goTo(MazeSectionPtr s);

	public:
		Puckman();
		virtual ~Puckman();

	public:
		virtual void init();
		virtual void update(sf::Time t);

		const MazeSectionPtr getLocation() const;
		void setLocation(MazeSectionPtr s);

		bool isInHorizontal();
		bool isInVertical();

		void stop();
		void goLeft();
		void goUp();
		void goDown();
		void goRight();

	};
}
