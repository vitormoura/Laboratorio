#pragma once
#include <SFML\Graphics.hpp>
#include "GameScene.h"

namespace my {

	template<typename T>
	class GameElement : public GameScene
	{

	protected:
		T* m_el;
					
	public:

		//Redefine posição do elemento
		virtual void setPosition(const sf::Vector2f& p) {
			m_el->setPosition(p);
		}

		//Recupera posição atual do elemento
		virtual const sf::Vector2f& getPosition() const {
			return m_el->getPosition();
		}

		//Move o elemento de acordo com o offset informado
		void move(const sf::Vector2f& offset) {
			m_el->move(offset);
		}
		
		//Recupera espaço ocupado pelo elemento visual
		sf::FloatRect getLocalBounds() const {
			return m_el->getLocalBounds();
		}

		//Recupera espaço ocupado pelo elemento levando em consideração suas transformações
		sf::FloatRect getGlobalBounds() const {
			return m_el->getGlobalBounds();
		}

		virtual void render(sf::RenderTarget* t) {
			GameScene::render(t);
			t->draw(*m_el);
		};
	};	
}

