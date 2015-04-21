#pragma once
#include <SFML\Graphics.hpp>
#include "GameScene.h"
#include "Puckman.h"

namespace my {

	//Representa o jogo propriamente dito e seu conjunto de cenas
	class Game
	{

	private:
		const float			DEFAULT_SPEED = 60.0;
		sf::RenderWindow*	m_canvas;
		GameScenePtr		m_current_scene;
		PuckmanPtr			m_player;
		sf::FloatRect		m_bounds;

	private:
		void handleRender();
		void handleEvents();
		void handleUpdates(sf::Time t);

	public:
		Game();
		~Game();
		
		PuckmanPtr			getPlayer() const;
		const sf::Vector2u&	getSize() const;
		void				run();

	};
}
