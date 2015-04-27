#pragma once
#include "PlayerController.h"

namespace my {

	//Controlador de jogo baseado em inputs externos obtidos por eventos da janela que executa o jogo
	class InputPlayerController : public PlayerController
	{

	private:
		sf::Window*		m_window;
		ControllablePtr m_target;

		
	public:
		InputPlayerController(sf::Window* win, ControllablePtr target );
		~InputPlayerController();

		virtual void update(sf::Time t);
	};
}

