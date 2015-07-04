#pragma once
#include "MazeSection.h"

namespace my {

	class PathFindingNode
	{

	public:
		MazeSectionPtr		location;
		PathFindingNode*	parent;
		int					h;
		int					g;
		
	public:
		PathFindingNode();
		~PathFindingNode();

		bool PathFindingNode::operator ==	(const PathFindingNode &b) const { return *location == *b.location; }
		bool PathFindingNode::operator !=	(const PathFindingNode &b) const { return !((*this) == b);	}

		bool PathFindingNode::operator <	(const PathFindingNode& rhs){ return (h + g) < (rhs.h + rhs.g); }
		bool PathFindingNode::operator >	(const PathFindingNode& rhs){ return (h + g) > (rhs.h + rhs.g); }
		bool PathFindingNode::operator <=	(const PathFindingNode& rhs){ return (h + g) <= (rhs.h + rhs.g); }
		bool PathFindingNode::operator >=	(const PathFindingNode& rhs){ return (h + g) >= (rhs.h + rhs.g); }
	};
}

