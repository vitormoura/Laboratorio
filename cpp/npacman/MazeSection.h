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
		virtual ~MazeSection();

		std::pair<int, int> getID();
	};

}
