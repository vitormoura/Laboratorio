#pragma once
#include "Player.h"
#include <SFML\Graphics.hpp>

namespace my {

	class PlayerController;
	typedef PlayerController* PlayerControllerPtr;

	//Interface capaz de controlar um jogador
	class PlayerController
	{
	
	public:

		//Sinaliza uma atualização ao controlador
		virtual void update(sf::Time t) = 0;

	};
}