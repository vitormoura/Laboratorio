#pragma once
#include "Constants.h"
#include <SFML\Graphics.hpp>
#include "GameScene.h"
#include "ResourceManager.h"
#include "Puckman.h"
#include "InputPlayerController.h"

namespace my {

	class Game;
	typedef Game* GamePtr;

	//Representa o jogo propriamente dito e seu conjunto de cenas
	class Game
	{

	private:
		const float			DEFAULT_SPEED = DEFAULT_GAME_SPEED;
		sf::RenderWindow*	m_canvas;
		GameScenePtr		m_current_scene;
		
		sf::FloatRect		m_bounds;
		my::ResourceManager	m_rm;

	private:
		void handleRender();
		void handleUpdates(sf::Time t);

	public:
		Game();
		virtual ~Game();
				
		sf::RenderWindow*		getCanvas() const;
		const sf::Vector2u&		getSize() const;
		const ResourceManager&	getResourceManager();

		void					run();

	};
}
