#pragma once
#include "PlayerController.h"
#include "Ghost.h"

namespace my {

	//Controlador automático básico de um fantasma
	class SimplePlayerController : public PlayerController
	{
	
	private:
		float			m_wait;
		ControllablePtr	m_target;

	public:
		SimplePlayerController(ControllablePtr target);
		virtual ~SimplePlayerController();

		virtual void update(sf::Time t);
	};

}
