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

		//Redefine posi��o do elemento
		virtual void setPosition(const sf::Vector2f& p) {
			m_el->setPosition(p);
		}

		//Recupera posi��o atual do elemento
		virtual sf::Vector2f getPosition() const {
			return m_el->getPosition();
		}

		//Move o elemento de acordo com o offset informado
		void move(const sf::Vector2f& offset) {
			m_el->move(offset);
		}
		
		//Recupera espa�o ocupado pelo elemento visual
		sf::FloatRect getLocalBounds() const {
			return m_el->getLocalBounds();
		}

		//Recupera espa�o ocupado pelo elemento levando em considera��o suas transforma��es
		sf::FloatRect getGlobalBounds() const {
			return m_el->getGlobalBounds();
		}

		//Verifica se o elemento atual toca os limites do outro
		bool hit(const sf::FloatRect& other) const {
			getLocalBounds().contains(other);
		}
				
		virtual void render(sf::RenderTarget* t) {
			GameScene::render(t);
			t->draw(*m_el);
		};
	};	
}

