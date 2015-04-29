#pragma once
#include <utility>
#include <memory>
#include <SFML\Graphics.hpp>
#include "Enums.h"
#include <array>

namespace my {

	class MazeSection;
	typedef MazeSection*	MazeSectionPtr;
	typedef MazeSectionPtr*	MazeSectionMatrix;

	class MazeSection
	{
	
	private:
		sf::Vector2i	m_id;
	
	public:
		MazeSectionPtr	N;
		MazeSectionPtr	S;
		MazeSectionPtr	W;
		MazeSectionPtr	E;

		bool			allowed;
				
	public:
		MazeSection(sf::Vector2i id);
		MazeSection(int id_x, int id_y);
		virtual ~MazeSection();

		const sf::Vector2i getID() const;

		MazeSectionPtr get(Directions d) const;

		bool MazeSection::operator ==(const MazeSection &b) const {
			return (m_id.x == b.m_id.x) && (m_id.y == b.m_id.y);
		}

		bool MazeSection::operator !=(const MazeSection &b) const {
			return !((*this) == b);
		}
	};

}
