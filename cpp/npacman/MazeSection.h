#pragma once
#include <utility>
#include <memory>

namespace my {

	class MazeSection;
	typedef MazeSection*	MazeSectionPtr;
	typedef MazeSectionPtr*	MazeSectionMatrix;

	class MazeSection
	{
	
	private:
		std::pair<int, int> m_id;
	
	public:
		MazeSectionPtr	N;
		MazeSectionPtr	S;
		MazeSectionPtr	W;
		MazeSectionPtr	E;

		bool			allowed;
				
	public:
		MazeSection(std::pair<int, int> id);
		MazeSection(int id_x, int id_y);
		virtual ~MazeSection();

		std::pair<int, int> getID();

		bool MazeSection::operator ==(const MazeSection &b) const {
			return (m_id.first == b.m_id.first) && (m_id.second == b.m_id.second);
		}

		bool MazeSection::operator !=(const MazeSection &b) const {
			return !((*this) == b);
		}
	};

}
