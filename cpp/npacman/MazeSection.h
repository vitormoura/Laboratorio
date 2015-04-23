#pragma once
#include <utility>

namespace my {

	class MazeSection;
	typedef MazeSection* MazeSectionPtr;

	class MazeSection
	{
	
	private:
		std::pair<int, int> id;
	
	public:
		MazeSectionPtr	N;
		MazeSectionPtr	S;
		MazeSectionPtr	W;
		MazeSectionPtr	E;

	public:
		MazeSection(std::pair<int, int> id);
		~MazeSection();
	};

}
