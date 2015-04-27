#pragma once
#include "PlayerController.h"
#include "Ghost.h"

namespace my {

	//Controlador autom�tico b�sico de um fantasma
	class SimpleGhostController : public PlayerController
	{
	
	private:
		float		m_wait;
		GhostPtr	m_ghost;

	public:
		SimpleGhostController(GhostPtr ghost);
		virtual ~SimpleGhostController();

		virtual void update(sf::Time t);
	};

}
