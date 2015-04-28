#pragma once
#include "MazeSection.h"
#include "PlayerController.h"
#include <utility>

namespace my {

	class AutoPlayerController;
	typedef AutoPlayerController* AutoPlayerControllerPtr;
	
	//Controlador automático
	class AutoPlayerController : public PlayerController
	{
	private:
		ControllablePtr m_target;
		ControllablePtr m_self;

		float			m_wait;

	public:
		AutoPlayerController(ControllablePtr m_self, ControllablePtr m_target);
		virtual ~AutoPlayerController();

		virtual void update(sf::Time t);
	};

}
