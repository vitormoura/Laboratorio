#include "GameScene.h"
#include "Game.h"
#include <iostream>

namespace my {

	GameScene::GameScene() : m_children(), m_parent(nullptr) {
	}

	GameScene::~GameScene() {

		#if _DEBUG
		std::cout << "GameScene::~GameScene" << std::endl;
		#endif		

		for (auto p : m_children)
			delete p;
	}

	void GameScene::init() {
	}

	void GameScene::update(sf::Time t) {
		for (auto s : m_children) {
			s->update(t);
		}
	}

	void GameScene::render(sf::RenderTarget* t) {
		for (auto s : m_children) {
			s->render(t);
		}
	}
}