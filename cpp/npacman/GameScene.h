#pragma once
#include <list>
#include <SFML\Graphics.hpp>

namespace my {
	
	class GameScene;
	class Game;

	typedef GameScene* GameScenePtr;
	typedef Game*		GamePtr;
	

	//Representa uma cena de jogo
	class GameScene
	{

	protected:
		GamePtr						m_game;
		GameScenePtr				m_parent;
		std::vector<GameScenePtr>	m_children;
		
	public:
		GameScene();
		virtual ~GameScene();

	public:
		virtual void init();
		virtual void update(sf::Time t);
		virtual void render(sf::RenderTarget* t);
	};

}